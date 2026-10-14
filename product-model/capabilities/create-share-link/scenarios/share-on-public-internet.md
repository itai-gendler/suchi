---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Share, picks Public internet and an expiry of up to 365 days
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product creates an active public Share link that expires when chosen
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: share-link, effect: creates, to: Active, facts: [Documents, Surface, Expiry, Password, Label, Created]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Share a document on the public internet

## Trigger

The operator enabled public sharing and the person chooses Public internet.

## Outcome

The link uses the public share address and expires after the chosen 1, 7, 30, 90 or 365 days; 7 is the default.
