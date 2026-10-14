---
kind: primary
routes:
  web: Web
steps:
- text: The Member clears a Custom field value
  kind: actor
  actor: member
  entities:
  - {entity: custom-field, effect: reads, facts: [Name]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product removes the value from the Document
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Custom field values]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Clear a field value

## Trigger

The person clears a shown custom field value.

## Outcome

The field no longer shows on the document and is offered again under Add custom field.
