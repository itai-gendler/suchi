---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects suggested Document dates and chooses Reject
  kind: actor
  actor: member
  entities:
  - {entity: document-date, effect: reads, facts: [Date, Document text]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Product rejects each selected Document date, recording the Member as reviewer
  kind: product
  actor: member
  entities:
  - {entity: document-date, effect: changes, from: Proposed, to: Rejected, facts: [Reviewed by]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Reject selected dates

## Trigger

The person selects suggested dates and chooses Reject.

## Outcome

The rejected dates never appear on Calendar.
