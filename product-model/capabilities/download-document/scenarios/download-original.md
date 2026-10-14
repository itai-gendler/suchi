---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Download on a Document
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product delivers the original file
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Original file]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Download the original

## Trigger

The person chooses Download.

## Outcome

The browser receives the unchanged original file.
