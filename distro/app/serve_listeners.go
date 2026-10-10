// SPDX-License-Identifier: AGPL-3.0-or-later

package app

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/johnnybravo-xyz/suchi/core/config"
)

type httpEndpoint struct {
	name      string
	addr      string
	publicURL string
	server    *http.Server
	listener  net.Listener
}

func bindHTTPEndpoints(ctx context.Context, cfg *config.Config, mainHandler, shareHandler http.Handler) ([]httpEndpoint, error) {
	if cfg == nil || mainHandler == nil {
		return nil, errors.New("bind HTTP endpoints: main configuration and handler are required")
	}
	if (cfg.ShareListenAddr == "") != (cfg.SharePublicURL == "") {
		return nil, errors.New("bind HTTP endpoints: share address and public URL must both be set")
	}
	var tlsConfig *tls.Config
	if cfg.TLSCertFile != "" {
		certificate, err := tls.LoadX509KeyPair(cfg.TLSCertFile, cfg.TLSKeyFile)
		if err != nil {
			return nil, fmt.Errorf("load HTTP TLS certificate: %w", err)
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{certificate},
			MinVersion:   tls.VersionTLS12,
		}
	}

	type endpointSpec struct {
		name, addr, publicURL string
		handler               http.Handler
	}
	specs := []endpointSpec{{name: "main", addr: cfg.ListenAddr, publicURL: cfg.PublicURL, handler: mainHandler}}
	if cfg.ShareListenAddr != "" {
		if shareHandler == nil {
			return nil, errors.New("bind HTTP endpoints: share handler is required when SHARE_LISTEN_ADDR is set")
		}
		specs = append(specs, endpointSpec{
			name: "share", addr: cfg.ShareListenAddr, publicURL: cfg.SharePublicURL, handler: shareHandler,
		})
	}

	endpoints := make([]httpEndpoint, 0, len(specs))
	closeBound := func() {
		for i := range endpoints {
			_ = endpoints[i].listener.Close()
		}
	}
	for _, spec := range specs {
		listener, err := (&net.ListenConfig{}).Listen(ctx, "tcp", spec.addr)
		if err != nil {
			closeBound()
			return nil, fmt.Errorf("bind %s listener %s: %w", spec.name, spec.addr, err)
		}
		if tlsConfig != nil {
			listener = tls.NewListener(listener, tlsConfig)
		}
		endpoints = append(endpoints, httpEndpoint{
			name:      spec.name,
			addr:      spec.addr,
			publicURL: spec.publicURL,
			server: &http.Server{
				Addr:              spec.addr,
				Handler:           spec.handler,
				ReadHeaderTimeout: 10 * time.Second,
				IdleTimeout:       60 * time.Second,
				BaseContext: func(net.Listener) context.Context {
					return ctx
				},
			},
			listener: listener,
		})
	}
	return endpoints, nil
}

func closeHTTPEndpoints(endpoints []httpEndpoint) {
	for i := range endpoints {
		_ = endpoints[i].listener.Close()
	}
}

type httpServeResult struct {
	name string
	err  error
}

func runHTTPEndpoints(ctx context.Context, cancel context.CancelFunc, endpoints []httpEndpoint, log *slog.Logger) error {
	if len(endpoints) == 0 {
		return errors.New("serve HTTP endpoints: no listeners")
	}
	results := make(chan httpServeResult, len(endpoints))
	for i := range endpoints {
		endpoint := &endpoints[i]
		log.Info(endpoint.name+".serve", "addr", endpoint.addr, "public_url", endpoint.publicURL)
		go func() {
			err := endpoint.server.Serve(endpoint.listener)
			results <- httpServeResult{name: endpoint.name, err: err}
		}()
	}

	var shutdownOnce sync.Once
	shutdownDone := make(chan struct{})
	shutdown := func() {
		shutdownOnce.Do(func() {
			log.Info("main.shutdown.begin")
			shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer shutdownCancel()
			var wait sync.WaitGroup
			for i := range endpoints {
				wait.Add(1)
				go func(server *http.Server) {
					defer wait.Done()
					_ = server.Shutdown(shutdownCtx)
				}(endpoints[i].server)
			}
			wait.Wait()
			close(shutdownDone)
		})
		<-shutdownDone
	}
	go func() {
		<-ctx.Done()
		shutdown()
	}()

	first := <-results
	wasCanceled := ctx.Err() != nil
	var serveErr error
	if !wasCanceled {
		switch {
		case first.err == nil:
			serveErr = fmt.Errorf("%s listener stopped unexpectedly", first.name)
		case errors.Is(first.err, http.ErrServerClosed):
			serveErr = fmt.Errorf("%s listener closed unexpectedly", first.name)
		default:
			serveErr = fmt.Errorf("%s listener: %w", first.name, first.err)
		}
		cancel()
	}
	shutdown()
	for range len(endpoints) - 1 {
		result := <-results
		if serveErr == nil && !wasCanceled && result.err != nil && !errors.Is(result.err, http.ErrServerClosed) {
			serveErr = fmt.Errorf("%s listener: %w", result.name, result.err)
		}
	}
	if serveErr != nil {
		log.Error("main.serve.err", "err", serveErr.Error())
		return fmt.Errorf("main.serve.err: %w", serveErr)
	}
	log.Info("main.shutdown.done")
	return nil
}
