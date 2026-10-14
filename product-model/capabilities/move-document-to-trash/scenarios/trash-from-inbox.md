---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Trash on a Document in the Inbox list and confirms
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product moves the Document to Trash
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, from: Inbox, to: Trashed, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Move an Inbox document to Trash

## Trigger

The person uses the Trash quick action on a document in the Inbox list.

## Outcome

The document leaves the Inbox for Trash.
