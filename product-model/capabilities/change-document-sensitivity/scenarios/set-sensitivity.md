---
kind: primary
routes:
  web: Web
steps:
- text: The Member picks a sensitivity
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product saves the Document sensitivity
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Sensitivity]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Set the sensitivity of a document

## Trigger

The person picks a sensitivity on the document.

## Outcome

The document carries the new sensitivity everywhere it is shown.

## Edge cases

- Sensitivity changes no access: it never grants or removes anyone's access to the document.
