---
appliesTo:
- type: entity
  id: document
  effect: changes
  from: Trashed
permits:
- actors:
  - administrator
- related:
  - verb: owns
    entity: member
references:
- kind: code
  role: implementation
  target: core/api/documents.go#RestoreDocument
- kind: code
  role: implementation
  target: ui/src/routes/DocumentDetail.svelte
- kind: doc
  role: intent
  target: docs/web-app.mdx
---

# Only the owner or an administrator restores a document from Trash

Trash belongs to the document's owner; administrators see and restore every trashed document of the filing system.
