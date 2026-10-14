---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Revoke access
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product removes the Document access
  kind: product
  actor: member
  entities:
  - {entity: document-access, effect: removes}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Revoke access

## Trigger

The owner chooses Revoke access.

## Outcome

The person or group can no longer open the document through that access.

## Edge cases

- Revoking cannot take back what was already downloaded.
