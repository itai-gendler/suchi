---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Restore on a Document in Trash
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::trash
- text: The Product restores the Document under its category
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, from: Trashed, to: Filed, facts: []}
  - {entity: category, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::trash
---

# Restore a filed document from Trash

## Trigger

The owner chooses Restore in Trash.

## Outcome

The document is live again under its category; links to it show again.
