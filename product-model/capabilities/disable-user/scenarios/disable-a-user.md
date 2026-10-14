---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator switches Active off for another person
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product disables the Account
  kind: product
  actor: administrator
  entities:
  - {entity: account, effect: changes, from: Active, to: Disabled, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Disable a user

## Trigger

An administrator switches a person's Active switch off.

## Outcome

The person cannot sign in; their mailboxes stop polling and their share links stop working.
