---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses another category for a filed Document
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: category, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product files the Document under the new category
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Filing category, Filing address]}
  - {entity: category, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Refile a filed document

## Trigger

The person chooses another category for a filed document.

## Outcome

The document is filed under the new category; the old filing address no longer resolves.
