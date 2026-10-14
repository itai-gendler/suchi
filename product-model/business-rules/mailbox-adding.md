---
appliesTo:
- type: entity
  id: mailbox
  effect: creates
permits:
- actors:
  - administrator
- actors:
  - member
  when:
  - entity: member
    fact: Manage mailboxes
    is: 'On'
references:
- kind: code
  role: implementation
  target: core/api/emailaccounts.go#CreateEmailAccount
---

# Only administrators and members with Manage mailboxes add mailboxes

A member's mailbox always imports documents for that member; an administrator chooses the owner.
