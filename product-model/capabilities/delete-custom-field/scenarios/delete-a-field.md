---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator deletes the Custom field and confirms
  kind: actor
  actor: administrator
  entities:
  - {entity: custom-field, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
- text: The Product deletes the Custom field and its values on every Document
  kind: product
  actor: administrator
  entities:
  - {entity: custom-field, effect: removes}
  - {entity: document, effect: changes, facts: [Custom field values]}
  contexts:
    web:
      place: suchi-web::archive-configuration::custom-fields
---

# Delete a custom field

## Trigger

An administrator deletes a field and confirms.

## Outcome

The field and its values are gone from every document.
