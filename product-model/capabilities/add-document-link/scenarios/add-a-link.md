---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses an unused link type and searches for the target by title, number or address
  kind: actor
  actor: member
  entities:
  - {entity: custom-field, effect: reads, facts: [Name]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product saves a Document link from the open Document to the chosen one
  kind: product
  actor: member
  entities:
  - {entity: document-link, effect: creates, facts: [Link type, Target document]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Link to another document

## Trigger

The person chooses Add link under Linked documents.

## Outcome

The open document links to the chosen document, and the chosen document shows the link back.

## Edge cases

- A pasted link to another server is rejected without being fetched.
- Linking across filing systems needs entry to both.
