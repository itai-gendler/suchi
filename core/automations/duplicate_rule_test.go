// SPDX-License-Identifier: AGPL-3.0-or-later

package automations_test

// Save (create or edit) refuses to produce a content-duplicate of an
// existing row; the store returns *ErrDuplicateRule pointing at the
// existing match.

import (
	"context"
	"errors"
	"testing"

	"github.com/johnnybravo-xyz/suchi/core/automations"
)

func TestStore_Update_ToggleIsInPlace(t *testing.T) {
	ctx := context.Background()
	d, _ := setup(t, ctx)

	s := automations.New(d, testActions(t))
	created, err := s.Create(ctx, 1, automations.Automation{Name: "user-owned",
		Enabled: true,
		Triggers: []automations.Trigger{{Type: automations.TriggerDocumentAdded,
			TypeCode: automations.TriggerToCode(automations.TriggerDocumentAdded)}}})
	if err != nil {
		t.Fatal(err)
	}
	disabled := false
	patched, err := s.Update(ctx, 1, created.ID, automations.AutomationPatch{Enabled: &disabled})
	if err != nil {
		t.Fatal(err)
	}
	if patched.ID != created.ID {
		t.Fatalf("patch must update in place; got new id %d", patched.ID)
	}
	if patched.Enabled {
		t.Fatalf("expected enabled=false")
	}
}

// Creating a rule whose triggers + actions content-match an existing
// row returns *ErrDuplicateRule pointing at the existing id. Enabled
// state and name are metadata; they don't disqualify a match.
func TestStore_Create_DuplicateRule_Refused(t *testing.T) {
	ctx := context.Background()
	d, _ := setup(t, ctx)
	s := automations.New(d, testActions(t))

	firstA, err := s.Create(ctx, 1, automations.Automation{Name: "first",
		Enabled: true,
		Triggers: []automations.Trigger{{
			Type:            automations.TriggerDocumentAdded,
			TypeCode:        automations.TriggerToCode(automations.TriggerDocumentAdded),
			FilterContentRE: "invoice",
		}},
		Actions: []automations.Action{{
			Kind:   "assign_title",
			Params: map[string]any{"template": "Filed {{title}}"},
		}}})
	if err != nil {
		t.Fatal(err)
	}

	_, err = s.Create(ctx, 1, automations.Automation{Name: "clone (different name, same content)",
		Enabled: false,
		Triggers: []automations.Trigger{{
			Type:            automations.TriggerDocumentAdded,
			TypeCode:        automations.TriggerToCode(automations.TriggerDocumentAdded),
			FilterContentRE: "invoice",
		}},
		Actions: []automations.Action{{
			Kind:   "assign_title",
			Params: map[string]any{"template": "Filed {{title}}"},
		}}})
	var dup *automations.ErrDuplicateRule
	if !errors.As(err, &dup) {
		t.Fatalf("expected ErrDuplicateRule, got %v", err)
	}
	if dup.ExistingID != firstA.ID {
		t.Errorf("ExistingID = %d, want %d", dup.ExistingID, firstA.ID)
	}
	if !dup.ExistingEnabled {
		t.Errorf("ExistingEnabled = false, want true")
	}
}

// Editing rule A to make its content match already-existing rule B
// refuses with *ErrDuplicateRule.
func TestStore_Update_DuplicateRule_Refused(t *testing.T) {
	ctx := context.Background()
	d, _ := setup(t, ctx)
	s := automations.New(d, testActions(t))

	a, err := s.Create(ctx, 1, automations.Automation{Name: "rule A",
		Triggers: []automations.Trigger{{
			Type: automations.TriggerDocumentAdded, FilterContentRE: "aaa",
		}}})
	if err != nil {
		t.Fatal(err)
	}
	b, err := s.Create(ctx, 1, automations.Automation{Name: "rule B",
		Triggers: []automations.Trigger{{
			Type: automations.TriggerDocumentAdded, FilterContentRE: "bbb",
		}}})
	if err != nil {
		t.Fatal(err)
	}

	// Point A at B's content — should refuse.
	newTriggers := []automations.Trigger{{
		Type: automations.TriggerDocumentAdded, FilterContentRE: "bbb",
	}}
	_, err = s.Update(ctx, 1, a.ID, automations.AutomationPatch{Triggers: &newTriggers})
	var dup *automations.ErrDuplicateRule
	if !errors.As(err, &dup) {
		t.Fatalf("expected ErrDuplicateRule, got %v", err)
	}
	if dup.ExistingID != b.ID {
		t.Errorf("ExistingID = %d, want %d", dup.ExistingID, b.ID)
	}
}

// Self-match must be allowed: an in-place patch that produces the
// same content is a no-op, not a duplicate.
func TestStore_Update_SelfMatchAllowed(t *testing.T) {
	ctx := context.Background()
	d, _ := setup(t, ctx)
	s := automations.New(d, testActions(t))

	created, err := s.Create(ctx, 1, automations.Automation{Name: "same shape",
		Triggers: []automations.Trigger{{
			Type: automations.TriggerDocumentAdded, FilterContentRE: "x",
		}}})
	if err != nil {
		t.Fatal(err)
	}
	// Patch with the same triggers — different pointer, same content.
	sameTriggers := []automations.Trigger{{
		Type: automations.TriggerDocumentAdded, FilterContentRE: "x",
	}}
	patched, err := s.Update(ctx, 1, created.ID, automations.AutomationPatch{Triggers: &sameTriggers})
	if err != nil {
		t.Fatalf("self-match must not error: %v", err)
	}
	if patched.ID != created.ID {
		t.Errorf("id = %d, want %d (in-place)", patched.ID, created.ID)
	}
}

// Two logically-equivalent triggers built with different map insert
// orders on Params produce byte-identical signatures. Guards against
// a regression where a Go map iteration change makes rules
// non-comparable.
func TestStore_Create_ParamsMapOrderIndependent(t *testing.T) {
	ctx := context.Background()
	d, _ := setup(t, ctx)
	s := automations.New(d, testActions(t))

	if _, err := s.Create(ctx, 1, automations.Automation{Name: "params-a",
		Triggers: []automations.Trigger{{
			Type: automations.TriggerDocumentAdded, FilterContentRE: "x",
		}},
		Actions: []automations.Action{{
			Kind:   "assign_title",
			Params: map[string]any{"template": "Filed {{title}}", "note": "keep"},
		}}}); err != nil {
		t.Fatal(err)
	}

	_, err := s.Create(ctx, 1, automations.Automation{Name: "params-b (different key insert order)",
		Triggers: []automations.Trigger{{
			Type: automations.TriggerDocumentAdded, FilterContentRE: "x",
		}},
		Actions: []automations.Action{{
			Kind:   "assign_title",
			Params: map[string]any{"note": "keep", "template": "Filed {{title}}"},
		}}})
	var dup *automations.ErrDuplicateRule
	if !errors.As(err, &dup) {
		t.Fatalf("expected ErrDuplicateRule despite map insert order, got %v", err)
	}
}
