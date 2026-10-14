---
kind: edge
routes:
  web: Web
steps:
- text: The Document is password-protected and no saved password opens it
  kind: condition
  unattended: true
  entities:
  - {entity: document, effect: reads, facts: [Original file]}
- text: The Product marks the Document Locked so it can be unlocked
  kind: product
  entities:
  - {entity: document, effect: changes, facts: [Password protection]}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Hold a password-protected document

## Trigger

The added file is password-protected and no saved password opens it.

## Outcome

The document shows as Locked until its owner unlocks it.
