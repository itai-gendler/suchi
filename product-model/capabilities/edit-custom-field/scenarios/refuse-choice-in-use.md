---
kind: validation
routes:
  web: Web
steps:
- text: The Administrator removes a choice that is in use
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
- text: The Product refuses the change to the Custom field because Documents use the choice
  kind: product
  actor: administrator
  entities:
  - {entity: document, effect: reads, facts: [Custom field values]}
  - {entity: custom-field, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
---

# Refuse removing a choice in use

## Trigger

An administrator removes a choice that documents use, or changes the type of a field with values.

## Outcome

Nothing changes.
