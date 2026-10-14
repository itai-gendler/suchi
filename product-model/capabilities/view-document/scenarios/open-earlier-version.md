---
kind: primary
routes:
  web: Web
steps:
- text: The Member opens an earlier version from Versions or its address
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product shows that exact version and points to the latest visible version
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title, Previous version]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Open an earlier version

## Trigger

The person opens a version that has been replaced.

## Outcome

The exact earlier version opens, labelled Earlier, with the way to the latest version the person may see.
