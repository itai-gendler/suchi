---
kind: validation
routes:
  web: Web
steps:
- text: The Member enters a password for a Locked Document
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Password protection]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product reports that the password opened nothing and keeps the Document Locked
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Refuse a wrong password

## Trigger

The password opens nothing.

## Outcome

The document stays Locked and the person is told the password opened nothing.
