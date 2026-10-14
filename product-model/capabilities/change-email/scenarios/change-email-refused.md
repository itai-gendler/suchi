---
kind: validation
routes:
  web: Web
steps:
- text: The Member enters the new address twice and their current password
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::profile
- text: The password is wrong or the address belongs to another Account
  kind: condition
  actor: member
  entities:
  - {entity: account, effect: reads, facts: [Password, Email]}
  contexts:
    web:
      place: suchi-web::archive::profile
- text: The Product refuses the change and keeps the current email
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::profile
---

# Refuse an email change

## Trigger

The person enters a wrong current password or an address that is already in use.

## Outcome

Nothing changes and the dialog explains why.
