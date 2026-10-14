---
references:
- kind: code
  role: implementation
  target: core/api/automations.go#CreateAutomation
- kind: doc
  role: intent
  target: docs/automations.mdx
---

# Automation

A deterministic filing rule of one filing system: event filters and an ordered list of metadata actions, optionally with one bounded question for the document model.

## Information kept

- **Name** — the automation name
- **Triggers** — the events and filters that select documents: when added, after extraction or after an update
- **Actions** — the ordered actions: title, tag, correspondent, category, folder layout, owner or custom-field assignment
- **Classifier question** — an optional Yes/No or multiple-choice question for the document model with an action per answer
- **Run order** — its position in the top-to-bottom run order

## States

### Enabled

The automation runs on matching documents.

### Disabled

The automation is kept but does not run.
