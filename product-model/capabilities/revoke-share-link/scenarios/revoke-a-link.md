---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Revoke on one of their active Share links
  kind: actor
  actor: member
  entities:
  - {entity: share-link, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product revokes the Share link
  kind: product
  actor: member
  entities:
  - {entity: share-link, effect: changes, from: Active, to: Revoked, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Revoke a share link

## Trigger

The creator chooses Revoke on an active link.

## Outcome

The link stops working at once and moves to History.

## Edge cases

- A revoked link cannot be switched to another surface; revoke and create a new one.
