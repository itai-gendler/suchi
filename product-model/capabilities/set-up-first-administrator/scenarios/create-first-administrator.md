---
kind: primary
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
- text: The token is the current one and has not expired, and no administrator exists yet
  kind: condition
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::setup
- text: The Product creates the Administrator with an active Account, signs them in and asks them to choose a filing tree
  kind: product
  actor: administrator
  entities:
  - {entity: administrator, effect: creates, facts: []}
  - {entity: account, effect: creates, to: Active, facts: [Email, Display name, Password, Sign-in method], with: administrator}
  contexts:
    web:
      place: suchi-web::sign-in::setup
---

# Create the first administrator with the setup token

## Trigger

Someone opens a new server for the first time.

## Outcome

The first Administrator exists, is signed in and is asked to choose a filing tree; the setup token no longer works.

## Edge cases

- A password shorter than 8 characters is refused and nothing is created.
