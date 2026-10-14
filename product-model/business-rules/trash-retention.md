---
appliesTo:
- type: capability
  id: move-document-to-trash
- type: capability
  id: restore-document
- type: capability
  id: delete-document-permanently
references:
- kind: code
  role: implementation
  target: core/trash/trash.go
- kind: doc
  role: intent
  target: docs/web-app.mdx
---

# A trashed document can be restored for 30 days, then it is deleted for good

While in Trash a document is read-only, leaves every list, search, link and share, and shows its deletion date.
