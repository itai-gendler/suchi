---
kind: primary
routes:
  web: Web
steps:
- text: The Member picks another level for existing access
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product replaces the Document access level
  kind: product
  actor: member
  entities:
  - {entity: document-access, effect: changes, facts: [Access level]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Change an access level

## Trigger

The owner picks another level for existing access.

## Outcome

The person or group has the new level.
