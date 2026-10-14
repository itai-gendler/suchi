---
appliesTo:
- type: entity
  id: mailbox
  effect: removes
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
references:
- kind: code
  role: implementation
  target: core/api/emailaccounts.go#DeleteEmailAccount
---

# Only the mailbox owner with Manage mailboxes or an administrator removes a mailbox

Documents the mailbox imported stay.
