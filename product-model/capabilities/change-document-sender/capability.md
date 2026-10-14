---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/documents_bulk.go#BulkEdit
- kind: code
  role: implementation
  target: core/api/document_correspondents.go
- kind: doc
  role: intent
  target: docs/web-app.mdx
---

# Change the document sender

A person chooses an existing correspondent as a document's primary sender, or clears the sender, from the document or for a selection. Recipient, cc and other correspondent roles stay as they are.
