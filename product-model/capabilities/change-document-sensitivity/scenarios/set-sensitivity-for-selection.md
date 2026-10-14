---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects Documents and chooses a sensitivity
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product saves the sensitivity on each selected Document the Member may change
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Sensitivity]}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Set the sensitivity of selected documents

## Trigger

The person selects documents and chooses Sensitivity.

## Outcome

Every selected document the person may change carries the chosen sensitivity.
