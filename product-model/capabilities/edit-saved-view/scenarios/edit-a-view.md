---
kind: primary
routes:
  web: Web
steps:
- text: The Member changes the name or filters of their Saved view and chooses Save changes
  kind: actor
  actor: member
  entities:
  - {entity: saved-view, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::views
- text: The Product saves the Saved view in place
  kind: product
  actor: member
  entities:
  - {entity: saved-view, effect: changes, facts: [Name, Filter]}
  contexts:
    web:
      place: suchi-web::archive::views
---

# Edit a view

## Trigger

The owner chooses Edit on their view.

## Outcome

The view shows the new name and filters.
