---
kind: validation
routes:
  web: Web
steps:
- text: The Member enters a value that does not fit the Custom field
  kind: actor
  actor: member
  entities:
  - {entity: custom-field, effect: reads, facts: [Name, Type]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product refuses the value
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
---

# Refuse an invalid value

## Trigger

The value does not fit the field type or its choices.

## Outcome

Nothing changes and the editor shows the problem.
