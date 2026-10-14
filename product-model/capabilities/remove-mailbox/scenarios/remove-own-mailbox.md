---
kind: primary
routes:
  web: Web
steps:
- text: The Member removes their Mailbox and confirms
  kind: actor
  actor: member
  entities:
  - {entity: mailbox, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::my-mailboxes
- text: The Product deletes the Mailbox
  kind: product
  actor: member
  entities:
  - {entity: mailbox, effect: removes}
  contexts:
    web:
      place: suchi-web::archive::my-mailboxes
---

# Remove your own mailbox

## Trigger

A member removes one of their mailboxes and confirms.

## Outcome

The mailbox is no longer polled.
