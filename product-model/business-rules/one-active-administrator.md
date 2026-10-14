---
appliesTo:
- type: capability
  id: disable-user
- type: context
  context:
    place: suchi-web::archive-configuration::users
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchUser
---

# At least one active administrator always remains, and nobody disables their own account

A change that would leave no active administrator is refused as a whole.
