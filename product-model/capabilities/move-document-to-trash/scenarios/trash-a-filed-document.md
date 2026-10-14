---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Trash and confirms
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product moves the Document to Trash
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, from: Filed, to: Trashed, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Move a document to Trash

## Trigger

The person chooses Trash on a document and confirms.

## Outcome

The document is in Trash, read-only, with its deletion date; it disappears from lists, search, links and share links.
