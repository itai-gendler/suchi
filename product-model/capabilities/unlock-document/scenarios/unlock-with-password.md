---
kind: primary
routes:
  web: Web
steps:
- text: The Member enters the password for a Locked Document
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Password protection]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product opens the file with the password, keeps an unlocked copy beside the original and reads it
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Password protection, Extracted text]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Unlock one document

## Trigger

The person opens a Locked document and enters its password.

## Outcome

The document shows Password unlocked, has searchable text and can be previewed; the original stays protected.

## Edge cases

- The person may save the password among their saved decryption passwords so later files open by themselves.
