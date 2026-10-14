---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects Documents and picks a Tag under Add tags
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: tag, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product adds the Tag to each selected Document the Member may change
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Tags]}
  - {entity: tag, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Add a tag to selected documents

## Trigger

The person selects documents and chooses Add tags.

## Outcome

The tag is on every selected document the person may change; a partial failure keeps the selection for retry.
