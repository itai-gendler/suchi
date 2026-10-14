---
appliesTo:
- type: capability
  id: delete-group
- type: capability
  id: revoke-document-access
references:
- kind: code
  role: implementation
  target: core/authz/store.go#DeleteGroup
---

# A group that holds document access cannot be deleted

Its access must be revoked first so nobody loses access by accident.
