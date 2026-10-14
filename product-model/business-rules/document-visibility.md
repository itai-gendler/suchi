---
appliesTo:
- type: capability
  id: browse-documents
- type: capability
  id: search-documents
- type: capability
  id: view-document
- type: capability
  id: ask-documents
- type: capability
  id: view-dashboard
references:
- kind: code
  role: implementation
  target: core/api/authz_helpers.go
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# A person sees only the documents they own or were given access to, in filing systems they may enter

Administrators see every document of a filing system. Counts, pages, snippets, similar documents and answers never include documents a person may not open.
