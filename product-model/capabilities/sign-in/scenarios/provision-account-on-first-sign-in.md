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
- text: No Account is bound to the identity and its verified email is unused
  kind: condition
  actor: member
  entities:
  - {entity: account, effect: reads, facts: [Email, Sign-in method]}
  contexts:
    web:
      place: suchi-web::sign-in::login
- text: The Product creates a Member with an active Account bound to that identity and signs them in
  kind: product
  actor: member
  entities:
  - {entity: member, effect: creates, facts: []}
  - {entity: account, effect: creates, to: Active, facts: [Email, Display name, Sign-in method], with: member}
  contexts:
    web:
      place: suchi-web::sign-in::login
---

# Create an account on first identity-provider sign-in

## Trigger

A person signs in through the identity provider for the first time.

## Outcome

A new member account bound to the provider identity exists and is signed in; if its email is the operator's named administrator email and no administrator exists yet, it is an administrator instead.

## Edge cases

- A verified email that already belongs to another account is refused rather than merged.
