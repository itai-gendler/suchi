---
kind: validation
routes:
  web: Web
steps:
- text: The Member saves a new view
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::views
- text: The Member already keeps 50 Saved views in this filing system
  kind: condition
  actor: member
  entities:
  - {entity: saved-view, effect: reads, facts: []}
  - {entity: filing-system, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::views
- text: The Product refuses the Saved view
  kind: product
  actor: member
  entities:
  - {entity: saved-view, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::views
---

# Refuse a view over the limit

## Trigger

The person already keeps 50 views in this filing system.

## Outcome

No view is saved and the builder keeps the draft with the error.
