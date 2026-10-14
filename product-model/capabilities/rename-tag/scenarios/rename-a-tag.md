---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator edits the Tag name
  kind: actor
  actor: administrator
  entities:
  - {entity: tag, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::tags
- text: The Product renames the Tag
  kind: product
  actor: administrator
  entities:
  - {entity: tag, effect: changes, facts: [Name, Slug]}
  contexts:
    web:
      place: suchi-web::archive-configuration::tags
---

# Rename a tag

## Trigger

An administrator edits a tag name.

## Outcome

Every document carrying the tag shows the new name.
