// SPDX-License-Identifier: AGPL-3.0-or-later

package api

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/auth"
)

func TestCreateAutomationAcceptsJSONNumberActionIDs(t *testing.T) {
	s := newBulkServer(t)
	seedUser(t, s.DB, 1)
	if _, err := s.DB.Write.ExecContext(context.Background(), `
		INSERT INTO correspondents(id, system_id, name, slug, created_at, updated_at)
		VALUES (1, 1, 'Existing correspondent', 'existing-correspondent', 0, 0);
		INSERT INTO tags(id, system_id, name, slug, created_at, updated_at)
		VALUES (1, 1, 'Existing tag', 'existing-tag', 0, 0)
	`); err != nil {
		t.Fatal(err)
	}

	for _, test := range []struct {
		name   string
		action string
	}{
		{name: "scalar id", action: `{"type":"assign_correspondent","params":{"correspondent_id":1}}`},
		{name: "id array", action: `{"type":"assign_tags","params":{"tag_ids":[1]}}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			body := []byte(fmt.Sprintf(`{
				"name":%q,
				"enabled":true,
				"triggers":[{"type":2}],
				"actions":[%s]
			}`, test.name, test.action))
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/automations/", bytes.NewReader(body))
			req = req.WithContext(auth.WithPrincipal(req.Context(), adminPrincipal(1)))

			s.CreateAutomation(rec, req)
			if rec.Code != http.StatusCreated {
				t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
			}
		})
	}
}
