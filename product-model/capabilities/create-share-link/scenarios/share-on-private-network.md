---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Share, keeps Private network and picks an expiry and optional password
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product creates an active Share link for the Document and shows it
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: share-link, effect: creates, to: Active, facts: [Documents, Surface, Expiry, Password, Label, Created]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Share a document on the private network

## Trigger

The person chooses Share on a document and keeps Private network.

## Outcome

The person sees the link, can copy it or show it as a QR code, and the document lists it among its active links.

## Edge cases

- A password is recommended for Confidential or Restricted documents but never added by itself.
