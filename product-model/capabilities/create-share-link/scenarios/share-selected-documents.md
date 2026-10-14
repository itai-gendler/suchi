---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects Documents they own and chooses Share
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product creates one active Share link for all selected Documents
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: share-link, effect: creates, to: Active, facts: [Documents, Surface, Expiry, Password, Label, Created]}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Share selected documents as one link

## Trigger

The person selects documents and chooses Share.

## Outcome

One link serves every selected document.
