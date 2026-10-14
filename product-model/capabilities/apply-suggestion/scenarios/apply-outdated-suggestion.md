---
kind: edge
routes:
  web: Web
steps:
- text: The Member chooses to apply a Suggestion
  kind: actor
  actor: member
  entities:
  - {entity: suggestion, effect: reads, facts: [Suggested value]}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Document or the suggested field changed after the Suggestion was made
  kind: condition
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: suggestion, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Product closes the Suggestion as outdated without changing the Document
  kind: product
  actor: member
  entities:
  - {entity: suggestion, effect: changes, from: Proposed, to: Outdated, facts: []}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Close an outdated suggestion

## Trigger

The owner applies a suggestion after the document or the field changed.

## Outcome

Nothing is applied; the suggestion closes as outdated and leaves Approvals.
