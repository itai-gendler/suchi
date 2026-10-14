---
kind: validation
routes:
  web: Web
steps:
- text: The Member chooses a person who may not enter the Filing system
  kind: actor
  actor: member
  entities:
  - {entity: filing-system, effect: reads, facts: [Members]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product refuses the Document access
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: document-access, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Refuse access for someone outside the filing system

## Trigger

The chosen person may not enter the document's filing system or is disabled.

## Outcome

No access is recorded.
