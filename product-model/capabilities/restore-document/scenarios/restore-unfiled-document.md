---
kind: primary
routes:
  web: Web
steps:
- text: The Member opens a trashed Document and chooses Restore
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product restores the Document to the Inbox
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, from: Trashed, to: Inbox, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Restore an unfiled document

## Trigger

The owner opens a trashed document that was never filed and chooses Restore.

## Outcome

The document is back in the Inbox and stays open for editing.

## Edge cases

- After 30 days no restore is offered.
