---
appliesTo:
- type: entity
  id: mailbox
  effect: changes
permits:
- actors:
  - administrator
- related:
  - verb: owns
    entity: member
  when:
  - entity: member
    fact: Manage mailboxes
    is: 'On'
- unattended: true
references:
- kind: code
  role: implementation
  target: core/api/emailaccounts.go#PatchEmailAccount
---

# Only the mailbox owner with Manage mailboxes or an administrator changes a mailbox

Polling records the sync status on the archive's own behalf.
