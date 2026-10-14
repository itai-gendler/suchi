---
references:
- kind: code
  role: implementation
  target: core/api/saved_views.go#CreateSavedView
---

# Saved view

A named, reusable document search kept by its owner in one filing system, optionally shared with everyone who may enter it.

## Information kept

- **Name** — the view name, unique for its owner
- **Filter** — the search text, category, tag, correspondent, custom-field presence, sensitivity and date filters, or an exact set of documents
- **Shared** — whether others in the filing system see the view
