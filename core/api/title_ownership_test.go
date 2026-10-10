// SPDX-License-Identifier: AGPL-3.0-or-later

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/auth"
	"github.com/johnnybravo-xyz/suchi/core/automations"
)

func TestPatchDocumentTitleOwnership(t *testing.T) {
	s := newBulkServer(t)
	docID := seedStatsDoc(t, s.DB, 1, "title-owner", "Generated title", seedStatsJDInbox(t, s.DB), false, 0)

	patch := func(body string) *httptest.ResponseRecorder {
		t.Helper()
		req := httptest.NewRequest(http.MethodPatch, "/api/documents/"+strconv.FormatInt(docID, 10), bytes.NewBufferString(body)).
			WithContext(auth.WithPrincipal(context.Background(), adminPrincipal(1)))
		req.SetPathValue("id", strconv.FormatInt(docID, 10))
		rec := httptest.NewRecorder()
		s.PatchDocument(rec, req)
		return rec
	}
	read := func() (title, owner string) {
		t.Helper()
		if err := s.DB.Read.QueryRow(`SELECT title,title_owner FROM documents WHERE id=?`, docID).Scan(&title, &owner); err != nil {
			t.Fatal(err)
		}
		return title, owner
	}

	if rec := patch(`{"title":"Household receipt"}`); rec.Code != http.StatusOK {
		t.Fatalf("rename status=%d body=%s", rec.Code, rec.Body.String())
	}
	if title, owner := read(); title != "Household receipt" || owner != "human" {
		t.Fatalf("rename stored title=%q owner=%q", title, owner)
	}
	var beforeOwner, afterOwner string
	if err := s.DB.Read.QueryRow(`
		SELECT json_extract(before_json,'$.title_owner'), json_extract(after_json,'$.title_owner')
		FROM audit_events WHERE action='document.update' AND object_id=? ORDER BY id LIMIT 1
	`, docID).Scan(&beforeOwner, &afterOwner); err != nil {
		t.Fatal(err)
	}
	if beforeOwner != "auto" || afterOwner != "human" {
		t.Fatalf("rename audit owner %q -> %q", beforeOwner, afterOwner)
	}

	if _, err := s.DB.Write.Exec(`UPDATE documents SET title_owner='auto' WHERE id=?`, docID); err != nil {
		t.Fatal(err)
	}
	if rec := patch(`{"title":"Household receipt"}`); rec.Code != http.StatusOK {
		t.Fatalf("same-value rename status=%d body=%s", rec.Code, rec.Body.String())
	}
	if _, owner := read(); owner != "human" {
		t.Fatalf("same-value rename owner=%q, want human", owner)
	}

	if rec := patch(`{"title_owner":"auto"}`); rec.Code != http.StatusOK {
		t.Fatalf("release status=%d body=%s", rec.Code, rec.Body.String())
	}
	if title, owner := read(); title != "Household receipt" || owner != "auto" {
		t.Fatalf("release stored title=%q owner=%q", title, owner)
	}

	if rec := patch(`{"title":"Suggested starting point","title_owner":"auto"}`); rec.Code != http.StatusOK {
		t.Fatalf("automatic title status=%d body=%s", rec.Code, rec.Body.String())
	}
	if title, owner := read(); title != "Suggested starting point" || owner != "auto" {
		t.Fatalf("automatic title stored title=%q owner=%q", title, owner)
	}

	if rec := patch(`{"title_owner":"rule"}`); rec.Code != http.StatusBadRequest || !strings.Contains(rec.Body.String(), "invalid_title_owner") {
		t.Fatalf("rule owner status=%d body=%s", rec.Code, rec.Body.String())
	}

	req := httptest.NewRequest(http.MethodGet, "/api/documents/"+strconv.FormatInt(docID, 10), nil).
		WithContext(auth.WithPrincipal(context.Background(), adminPrincipal(1)))
	req.SetPathValue("id", strconv.FormatInt(docID, 10))
	rec := httptest.NewRecorder()
	s.GetDocument(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status=%d body=%s", rec.Code, rec.Body.String())
	}
	var detail DocumentDetail
	if err := json.Unmarshal(rec.Body.Bytes(), &detail); err != nil {
		t.Fatal(err)
	}
	if detail.TitleOwner != "auto" {
		t.Fatalf("detail title_owner=%q", detail.TitleOwner)
	}
}

func TestRenameSurvivesDocumentUpdatedTitleRule(t *testing.T) {
	s := newBulkServer(t)
	docID := seedStatsDoc(t, s.DB, 1, "rename-rule", "Generated title", seedStatsJDInbox(t, s.DB), false, 0)
	store := automations.New(s.DB, s.Actions)
	if _, err := store.Create(t.Context(), 1, automations.Automation{
		Name: "retitle updates", Enabled: true,
		Triggers: []automations.Trigger{{Type: automations.TriggerDocumentUpdated}},
		Actions: []automations.Action{{
			Kind: "assign_title", Params: map[string]any{"template": "Rule title"},
		}},
	}); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPatch, "/api/documents/"+strconv.FormatInt(docID, 10), strings.NewReader(`{"title":"My title"}`)).
		WithContext(auth.WithPrincipal(context.Background(), adminPrincipal(1)))
	req.SetPathValue("id", strconv.FormatInt(docID, 10))
	rec := httptest.NewRecorder()
	s.PatchDocument(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("rename status=%d body=%s", rec.Code, rec.Body.String())
	}
	var title, owner string
	if err := s.DB.Read.QueryRow(`SELECT title,title_owner FROM documents WHERE id=?`, docID).Scan(&title, &owner); err != nil {
		t.Fatal(err)
	}
	if title != "My title" || owner != "human" {
		t.Fatalf("rename after rule title=%q owner=%q", title, owner)
	}
}
