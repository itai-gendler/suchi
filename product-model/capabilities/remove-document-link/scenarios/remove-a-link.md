---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Remove link
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product removes the Document link
  kind: product
  actor: member
  entities:
  - {entity: document-link, effect: removes}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Remove a link

## Trigger

The person chooses Remove link.

## Outcome

The link is gone from both documents; its link type is offered again under Add link.
