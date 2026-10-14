---
kind: primary
routes:
  web: Web
steps:
- text: The Member deletes their Saved view and confirms
  kind: actor
  actor: member
  entities:
  - {entity: saved-view, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::views
- text: The Product deletes the Saved view
  kind: product
  actor: member
  entities:
  - {entity: saved-view, effect: removes}
  contexts:
    web:
      place: suchi-web::archive::views
---

# Delete a view

## Trigger

The owner chooses Delete on their view and confirms.

## Outcome

The view is gone for everyone it was shared with.
