---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator selects up to 500 Tags and confirms Delete selected
  kind: actor
  actor: administrator
  entities:
  - {entity: tag, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::tags
- text: The Product deletes the Tags, removes them from every Document and makes their child tags top-level
  kind: product
  actor: administrator
  entities:
  - {entity: tag, effect: removes}
  - {entity: document, effect: changes, facts: [Tags]}
  contexts:
    web:
      place: suchi-web::archive-configuration::tags
---

# Delete selected tags

## Trigger

An administrator selects tags and chooses Delete selected.

## Outcome

The tags are gone from the catalog and from every document; their child tags become top-level. This cannot be undone.
