---
kind: primary
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
- text: The Account signs in with a password, the password matches and the address is unused
  kind: condition
  actor: member
  entities:
  - {entity: account, effect: reads, facts: [Sign-in method, Password, Email]}
  contexts:
    web:
      place: suchi-web::archive::profile
- text: The Product changes the Account email immediately and replaces every browser session with one for this browser
  kind: product
  actor: member
  entities:
  - {entity: account, effect: changes, facts: [Email]}
  contexts:
    web:
      place: suchi-web::archive::profile
---

# Change email with the current password

## Trigger

A local-password account holder chooses Change email.

## Outcome

The account signs in with the new address; every other browser session is signed out and this browser stays signed in.

## Edge cases

- Suchi sends no confirmation or notice to either address.
