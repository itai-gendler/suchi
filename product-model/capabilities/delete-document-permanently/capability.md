---
domain: trash
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/trash.go#PurgeTrashDocument
- kind: code
  role: implementation
  target: core/trash/trash.go#PurgeExpired
---

# Delete a document permanently

A document in Trash is deleted for good when its owner or an administrator confirms, or by itself 30 days after it was trashed. Links other documents held to it go with it.
