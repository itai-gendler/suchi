---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Let Suchi rename this
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product marks the title automatic without changing it
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Title ownership]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Let Suchi rename this

## Trigger

The person chooses Let Suchi rename this on a title they or someone else set.

## Outcome

The title is automatic again and unchanged for now.
