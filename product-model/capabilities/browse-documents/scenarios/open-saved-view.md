---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects a Saved view
  kind: actor
  actor: member
  entities:
  - {entity: saved-view, effect: reads, facts: [Name]}
  contexts:
    web:
      place: suchi-web::archive::views
- text: The Product lists the Documents matching the Saved view that the Member may open
  kind: product
  actor: member
  entities:
  - {entity: saved-view, effect: reads, facts: [Filter]}
  - {entity: document, effect: reads, facts: [Title]}
  contexts:
    web:
      place: suchi-web::archive::views
---

# Open a saved view

## Trigger

The person selects a saved view in Views.

## Outcome

The documents matching the view's filter, or its exact documents, are listed.

## Edge cases

- A view whose custom field was renamed or deleted fails with a query error rather than showing more documents.
