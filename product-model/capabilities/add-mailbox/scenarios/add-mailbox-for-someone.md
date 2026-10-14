---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator chooses the provider, credentials, intake rules and the person who owns what it imports
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::email-intake
- text: The Product saves the Mailbox for that owner and starts polling it
  kind: product
  actor: administrator
  entities:
  - {entity: mailbox, effect: creates, facts: [Name, Provider, Username, Credentials, Source folder, Intake rules, After import, Sync mail from, Poll interval, Enabled]}
  contexts:
    web:
      place: suchi-web::archive-configuration::email-intake
---

# Add a mailbox for a person

## Trigger

An administrator chooses Add mailbox under Email intake.

## Outcome

The mailbox imports documents owned by the chosen person.
