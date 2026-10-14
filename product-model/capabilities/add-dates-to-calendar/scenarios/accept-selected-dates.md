---
kind: primary
routes:
  web: Web
steps:
- text: The Member checks the date, role and exact text of suggested Document dates and selects them
  kind: actor
  actor: member
  entities:
  - {entity: document-date, effect: reads, facts: [Date, Role, Document text, Model confidence]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Product accepts each selected Document date whose text still matches the document, recording the Member as reviewer
  kind: product
  actor: member
  entities:
  - {entity: document-date, effect: changes, from: Proposed, to: Accepted, facts: [Reviewed by]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Add selected dates to Calendar

## Trigger

The person selects suggested dates in Approvals and chooses Add to Calendar.

## Outcome

The selected dates are on Calendar as reviewed by the person; unselected dates stay pending.

## Edge cases

- A suggestion made from an earlier reading of the document cannot be accepted, only rejected.
