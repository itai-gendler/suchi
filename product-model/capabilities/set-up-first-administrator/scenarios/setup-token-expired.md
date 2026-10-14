---
kind: validation
routes:
  web: Web
steps:
- text: The Administrator enters the setup token with their email, display name and password
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::setup
- text: The token has expired
  kind: condition
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::setup
- text: The Product refuses the setup and writes a replacement token to the server log
  kind: product
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::setup
---

# Refuse an expired setup token

## Trigger

Someone enters a setup token more than 24 hours after it was issued.

## Outcome

No account is created; a replacement token is written to the server log for the operator.
