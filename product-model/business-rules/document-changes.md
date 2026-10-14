---
appliesTo:
- type: entity
  id: document
  effect: changes
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
    is-not: View
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
    is-not: View
- unattended: true
references:
- kind: code
  role: implementation
  target: core/api/documents.go#PatchDocument
- kind: code
  role: implementation
  target: core/authz/authz.go
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Only the owner, an administrator or someone with Edit access changes a document

A document's owner and administrators change it; anyone else needs Edit or Full control access, granted to them or to one of their groups, and must be able to enter its filing system. Automations, matching and the document model change documents on the archive's own behalf.
