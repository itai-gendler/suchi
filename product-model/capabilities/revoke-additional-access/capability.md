---
domain: users
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchUser
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Revoke additional access

An administrator turns off one of a member's Additional access abilities; what that ability opened closes with it.
