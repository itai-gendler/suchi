---
kind: edge
routes:
  web: Web
steps:
- text: The document model cannot be reached or its answer is invalid
  kind: condition
  unattended: true
  entities:
  - {entity: archive-settings, effect: reads, facts: [Document model]}
  - {entity: document, effect: reads, facts: []}
- text: The Product leaves the Document unchanged in the Inbox
  kind: product
  entities:
  - {entity: document, effect: reads, facts: [Filing category]}
  contexts:
    web:
      place: suchi-web::archive::inbox
---

# Keep the document unchanged when the model fails

## Trigger

The document model cannot be reached or returns an invalid answer.

## Outcome

The document stays in the Inbox with its details unchanged and its review marker kept; the step is retried.
