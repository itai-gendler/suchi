---
entities:
- entity: mailbox
  shows:
  - Name
  - Provider
  - Username
  - Source folder
  - Intake rules
  - After import
  - Sync mail from
  - Poll interval
  - Enabled
  - Sync status
references:
- kind: code
  role: implementation
  target: ui/src/lib/EmailAccounts.svelte
---

# Mailboxes

The mailboxes the person owns, under My account, with their sync status.
