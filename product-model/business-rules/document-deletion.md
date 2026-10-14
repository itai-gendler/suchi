---
appliesTo:
- type: entity
  id: document
  effect: removes
permits:
- actors:
  - administrator
- related:
  - verb: owns
    entity: member
- unattended: true
references:
- kind: code
  role: implementation
  target: core/api/trash.go#PurgeTrashDocument
- kind: code
  role: implementation
  target: core/trash/trash.go#PurgeExpired
---

# Only the owner or an administrator deletes a document permanently

Permanent deletion is confirmed in Trash by the owner or an administrator, or happens by itself 30 days after the document was trashed.
