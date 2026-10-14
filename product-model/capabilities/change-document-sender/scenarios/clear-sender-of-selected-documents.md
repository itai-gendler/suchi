---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects Documents and chooses Clear sender
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product removes the sender role from each selected Document the Member may change
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Correspondents]}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Clear the sender of selected documents

## Trigger

The person selects documents and chooses Clear sender.

## Outcome

The selected documents have no sender; other correspondent roles stay.
