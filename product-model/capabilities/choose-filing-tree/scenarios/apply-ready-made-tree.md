---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator picks a ready-made tree or builds one and chooses Preview
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product shows the complete target tree, every name collision and how many documents use each affected category
  kind: product
  actor: administrator
  entities:
  - {entity: category, effect: reads, facts: [Code, Name]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Administrator decides each collision and chooses Apply
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product adds and updates the categories as previewed and records the choice for the Filing system
  kind: product
  actor: administrator
  entities:
  - {entity: category, effect: creates, facts: [Code, Name, Description, Area]}
  - {entity: filing-system, effect: changes, facts: [Filing tree choice]}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
---

# Apply a ready-made or composed tree

## Trigger

An administrator opens Filing tree and picks a tree.

## Outcome

The filing system has the chosen categories; existing documents keep their categories, and the setup reminder disappears.

## Edge cases

- A same-meaning built-in category is updated in place and keeps its documents.
- The System area 40–49 with its Inbox is generated and never replaced.
- Refiling existing documents with the new tree is a separate, explicit choice.
