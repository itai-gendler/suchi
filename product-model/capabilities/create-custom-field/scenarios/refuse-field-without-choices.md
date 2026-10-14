---
kind: validation
routes:
  web: Web
steps:
- text: The Administrator creates a select Custom field with no choices
  kind: actor
  actor: administrator
  entities:
  - {entity: custom-field, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
- text: The Product refuses the Custom field
  kind: product
  actor: administrator
  entities:
  - {entity: custom-field, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
---

# Refuse a choice field without choices

## Trigger

A select or multi-choice field has no choices, or repeats one.

## Outcome

Nothing is created.
