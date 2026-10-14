---
kind: primary
routes:
  web: Web
steps:
- text: The Member picks an existing Tag
  kind: actor
  actor: member
  entities:
  - {entity: tag, effect: reads, facts: [Name, Slug]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product adds the Tag to the Document
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Tags]}
  - {entity: tag, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Add a tag on the document

## Trigger

The person chooses Edit beside Tags and picks a tag.

## Outcome

The tag is on the document at once.

## Edge cases

- Adding needs-review by hand makes the marker the person's own, so later edits do not remove it.
