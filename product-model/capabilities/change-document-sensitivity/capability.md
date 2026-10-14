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
---

# Change document sensitivity

A person sets a document's sensitivity to Public, Internal, Confidential, Restricted or Unset, from the document or for a selection. Confidential and Restricted documents hide their thumbnail, preview and text until revealed.
