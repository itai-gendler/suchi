---
entities:
- entity: account
  shows:
  - Email
  - Display name
- entity: member
  shows:
  - Ask your documents
  - Manage document dates
  - Manage mailboxes
  - Create share links
  - Share saved views
references:
- kind: code
  role: implementation
  target: ui/src/routes/PeopleSettings.svelte
---

# Users

Everyone with an account, their role, whether they are active and the additional access of each member.
