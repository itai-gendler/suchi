---
kind: edge
routes:
  web: Web
steps:
- text: The Member asks a question
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::ask-your-documents
- text: The document model cannot be reached
  kind: condition
  actor: member
  entities:
  - {entity: archive-settings, effect: reads, facts: [Document model]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::ask-your-documents
- text: The Product reports that no answer is available
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::ask-your-documents
---

# Answer nothing when the model fails

## Trigger

The document model cannot be reached.

## Outcome

The person sees that no answer is available; nothing is saved.
