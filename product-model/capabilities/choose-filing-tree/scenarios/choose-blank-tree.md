---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator chooses Blank and confirms
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product records the blank choice for the Filing system
  kind: product
  actor: administrator
  entities:
  - {entity: filing-system, effect: changes, facts: [Filing tree choice]}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
---

# Start from a blank tree

## Trigger

An administrator chooses Blank.

## Outcome

The filing system has no user categories yet, and setup counts as complete.
