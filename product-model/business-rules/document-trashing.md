---
appliesTo:
- type: entity
  id: document
  effect: changes
  to: Trashed
permits:
- actors:
  - administrator
- related:
  - verb: owns
    entity: member
- related:
  - verb: has
    entity: document-access
  - verb: holds
    entity: member
  when:
  - entity: document-access
    fact: Access level
    is: Full control
- related:
  - verb: has
    entity: document-access
  - verb: holds
    entity: group
  - verb: includes
    entity: member
  when:
  - entity: document-access
    fact: Access level
    is: Full control
- unattended: true
references:
- kind: code
  role: implementation
  target: core/api/documents.go#SoftDeleteDocument
---

# Only the owner, an administrator or someone with Full control moves a document to Trash

Moving to Trash needs delete permission: ownership, the administrator role or Full control access. An automation's discard action may also move a document to Trash.
