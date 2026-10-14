---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator switches off Manage mailboxes for a Member
  kind: actor
  actor: administrator
  entities:
  - {entity: member, effect: reads, facts: [Manage mailboxes]}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product removes the ability and disables the Member's Mailboxes
  kind: product
  actor: administrator
  entities:
  - {entity: member, effect: changes, facts: [Manage mailboxes]}
  - {entity: mailbox, effect: changes, facts: [Enabled]}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Stop a member managing mailboxes

## Trigger

An administrator switches off Manage mailboxes for a member.

## Outcome

The member's mailboxes stop polling and stay disabled even if the ability is granted again.
