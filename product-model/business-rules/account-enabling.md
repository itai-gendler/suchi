---
appliesTo:
- type: entity
  id: account
  effect: changes
  to: Active
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchUser
---

# Only administrators enable accounts

A disabled person cannot enable themselves.
