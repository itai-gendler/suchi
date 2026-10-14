---
kind: validation
routes:
  web: Web
steps:
- text: The Member enters their email and password
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::login
- text: The password does not match the Account
  kind: condition
  actor: member
  entities:
  - {entity: account, effect: reads, facts: [Password]}
  contexts:
    web:
      place: suchi-web::sign-in::login
- text: The Product refuses the sign-in without starting a session
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::login
---

# Refuse a wrong password

## Trigger

A person enters credentials that do not match.

## Outcome

No session is started and the sign-in page says the credentials were not accepted.
