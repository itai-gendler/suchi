---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses a Custom field and enters a value
  kind: actor
  actor: member
  entities:
  - {entity: custom-field, effect: reads, facts: [Name, Type]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product checks the value against the field type and choices and saves it
  kind: product
  actor: member
  entities:
  - {entity: custom-field, effect: reads, facts: [Choices]}
  - {entity: document, effect: changes, facts: [Custom field values]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Set a field value

## Trigger

The person adds a field under Add custom field or edits a shown value.

## Outcome

The document carries the value, formatted for its type.

## Edge cases

- Invoice fields that extraction manages are not offered for adding but stay editable once extraction filled them.
