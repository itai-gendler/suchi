---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/documents.go#PatchDocument
- kind: code
  role: implementation
  target: core/api/documents_bulk.go#BulkEdit
- kind: doc
  role: intent
  target: docs/jd.mdx
---

# File a document

A person files a document under one category of the filing tree, from the document itself or for a selection in Documents. Its filing address follows the category; its number never changes.
