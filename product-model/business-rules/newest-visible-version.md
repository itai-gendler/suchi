---
appliesTo:
- type: capability
  id: browse-documents
- type: capability
  id: search-documents
- type: capability
  id: browse-calendar
- type: capability
  id: ask-documents
references:
- kind: code
  role: implementation
  target: core/api/documents_list.go
- kind: doc
  role: intent
  target: docs/document-versions-and-links.mdx
---

# Lists, search, Calendar and questions show the newest version each person may see

A newer version the person cannot see, or that is in Trash, never hides the older one or reveals that it exists.
