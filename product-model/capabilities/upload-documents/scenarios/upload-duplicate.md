---
kind: edge
routes:
  web: Web
steps:
- text: The Member chooses or drops a file to upload
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::upload-queue
- text: The Member already owns a live Document with identical bytes in this filing system
  kind: condition
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Original file]}
  - {entity: filing-system, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::upload-queue
- text: The Product returns the existing Document and records where the file came from this time, without reading it again
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Added from]}
  contexts:
    web:
      place: suchi-web::archive::upload-queue
---

# Upload a file the person already has

## Trigger

The person uploads a file identical to a live document they already own in this filing system.

## Outcome

No second document is created; the existing document is returned and records the new source.
