---
availability:
- {place: 'suchi-web::archive'}
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/emailaccounts.go#DeleteEmailAccount
- kind: code
  role: implementation
  target: ui/src/lib/EmailAccounts.svelte
---

# Remove a mailbox

The mailbox owner, or an administrator, removes a mailbox. Documents it imported stay and keep its name as their source.
