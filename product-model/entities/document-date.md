---
references:
- kind: code
  role: implementation
  target: core/api/intelligence.go#ResolveIntelligence
- kind: doc
  role: intent
  target: docs/document-dates.mdx
---

# Document date

An important date the document model found in a document's text, kept with the exact text it came from and shown on Calendar once accepted.

## Information kept

- **Date** — the date, at its known precision
- **Precision** — day, month or year
- **Role** — issued, due, start, end, expiry, renewal, service or other
- **Document text** — the exact text the date was read from
- **Model confidence** — the model's score for this date
- **Reviewed by** — who accepted or rejected it and when, or automatic acceptance

## States

### Proposed

Waiting for review in Approvals.

### Accepted

On Calendar and usable in date queries.

### Rejected

Dismissed; never shown on Calendar.
