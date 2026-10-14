---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator turns off the view-sharing ability for a Member
  kind: actor
  actor: administrator
  entities:
  - {entity: member, effect: reads, facts: [Share saved views]}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product removes the ability and makes the Member's Saved views private
  kind: product
  actor: administrator
  entities:
  - {entity: member, effect: changes, facts: [Share saved views]}
  - {entity: saved-view, effect: changes, facts: [Shared]}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Stop a member sharing views

## Trigger

An administrator switches off Share saved views for a member.

## Outcome

The member's shared views become private; they keep them.
