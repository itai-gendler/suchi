---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator edits the Correspondent name
  kind: actor
  actor: administrator
  entities:
  - {entity: correspondent, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::correspondents
- text: The Product renames the Correspondent
  kind: product
  actor: administrator
  entities:
  - {entity: correspondent, effect: changes, facts: [Name]}
  contexts:
    web:
      place: suchi-web::archive-configuration::correspondents
---

# Rename a correspondent

## Trigger

An administrator edits a correspondent name.

## Outcome

Every document with the correspondent shows the new name.
