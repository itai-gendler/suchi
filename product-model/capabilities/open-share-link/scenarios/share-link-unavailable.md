---
kind: edge
routes:
  web: Web
steps:
- text: The Link recipient opens the Share link
  kind: actor
  actor: link-recipient
  entities:
  - {entity: share-link, effect: reads, facts: []}
  contexts:
    web:
      place: share-page::shared-documents
- text: The link is revoked or expired, or its creator can no longer share from its filing system
  kind: condition
  actor: link-recipient
  entities:
  - {entity: share-link, effect: reads, facts: [Expiry]}
  - {entity: filing-system, effect: reads, facts: []}
  contexts:
    web:
      place: share-page::shared-documents
- text: The Product shows that the link has expired
  kind: product
  actor: link-recipient
  entities: []
  contexts:
    web:
      place: share-page::shared-documents
---

# Show an unavailable share link

## Trigger

A recipient opens a link that was revoked or has expired, or whose creator was disabled or lost access.

## Outcome

The page says the link has expired; nothing is served.
