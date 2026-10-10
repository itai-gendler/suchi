// SPDX-License-Identifier: AGPL-3.0-or-later

package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	pluginapi "github.com/johnnybravo-xyz/suchi/plugin-api"
)

func TestCreateShareLinkRequiresVisibilityAndBoundsPublicExpiry(t *testing.T) {
	s := newStatsServer(t)
	s.PublicURL = "https://archive.internal.example"
	category := seedStatsJDInbox(t, s.DB)
	docID := seedStatsDoc(t, s.DB, 1, "share-visibility", "Share visibility", category, false, 0)
	principal := &pluginapi.Principal{Kind: "user", UserID: 1, Role: "admin", Display: "Owner"}

	for _, tc := range []struct {
		name string
		body string
		code int
	}{
		{"missing visibility", `{"doc_ids":[` + itoa(docID) + `]}`, http.StatusBadRequest},
		{"public without expiry", `{"doc_ids":[` + itoa(docID) + `],"visibility":"public"}`, http.StatusBadRequest},
		{"public beyond one year", `{"doc_ids":[` + itoa(docID) + `],"visibility":"public","expires_in_sec":31536001}`, http.StatusBadRequest},
		{"unknown visibility", `{"doc_ids":[` + itoa(docID) + `],"visibility":"private"}`, http.StatusBadRequest},
		{"public ingress unavailable", `{"doc_ids":[` + itoa(docID) + `],"visibility":"public","expires_in_sec":604800}`, http.StatusConflict},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rec := shareCall(t, s, http.MethodPost, "/api/share_links/", tc.body, principal)
			if rec.Code != tc.code {
				t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
			}
		})
	}

	s.SharePublicURL = "https://share.example.test:8443"
	createdAt := time.Now().Unix()
	rec := shareCall(t, s, http.MethodPost, "/api/share_links/",
		`{"doc_ids":[`+itoa(docID)+`],"visibility":"public","expires_in_sec":604800}`, principal)
	if rec.Code != http.StatusCreated {
		t.Fatalf("public create status=%d body=%s", rec.Code, rec.Body.String())
	}
	var created struct {
		ID         int64  `json:"id"`
		Visibility string `json:"visibility"`
		ExpiresAt  int64  `json:"expires_at"`
		PublicURL  string `json:"public_url"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatal(err)
	}
	if created.Visibility != "public" || !strings.HasPrefix(created.PublicURL, s.SharePublicURL+"/s/") {
		t.Fatalf("created share=%+v", created)
	}
	if created.ExpiresAt < createdAt+604800 || created.ExpiresAt > time.Now().Unix()+604800 {
		t.Fatalf("expires_at=%d outside creation window", created.ExpiresAt)
	}
	var visibility string
	var expiresAt int64
	if err := s.DB.Read.QueryRow(`SELECT visibility,expires_at FROM share_links WHERE id=?`, created.ID).Scan(&visibility, &expiresAt); err != nil {
		t.Fatal(err)
	}
	if visibility != "public" || expiresAt != created.ExpiresAt {
		t.Fatalf("stored visibility=%q expires_at=%d", visibility, expiresAt)
	}

	rec = shareCall(t, s, http.MethodPost, "/api/share_links/",
		`{"doc_ids":[`+itoa(docID)+`],"visibility":"network"}`, principal)
	if rec.Code != http.StatusCreated {
		t.Fatalf("network create status=%d body=%s", rec.Code, rec.Body.String())
	}
	created = struct {
		ID         int64  `json:"id"`
		Visibility string `json:"visibility"`
		ExpiresAt  int64  `json:"expires_at"`
		PublicURL  string `json:"public_url"`
	}{}
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatal(err)
	}
	if created.Visibility != "network" || !strings.HasPrefix(created.PublicURL, s.PublicURL+"/s/") || created.ExpiresAt != 0 {
		t.Fatalf("network share=%+v", created)
	}
}

func TestListShareLinksFiltersBeforePagination(t *testing.T) {
	s := newStatsServer(t)
	s.PublicURL = "https://archive.internal.example"
	s.SharePublicURL = "https://share.example.test"
	now := time.Now().Unix()
	if _, err := s.DB.ExecWrite(t.Context(), `
		INSERT INTO share_links(system_id,token,doc_ids_json,created_by,label,visibility,expires_at,created_at) VALUES
			(1,'active-public','[42]',1,'active public','public',?,?),
			(1,'active-other','[43]',1,'active other','network',NULL,?),
			(1,'expired-target','[42]',1,'expired','network',?,?),
			(1,'revoked-target','[42]',1,'revoked','network',NULL,?),
			(1,'invalid-public','[42]',1,'invalid public','public',NULL,?)
	`, now+3600, now-10, now-20, now-1, now-30, now-40, now-50); err != nil {
		t.Fatal(err)
	}
	if _, err := s.DB.ExecWrite(t.Context(),
		`UPDATE share_links SET revoked_at=? WHERE token='revoked-target'`, now-1); err != nil {
		t.Fatal(err)
	}
	principal := &pluginapi.Principal{Kind: "user", UserID: 1, Role: "admin"}

	active := shareCall(t, s, http.MethodGet, "/api/share_links/?doc_id=42&status=active&page_size=1", "", principal)
	var activePage struct {
		Count   int            `json:"count"`
		Results []ShareLinkRow `json:"results"`
	}
	if active.Code != http.StatusOK || json.Unmarshal(active.Body.Bytes(), &activePage) != nil {
		t.Fatalf("active status=%d body=%s", active.Code, active.Body.String())
	}
	if activePage.Count != 1 || len(activePage.Results) != 1 || activePage.Results[0].Visibility != "public" ||
		activePage.Results[0].PublicURL != s.SharePublicURL+"/s/active-public" {
		t.Fatalf("active page=%+v", activePage)
	}
	if strings.Contains(active.Body.String(), "view_count") {
		t.Fatalf("unused view count leaked: %s", active.Body.String())
	}

	inactive := shareCall(t, s, http.MethodGet, "/api/share_links/?doc_id=42&status=inactive&page_size=1", "", principal)
	var inactivePage struct {
		Count   int            `json:"count"`
		Results []ShareLinkRow `json:"results"`
	}
	if inactive.Code != http.StatusOK || json.Unmarshal(inactive.Body.Bytes(), &inactivePage) != nil {
		t.Fatalf("inactive status=%d body=%s", inactive.Code, inactive.Body.String())
	}
	if inactivePage.Count != 3 || len(inactivePage.Results) != 1 {
		t.Fatalf("inactive page=%+v", inactivePage)
	}

	for _, path := range []string{
		"/api/share_links/?doc_id=not-a-number",
		"/api/share_links/?doc_id=0",
		"/api/share_links/?status=all",
	} {
		rec := shareCall(t, s, http.MethodGet, path, "", principal)
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("%s status=%d body=%s", path, rec.Code, rec.Body.String())
		}
	}
}

func TestPublicShareExpiryIsRecheckedAndMetadataReadsDoNotWrite(t *testing.T) {
	s := newStatsServer(t)
	s.PublicURL = "https://archive.internal.example"
	category := seedStatsJDInbox(t, s.DB)
	docID := seedStatsDoc(t, s.DB, 1, "share-read-only", "Read-only share", category, false, 0)
	now := time.Now().Unix()
	networkToken := strings.Repeat("a", 64)
	missingExpiryToken := strings.Repeat("b", 64)
	oversizedExpiryToken := strings.Repeat("c", 64)
	if _, err := s.DB.ExecWrite(t.Context(), `
		INSERT INTO share_links(system_id,token,doc_ids_json,created_by,visibility,created_at)
		VALUES(1,?,'[`+itoa(docID)+`]',1,'network',?)`, networkToken, now); err != nil {
		t.Fatal(err)
	}
	if _, err := s.DB.ExecWrite(t.Context(), `
		INSERT INTO share_links(system_id,token,doc_ids_json,created_by,visibility,created_at)
		VALUES(1,?,'[`+itoa(docID)+`]',1,'public',?)`, missingExpiryToken, now); err != nil {
		t.Fatal(err)
	}
	if _, err := s.DB.ExecWrite(t.Context(), `
		INSERT INTO share_links(system_id,token,doc_ids_json,created_by,visibility,expires_at,created_at)
		VALUES(1,?,'[`+itoa(docID)+`]',1,'public',?,?)`,
		oversizedExpiryToken, now+maxShareExpirySeconds+1, now); err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	s.RegisterShareSurface(mux, ShareSurface{Origin: s.PublicURL})
	request := func(token string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/s/"+token, http.NoBody)
		req.Header.Set("Accept", "application/json")
		mux.ServeHTTP(rec, req)
		return rec
	}
	if rec := request(networkToken); rec.Code != http.StatusOK {
		t.Fatalf("network status=%d body=%s", rec.Code, rec.Body.String())
	}
	for _, token := range []string{missingExpiryToken, oversizedExpiryToken} {
		if rec := request(token); rec.Code != http.StatusNotFound {
			t.Fatalf("invalid public token status=%d body=%s", rec.Code, rec.Body.String())
		}
	}
	var views int64
	if err := s.DB.Read.QueryRow(`SELECT view_count FROM share_links WHERE token=?`, networkToken).Scan(&views); err != nil {
		t.Fatal(err)
	}
	if views != 0 {
		t.Fatalf("metadata read wrote view_count=%d", views)
	}
}

func TestPublicShareSurfaceIsolatedFromNetworkLinksAndHostHeaders(t *testing.T) {
	s := newStatsServer(t)
	category := seedStatsJDInbox(t, s.DB)
	docID := seedStatsDoc(t, s.DB, 1, "share-surface", "Share surface", category, false, 0)
	now := time.Now().Unix()
	networkToken := strings.Repeat("d", 64)
	publicToken := strings.Repeat("e", 64)
	if _, err := s.DB.ExecWrite(t.Context(), `
		INSERT INTO share_links(system_id,token,doc_ids_json,created_by,label,visibility,created_at)
		VALUES(1,?,'[`+itoa(docID)+`]',1,'Network share','network',?)`, networkToken, now); err != nil {
		t.Fatal(err)
	}
	if _, err := s.DB.ExecWrite(t.Context(), `
		INSERT INTO share_links(system_id,token,doc_ids_json,created_by,label,visibility,expires_at,created_at)
		VALUES(1,?,'[`+itoa(docID)+`]',1,'Public share','public',?,?)`,
		publicToken, now+3600, now); err != nil {
		t.Fatal(err)
	}

	publicMux := http.NewServeMux()
	s.RegisterShareSurface(publicMux, ShareSurface{
		PublicOnly: true,
		Origin:     "https://share.example.test:8443",
	})
	publicMux.HandleFunc("/", http.NotFound)
	request := func(method, path, accept string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(method, path, http.NoBody)
		req.Host = "attacker.example"
		req.Header.Set("X-Forwarded-Host", "also-attacker.example")
		req.Header.Set("X-Forwarded-Proto", "http")
		if accept != "" {
			req.Header.Set("Accept", accept)
		}
		publicMux.ServeHTTP(rec, req)
		return rec
	}

	for _, tc := range []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/"},
		{http.MethodGet, "/api/whoami"},
		{http.MethodGet, "/assets/brand/favicon.svg"},
		{http.MethodGet, "/s/" + networkToken},
		{http.MethodPost, "/s/" + networkToken},
		{http.MethodGet, "/s/" + networkToken + "/" + itoa(docID) + "/download"},
		{http.MethodGet, "/s/short"},
		{http.MethodDelete, "/s/" + publicToken},
	} {
		rec := request(tc.method, tc.path, "application/json")
		if rec.Code != http.StatusNotFound {
			t.Fatalf("%s %s status=%d body=%s", tc.method, tc.path, rec.Code, rec.Body.String())
		}
	}

	rec := request(http.MethodGet, "/s/"+publicToken, "application/json")
	if rec.Code != http.StatusOK {
		t.Fatalf("public status=%d body=%s", rec.Code, rec.Body.String())
	}
	var payload ShareLinkPublic
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatal(err)
	}
	if payload.InstanceHost != "share.example.test" || payload.ExpiresAt != now+3600 {
		t.Fatalf("public payload=%+v", payload)
	}

	htmlResponse := request(http.MethodGet, "/s/"+publicToken, "text/html")
	if htmlResponse.Code != http.StatusOK || !strings.Contains(htmlResponse.Body.String(), "Available until") {
		t.Fatalf("HTML status=%d body=%s", htmlResponse.Code, htmlResponse.Body.String())
	}
	if strings.Contains(htmlResponse.Body.String(), "favicon") || strings.Contains(htmlResponse.Body.String(), "/assets/") {
		t.Fatalf("share-only HTML references unavailable assets: %s", htmlResponse.Body.String())
	}

	mainMux := http.NewServeMux()
	s.RegisterShareSurface(mainMux, ShareSurface{Origin: "https://archive.internal.example:9443"})
	mainRec := httptest.NewRecorder()
	mainReq := httptest.NewRequest(http.MethodGet, "/s/"+networkToken, http.NoBody)
	mainReq.Header.Set("Accept", "application/json")
	mainMux.ServeHTTP(mainRec, mainReq)
	if mainRec.Code != http.StatusOK {
		t.Fatalf("main network status=%d body=%s", mainRec.Code, mainRec.Body.String())
	}
	payload = ShareLinkPublic{}
	if err := json.Unmarshal(mainRec.Body.Bytes(), &payload); err != nil {
		t.Fatal(err)
	}
	if payload.InstanceHost != "archive.internal.example" {
		t.Fatalf("main payload host=%q", payload.InstanceHost)
	}
}
