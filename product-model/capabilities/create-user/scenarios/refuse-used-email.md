---
kind: validation
routes:
  web: Web
steps:
- text: The Administrator enters an email another Account uses
  kind: actor
  actor: administrator
  entities:
  - {entity: account, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product refuses the new user
  kind: product
  actor: administrator
  entities:
  - {entity: account, effect: reads, facts: [Email]}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Refuse an email already in use

## Trigger

The email belongs to another account.

## Outcome

No account is created.
