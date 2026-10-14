---
kind: primary
routes:
  web: Web
steps:
- text: The Account is bound to the identity provider
  kind: condition
  actor: member
  entities:
  - {entity: account, effect: reads, facts: [Sign-in method]}
  contexts:
    web:
      place: suchi-web::archive::profile
- text: The Member chooses Sync from identity provider and signs in at the provider again
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::profile
- text: The Product adopts the provider's verified email for the Account, refusing an address another Account uses
  kind: product
  actor: member
  entities:
  - {entity: account, effect: changes, facts: [Email]}
  contexts:
    web:
      place: suchi-web::archive::profile
---

# Sync email from the identity provider

## Trigger

A person whose account is bound to the identity provider chooses Sync from identity provider.

## Outcome

The account carries the provider's current verified email, or the person learns nothing changed; browser sessions are replaced when it changes.
