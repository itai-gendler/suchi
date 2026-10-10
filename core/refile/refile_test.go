// SPDX-License-Identifier: AGPL-3.0-or-later

package refile

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"path/filepath"
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/automations"
	"github.com/johnnybravo-xyz/suchi/core/db"
	"github.com/johnnybravo-xyz/suchi/core/db/migrations"
)

func TestRenderOnlyRefileSelectsOneSystemAcrossOwners(t *testing.T) {
	d, log := refileTestDB(t)
	stats, err := All(t.Context(), d, testActions(t), log, Options{SystemID: 2, SkipAutomations: true})
	if err != nil {
		t.Fatal(err)
	}
	if stats.DocsScanned != 2 || stats.RenderEnqueued != 2 || stats.AutomationsApplied != 0 || stats.Errors != 0 {
		t.Fatalf("wrong refile selection: %+v", stats)
	}
	rows, err := d.Read.Query(`SELECT kind,doc_id,system_id FROM jobs ORDER BY doc_id`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	var ids []int64
	for rows.Next() {
		var kind string
		var docID, systemID int64
		if err := rows.Scan(&kind, &docID, &systemID); err != nil {
			t.Fatal(err)
		}
		if kind != "render" || systemID != 2 {
			t.Fatalf("refile scheduled extraction or foreign work: %s %d %d", kind, docID, systemID)
		}
		ids = append(ids, docID)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	if len(ids) != 2 || ids[0] != 2 || ids[1] != 3 {
		t.Fatalf("foreign or missing documents: %v", ids)
	}
	if _, err := All(t.Context(), d, testActions(t), log, Options{SkipAutomations: true}); err == nil {
		t.Fatal("unqualified refile admitted")
	}
}

func TestRefileRefreshesRuleTitlesAndPreservesHumanTitles(t *testing.T) {
	d, log := refileTestDB(t)
	if _, err := d.ExecWrite(t.Context(), `
		UPDATE documents SET title='Old rule title',title_owner='rule' WHERE id=1;
		INSERT INTO documents(system_id,id,owner_id,original_blob,original_size,title,title_owner,jd_category_id,created_at,updated_at)
		VALUES(1,4,1,'human',1,'Human title','human',1,0,0)
	`); err != nil {
		t.Fatal(err)
	}
	store := automations.New(d, testActions(t))
	if _, err := store.Create(t.Context(), 1, automations.Automation{
		Name: "retitle", Enabled: true,
		Triggers: []automations.Trigger{{Type: automations.TriggerDocumentAdded}},
		Actions: []automations.Action{{
			Kind: "assign_title", Params: map[string]any{"template": "Updated rule title"},
		}},
	}); err != nil {
		t.Fatal(err)
	}
	stats, err := All(t.Context(), d, testActions(t), log, Options{SystemID: 1, SkipRender: true})
	if err != nil {
		t.Fatal(err)
	}
	if stats.DocsScanned != 2 || stats.AutomationsApplied != 2 || stats.Errors != 0 {
		t.Fatalf("refile stats=%+v", stats)
	}
	rows, err := d.Read.Query(`SELECT id,title,title_owner FROM documents WHERE system_id=1 ORDER BY id`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	want := []struct {
		id    int64
		title string
		owner string
	}{{1, "Updated rule title", "rule"}, {4, "Human title", "human"}}
	i := 0
	for ; rows.Next(); i++ {
		if i >= len(want) {
			t.Fatal("unexpected document")
		}
		var id int64
		var title, owner string
		if err := rows.Scan(&id, &title, &owner); err != nil {
			t.Fatal(err)
		}
		if id != want[i].id || title != want[i].title || owner != want[i].owner {
			t.Fatalf("document %d title=%q owner=%q", id, title, owner)
		}
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	if i != len(want) {
		t.Fatalf("documents=%d, want %d", i, len(want))
	}
}

func refileTestDB(t *testing.T) (*db.DB, *slog.Logger) {
	t.Helper()
	d, err := db.Open(t.Context(), filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { d.Close() })
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	migs, err := db.LoadMigrations(migrations.FS, ".")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Migrate(t.Context(), d, migs, log); err != nil {
		t.Fatal(err)
	}
	_, err = d.ExecWrite(t.Context(), `INSERT INTO users(id,email,display_name,role,created_at,updated_at) VALUES(1,'first@test','First','admin',0,0),(2,'second@test','Second','member',0,0);
 UPDATE jd_systems SET code='S01' WHERE id=1;
 INSERT INTO jd_systems(id,code,name,taxonomy,created_at,updated_at) VALUES(2,'S02','Other','jd',0,0);
 INSERT INTO jd_areas(system_id,code_start,code_end,name,position) VALUES(1,10,19,'Records',0),(2,10,19,'Records',0);
 INSERT INTO jd_categories(system_id,id,area_start,code,name,system) VALUES(1,1,10,13,'Tax',0),(2,2,10,13,'Tax',0);
 INSERT INTO documents(system_id,id,owner_id,original_blob,original_size,title,jd_category_id,created_at,updated_at) VALUES(1,1,1,'same',0,'Original',1,0,0),(2,2,1,'same',0,'Other first owner',2,0,0),(2,3,2,'same',0,'Other second owner',2,0,0);`)
	if err != nil {
		t.Fatal(err)
	}
	return d, log
}

// Pause through the existing logger after selection, before any writer turn.
type refilePauseWriter struct{ reached, release chan struct{} }

func (w refilePauseWriter) Write(data []byte) (int, error) {
	if bytes.Contains(data, []byte("msg=refile.begin")) {
		close(w.reached)
		<-w.release
	}
	return len(data), nil
}

func TestRefileRejectsActorLosingAdminBeforeWriterTurn(t *testing.T) {
	for _, change := range []string{"role='member'", "disabled=1"} {
		t.Run(change, func(t *testing.T) {
			d, log := refileTestDB(t)
			pause := refilePauseWriter{make(chan struct{}), make(chan struct{})}
			result := make(chan error, 1)
			go func() {
				_, err := All(t.Context(), d, testActions(t), slog.New(slog.NewTextHandler(pause, nil)), Options{SystemID: 2, ActorID: 1, SkipAutomations: true})
				result <- err
			}()
			<-pause.reached
			if _, err := d.ExecWrite(t.Context(), `UPDATE users SET `+change+` WHERE id=1`); err != nil {
				close(pause.release)
				t.Fatal(err)
			}
			close(pause.release)
			if err := <-result; !errors.Is(err, ErrActorUnavailable) {
				t.Fatalf("actor loss was not rejected: %v", err)
			}
			var jobs int
			if err := d.Read.QueryRow(`SELECT COUNT(*) FROM jobs`).Scan(&jobs); err != nil {
				t.Fatal(err)
			}
			if jobs != 0 {
				t.Fatalf("rejected refile left %d jobs", jobs)
			}
			// Trusted local CLI remains an explicit separate actor, not a stale API role.
			stats, err := All(t.Context(), d, testActions(t), log, Options{SystemID: 2, SkipAutomations: true})
			if err != nil || stats.RenderEnqueued != 2 {
				t.Fatalf("trusted CLI refile failed: %+v %v", stats, err)
			}
		})
	}
}
