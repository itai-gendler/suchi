---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator enters a new Tag name
  kind: actor
  actor: administrator
  entities:
  - {entity: tag, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::tags
- text: The Product creates the Tag
  kind: product
  actor: administrator
  entities:
  - {entity: tag, effect: creates, facts: [Name, Slug]}
  contexts:
    web:
      place: suchi-web::archive-configuration::tags
---

# Create a tag

## Trigger

An administrator enters a new tag name.

## Outcome

People can put the tag on documents.
