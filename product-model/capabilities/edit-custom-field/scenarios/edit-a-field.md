---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator renames the Custom field or adds choices and saves
  kind: actor
  actor: administrator
  entities:
  - {entity: custom-field, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
- text: The Product saves the Custom field
  kind: product
  actor: administrator
  entities:
  - {entity: custom-field, effect: changes, facts: [Name, Choices]}
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
---

# Rename a field or add choices

## Trigger

An administrator renames a field or saves new choices.

## Outcome

The field shows its new name or offers the new choices.

## Edge cases

- A saved view or query that names the field by its old name fails with a query error.
