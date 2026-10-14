---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator enters a name of at most 80 characters and saves
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-system
- text: The Product saves the Filing system name
  kind: product
  actor: administrator
  entities:
  - {entity: filing-system, effect: changes, facts: [Name]}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-system
---

# Rename a filing system

## Trigger

An administrator edits the system name and saves.

## Outcome

The new name shows beside the code everywhere.
