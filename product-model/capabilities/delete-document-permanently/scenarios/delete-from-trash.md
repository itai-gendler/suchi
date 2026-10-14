---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Delete permanently on a Document in Trash and confirms
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::trash
- text: The Product deletes the Document with everything kept for it
  kind: product
  actor: member
  entities:
  - {entity: document, effect: removes, from: Trashed}
  - {entity: note, effect: removes, with: document}
  - {entity: document-link, effect: removes, with: document}
  - {entity: document-access, effect: removes, with: document}
  contexts:
    web:
      place: suchi-web::archive::trash
---

# Delete a trashed document permanently

## Trigger

The owner chooses Delete permanently and confirms.

## Outcome

The document, its notes, links, access, suggestions and dates are gone and cannot be restored.
