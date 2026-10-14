---
appliesTo:
- type: entity
  id: member
  effect: changes
  facts:
  - Ask your documents
  - Manage document dates
  - Manage mailboxes
  - Create share links
  - Share saved views
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchUser
---

# Only administrators change a member's additional access

Administrators have every additional access implicitly.
