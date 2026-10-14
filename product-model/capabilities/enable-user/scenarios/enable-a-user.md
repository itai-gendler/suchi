---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator switches Active on for a disabled person
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product enables the Account
  kind: product
  actor: administrator
  entities:
  - {entity: account, effect: changes, from: Disabled, to: Active, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Enable a user

## Trigger

An administrator switches a disabled person's Active switch on.

## Outcome

The person can sign in again with everything still attached.
