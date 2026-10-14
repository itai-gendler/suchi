---
kind: primary
routes:
  web: Web
steps:
- text: A Document has been in Trash for 30 days
  kind: condition
  unattended: true
  entities:
  - {entity: document, effect: reads, facts: []}
- text: The Product deletes the Document with everything kept for it
  kind: product
  entities:
  - {entity: document, effect: removes, from: Trashed}
  - {entity: note, effect: removes, with: document}
  - {entity: document-link, effect: removes, with: document}
  - {entity: document-access, effect: removes, with: document}
  contexts:
    web:
      place: suchi-web::archive::trash
---

# Delete a document after 30 days in Trash

## Trigger

A document has been in Trash for 30 days.

## Outcome

The document is permanently deleted without anyone confirming.
