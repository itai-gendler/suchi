---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator deletes the Correspondent and confirms
  kind: actor
  actor: administrator
  entities:
  - {entity: correspondent, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::correspondents
- text: The Product deletes the Correspondent and removes it from every Document
  kind: product
  actor: administrator
  entities:
  - {entity: correspondent, effect: removes}
  - {entity: document, effect: changes, facts: [Correspondents]}
  contexts:
    web:
      place: suchi-web::archive-configuration::correspondents
---

# Delete a correspondent

## Trigger

An administrator deletes a correspondent and confirms.

## Outcome

The correspondent is gone from the catalog and from every document.
