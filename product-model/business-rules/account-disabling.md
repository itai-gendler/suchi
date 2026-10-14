---
appliesTo:
- type: entity
  id: account
  effect: changes
  to: Disabled
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchUser
---

# Only administrators disable accounts

Nobody disables their own account, and at least one active administrator always remains.
