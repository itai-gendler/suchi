---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses to sign in and authenticates at the identity provider
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::login
- text: The verified identity is bound to an active Account
  kind: condition
  actor: member
  entities:
  - {entity: account, effect: reads, facts: [Email, Sign-in method]}
  contexts:
    web:
      place: suchi-web::sign-in::login
- text: The Product starts a session and opens the archive
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::login
---

# Sign in through the identity provider

## Trigger

A person whose account is bound to the identity provider opens the archive.

## Outcome

The person has a session and sees the archive.
