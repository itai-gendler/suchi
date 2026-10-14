---
kind: primary
routes:
  web: Web
steps:
- text: The Member picks an existing Correspondent as the sender and saves
  kind: actor
  actor: member
  entities:
  - {entity: correspondent, effect: reads, facts: [Name]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product makes it the Document's primary sender
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Correspondents]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Set the sender of a document

## Trigger

The person chooses Edit beside Correspondents and picks an existing correspondent.

## Outcome

The chosen correspondent is the document's sender; its other correspondents keep their roles.
