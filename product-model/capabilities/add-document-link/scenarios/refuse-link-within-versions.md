---
kind: validation
routes:
  web: Web
steps:
- text: The Member chooses the open Document or one of its versions as the target
  kind: actor
  actor: member
  entities:
  - {entity: custom-field, effect: reads, facts: [Name]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product refuses the Document link
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: document-link, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Refuse a link to itself or its own versions

## Trigger

The person chooses the document itself or a version of it as the target.

## Outcome

No link is saved; Versions already relates the versions.
