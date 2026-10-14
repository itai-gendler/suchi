---
appliesTo:
- type: entity
  id: document-access
  effect: changes
permits:
- actors:
  - administrator
- related:
  - verb: has
    entity: document
  - verb: owns
    entity: member
references:
- kind: code
  role: implementation
  target: core/api/acls.go#PutGrant
---

# Only the owner or an administrator changes access to a document

Access granted to someone never lets them change it.
