---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator previews a tree file that names a new filing system code and chooses Apply
  kind: actor
  actor: administrator
  entities:
  - {entity: filing-system, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product creates the Filing system with its categories, admitting only administrators
  kind: product
  actor: administrator
  entities:
  - {entity: filing-system, effect: creates, facts: [Code, Name, Members, Filing tree choice]}
  - {entity: category, effect: creates, facts: [Code, Name, Description, Area], with: filing-system}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
---

# Import a tree into a new filing system

## Trigger

An administrator previews a file whose header names a filing system code that does not exist yet.

## Outcome

A new filing system with that code exists with the imported tree and admits administrators only; the first such import instead lets the administrator give the original archive that code.
