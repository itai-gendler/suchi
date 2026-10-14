---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator chooses or pastes a HuML or TOML tree and chooses Preview
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product validates the file and shows the additions and name collisions
  kind: product
  actor: administrator
  entities:
  - {entity: category, effect: reads, facts: [Code, Name]}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Administrator skips or remaps each collision and chooses Apply
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product adds the categories and records the import for the Filing system
  kind: product
  actor: administrator
  entities:
  - {entity: category, effect: creates, facts: [Code, Name, Description, Area]}
  - {entity: filing-system, effect: changes, facts: [Filing tree choice]}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
---

# Import a tree file

## Trigger

An administrator chooses Import a file and previews it.

## Outcome

The new categories are added; same code and name keep the existing category and its description.

## Edge cases

- Imports never create, change or remove automations.
- Validation checks structure, not whether the tree suits the documents.
