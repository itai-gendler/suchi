---
kind: edge
routes:
  web: Web
steps:
- text: The Member chooses Upload replacement and picks the new file
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: A newer version of the Document exists, live or in Trash
  kind: condition
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Previous version]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product refuses the upload with a conflict
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
---

# Refuse a replacement of an earlier version

## Trigger

The person uploads a replacement on a version that is no longer the latest, or a newer version is in Trash.

## Outcome

Nothing is added; the person gets a conflict that does not reveal a newer version they cannot see.
