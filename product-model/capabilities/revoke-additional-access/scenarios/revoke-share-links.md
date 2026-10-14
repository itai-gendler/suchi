---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator turns off the link-sharing ability for a Member
  kind: actor
  actor: administrator
  entities:
  - {entity: member, effect: reads, facts: [Create share links]}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product removes the ability and revokes the Member's active Share links
  kind: product
  actor: administrator
  entities:
  - {entity: member, effect: changes, facts: [Create share links]}
  - {entity: share-link, effect: changes, from: Active, to: Revoked, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Stop a member creating share links

## Trigger

An administrator switches off Create share links for a member.

## Outcome

The member's live share links are revoked and do not return if the ability is granted again.
