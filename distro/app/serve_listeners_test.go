// SPDX-License-Identifier: AGPL-3.0-or-later

package app

import (
	"context"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/johnnybravo-xyz/suchi/core/config"
)

func TestBindHTTPEndpointsRejectsAliasedAddressesAtomically(t *testing.T) {
	probe, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := probe.Addr().String()
	if err := probe.Close(); err != nil {
		t.Fatal(err)
	}
	cfg := &config.Config{
		ListenAddr:      addr,
		PublicURL:       "http://" + addr,
		ShareListenAddr: addr,
		SharePublicURL:  "https://share.example.test",
	}
	handler := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	if _, err := bindHTTPEndpoints(t.Context(), cfg, handler, handler); err == nil || !strings.Contains(err.Error(), "bind share listener") {
		t.Fatalf("bind error=%v", err)
	}

	// The failed second bind must release the already-bound main socket.
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatalf("main listener leaked after atomic bind failure: %v", err)
	}
	_ = listener.Close()
}

func TestRunHTTPEndpointsStopsSiblingOnListenerFailure(t *testing.T) {
	cfg := &config.Config{
		ListenAddr:      "127.0.0.1:0",
		PublicURL:       "http://archive.example.test",
		ShareListenAddr: "127.0.0.1:0",
		SharePublicURL:  "https://share.example.test",
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, "ok")
	})
	ctx, cancel := context.WithCancel(t.Context())
	endpoints, err := bindHTTPEndpoints(ctx, cfg, handler, handler)
	if err != nil {
		t.Fatal(err)
	}
	if len(endpoints) != 2 {
		t.Fatalf("endpoints=%d", len(endpoints))
	}
	mainAddr := endpoints[0].listener.Addr().String()
	result := make(chan error, 1)
	go func() { result <- runHTTPEndpoints(ctx, cancel, endpoints, testLogger()) }()
	if err := endpoints[1].listener.Close(); err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-result:
		if err == nil || !strings.Contains(err.Error(), "share listener") {
			t.Fatalf("serve error=%v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("listener failure did not stop sibling")
	}

	conn, err := net.DialTimeout("tcp", mainAddr, 200*time.Millisecond)
	if err == nil {
		_ = conn.Close()
		t.Fatal("main listener still accepts after share listener failure")
	}
}
