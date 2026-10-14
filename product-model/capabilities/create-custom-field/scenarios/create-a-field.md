---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator enters a name, a type and any choices
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
- text: The Product creates the Custom field
  kind: product
  actor: administrator
  entities:
  - {entity: custom-field, effect: creates, facts: [Name, Type, Choices]}
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
---

# Create a custom field

## Trigger

An administrator enters a name and type, and choices for a select or multi-choice field.

## Outcome

The field can be set on documents and used in searches, views and automations.
