---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator chooses Export and a format
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product produces a file of the filing tree
  kind: product
  actor: administrator
  entities:
  - {entity: category, effect: reads, facts: [Code, Name, Description, Area]}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
---

# Export the tree

## Trigger

An administrator chooses Export.

## Outcome

The administrator has a HuML or TOML file of the tree that can be imported elsewhere; it is not a backup.
