---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator changes a Mailbox and saves
  kind: actor
  actor: administrator
  entities:
  - {entity: mailbox, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::email-intake
- text: The Product saves the Mailbox and reloads its polling
  kind: product
  actor: administrator
  entities:
  - {entity: mailbox, effect: changes, facts: [Intake rules, Poll interval, Enabled]}
  contexts:
    web:
      place: suchi-web::archive-configuration::email-intake
---

# Edit a mailbox under Email intake

## Trigger

An administrator edits a mailbox and saves.

## Outcome

The mailbox polls with the new settings.
