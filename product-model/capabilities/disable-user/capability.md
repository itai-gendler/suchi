---
domain: users
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchUser
- kind: code
  role: implementation
  target: ui/src/routes/PeopleSettings.svelte
---

# Disable a user

An administrator stops a person signing in without deleting anything they own.
