---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Change link and picks another Document
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product saves the new target of the Document link
  kind: product
  actor: member
  entities:
  - {entity: document-link, effect: changes, facts: [Target document]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Change the link target

## Trigger

The person chooses Change link on an outgoing link.

## Outcome

The link points to the newly chosen document.
