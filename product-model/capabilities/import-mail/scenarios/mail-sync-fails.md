---
kind: edge
routes:
  web: Web
steps:
- text: A poll of an enabled Mailbox fails
  kind: condition
  unattended: true
  entities:
  - {entity: mailbox, effect: reads, facts: [Credentials, Source folder]}
- text: The Product records the error as the Mailbox sync status
  kind: product
  entities:
  - {entity: mailbox, effect: changes, facts: [Sync status]}
  contexts:
    web:
      place: suchi-web::archive::my-mailboxes
---

# Record a failed mail poll

## Trigger

A poll cannot connect, authenticate or download.

## Outcome

The mailbox shows the error; messages imported before the failure stay, and the next poll resumes from the last completed batch.
