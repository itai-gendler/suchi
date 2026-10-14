---
domain: trash
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/documents.go#RestoreDocument
- kind: code
  role: implementation
  target: core/trash/trash.go
---

# Restore a document

The owner of a document in Trash, or an administrator, restores it within 30 days. It returns to where it was filed.
