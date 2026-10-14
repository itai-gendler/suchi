---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Dismiss on a Suggestion
  kind: actor
  actor: member
  entities:
  - {entity: suggestion, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Product closes the Suggestion without changing the Document
  kind: product
  actor: member
  entities:
  - {entity: suggestion, effect: changes, from: Proposed, to: Dismissed, facts: []}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Dismiss a suggestion

## Trigger

The owner chooses Dismiss on a suggestion card.

## Outcome

The suggestion is closed and the document is unchanged.

## Edge cases

- A later reading of the document can make a fresh suggestion.
