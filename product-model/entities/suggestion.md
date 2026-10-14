---
references:
- kind: code
  role: implementation
  target: core/approvals/document_change.go
- kind: doc
  role: intent
  target: docs/approvals.mdx
---

# Suggestion

A proposed change to one field of a document from local archive matching, the document model or an automation question, waiting in Approvals for a decision.

## Information kept

- **Field** — the field it would change: filing category, correspondent, tag, title, language or a custom field
- **Current value** — the field's value when the suggestion was made
- **Suggested value** — the value it would set
- **Why suggested** — the reason and the supporting documents or text behind it
- **Model confidence** — the score the suggestion was made with
- **Suggested by** — archive matching, the document model, or the automation whose question produced it

## States

### Proposed

Waiting for the document owner or an administrator.

### Accepted

Applied to the document.

### Dismissed

Closed without changing the document.

### Outdated

Closed because the document or the field changed first.
