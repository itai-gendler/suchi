---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator removes a Mailbox and confirms
  kind: actor
  actor: administrator
  entities:
  - {entity: mailbox, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::email-intake
- text: The Product deletes the Mailbox
  kind: product
  actor: administrator
  entities:
  - {entity: mailbox, effect: removes}
  contexts:
    web:
      place: suchi-web::archive-configuration::email-intake
---

# Remove a mailbox under Email intake

## Trigger

An administrator removes a mailbox and confirms.

## Outcome

The mailbox is no longer polled.
