---
kind: primary
routes:
  web: Web
steps:
- text: The Member removes a Tag from the Document
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: tag, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product removes the Tag from the Document
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Tags]}
  - {entity: tag, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Remove a tag from the document

## Trigger

The person chooses Edit beside Tags and removes a tag.

## Outcome

The tag is no longer on the document.
