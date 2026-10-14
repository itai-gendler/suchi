---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator opens the Dashboard
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive::dashboard
- text: No filing tree has been chosen for the Filing system
  kind: condition
  actor: administrator
  entities:
  - {entity: filing-system, effect: reads, facts: [Filing tree choice]}
  contexts:
    web:
      place: suchi-web::archive::dashboard
- text: The Product shows Choose your filing tree with Continue setup
  kind: product
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive::dashboard
---

# Ask for a filing tree before first use

## Trigger

An administrator opens the Dashboard of a filing system without a chosen filing tree.

## Outcome

The Dashboard shows Choose your filing tree with Continue setup, which stays until a tree is chosen.
