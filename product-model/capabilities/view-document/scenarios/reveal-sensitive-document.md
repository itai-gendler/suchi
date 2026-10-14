---
kind: primary
routes:
  web: Web
steps:
- text: The Member opens a Confidential or Restricted Document
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Sensitivity]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product hides the preview and extracted text until the Member chooses Reveal
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Member chooses Reveal
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product shows the preview and extracted text
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Extracted text]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Reveal a sensitive document

## Trigger

The person opens a Confidential or Restricted document.

## Outcome

The preview and text stay hidden until the person chooses Reveal.
