---
kind: primary
routes:
  web: Web
steps:
- text: The Member turns on Share this view and saves
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::views
- text: The Product shares the Saved view in the filing system
  kind: product
  actor: member
  entities:
  - {entity: saved-view, effect: changes, facts: [Shared]}
  - {entity: filing-system, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::views
---

# Share a view

## Trigger

The owner turns on Share this view.

## Outcome

Everyone who may enter the filing system sees the view, and still only the documents they may open.
