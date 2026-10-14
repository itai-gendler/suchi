---
appliesTo:
- type: entity
  id: document-access
  effect: removes
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
  target: core/api/acls.go#DeleteGrant
---

# Only the owner or an administrator revokes access to a document

Access granted to someone never lets them revoke it.
