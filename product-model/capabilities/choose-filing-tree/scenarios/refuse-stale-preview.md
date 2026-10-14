---
kind: edge
routes:
  web: Web
steps:
- text: The Administrator chooses Apply on an older preview
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The filing tree or its documents changed since the preview
  kind: condition
  actor: administrator
  entities:
  - {entity: category, effect: reads, facts: []}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product refuses the apply and asks for a new preview
  kind: product
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
---

# Refuse an apply after the archive changed

## Trigger

Relevant categories or documents changed after the preview.

## Outcome

Nothing is applied; the input is kept and the administrator previews again.
