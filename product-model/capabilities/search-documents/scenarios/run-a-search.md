---
kind: primary
routes:
  web: Web
steps:
- text: 'The Member enters a search, with query terms such as from: or has-field: if wanted'
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::search
- text: The Product ranks the matching Documents the Member may open
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title, Filing address, Extracted text]}
  contexts:
    web:
      place: suchi-web::archive::search
---

# Search everything

## Trigger

The person types into search and submits.

## Outcome

The person sees ranked matching documents with snippets, counted only over documents they may open.
