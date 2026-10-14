---
domain: users
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/setup.go#CreateUser
- kind: code
  role: implementation
  target: ui/src/lib/UserCreateForm.svelte
---

# Create a user

An administrator creates an account with a role and, for a member, the additional access they need. With an identity provider, people create their own accounts by signing in instead.
