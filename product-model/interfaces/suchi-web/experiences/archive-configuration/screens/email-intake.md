---
entities:
- entity: mailbox
  shows:
  - Name
  - Provider
  - Username
  - Intake rules
  - Poll interval
  - Enabled
  - Sync status
references:
- kind: code
  role: implementation
  target: ui/src/lib/EmailAccounts.svelte
---

# Email intake

Every mailbox of the current filing system with its owner and sync status.
