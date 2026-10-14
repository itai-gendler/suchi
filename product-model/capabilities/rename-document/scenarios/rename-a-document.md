---
kind: primary
routes:
  web: Web
steps:
- text: The Member edits the Document title
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product saves the title as set by a person and closes any pending title Suggestion as outdated
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Title, Title ownership]}
  - {entity: suggestion, effect: changes, from: Proposed, to: Outdated, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Rename a document

## Trigger

The person chooses Rename on the document title.

## Outcome

The document carries the new title, shows that suggestions and rules will not change it, and any pending title suggestion is closed as outdated.
