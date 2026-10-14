---
kind: primary
routes:
  web: Web
steps:
- text: The Member names the view and adds filters for text, category, tag, correspondent, custom field presence, sensitivity or dates
  kind: actor
  actor: member
  entities:
  - {entity: tag, effect: reads, facts: []}
  - {entity: correspondent, effect: reads, facts: []}
  - {entity: custom-field, effect: reads, facts: []}
  - {entity: category, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::views
- text: The Product saves the Saved view for the Member
  kind: product
  actor: member
  entities:
  - {entity: saved-view, effect: creates, facts: [Name, Filter, Shared]}
  contexts:
    web:
      place: suchi-web::archive::views
---

# Create a view

## Trigger

The person chooses New view or a starter pattern.

## Outcome

The view is listed with its filter sentence and current document count.

## Edge cases

- Choosing Share this view requires Share saved views.
- A view from Ask your documents keeps the exact documents of the answer.
