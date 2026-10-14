---
kind: primary
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
- text: The Account is active and the password matches
  kind: condition
  actor: member
  entities:
  - {entity: account, effect: reads, facts: [Email, Password]}
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

# Sign in with email and password

## Trigger

A person with a local-password account opens the archive.

## Outcome

The person has a session and sees the archive.

## Edge cases

- A disabled Account cannot sign in.
- Without membership in any filing system the person sees only their account and a request to ask an administrator for access.
