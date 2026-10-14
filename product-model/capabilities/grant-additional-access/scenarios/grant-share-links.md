---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator turns on the link-sharing ability for a Member
  kind: actor
  actor: administrator
  entities:
  - {entity: member, effect: reads, facts: [Create share links]}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product grants the Member the ability
  kind: product
  actor: administrator
  entities:
  - {entity: member, effect: changes, facts: [Create share links]}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Let a member create share links

## Trigger

An administrator switches on Create share links for a member.

## Outcome

The member can create share links for their own documents; links revoked earlier stay revoked.
