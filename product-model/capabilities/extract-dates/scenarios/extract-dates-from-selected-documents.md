---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects Documents and chooses Extract dates
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product queues date reading for each selected Document the Member may read
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product keeps each date the document model finds with its exact text as a Proposed Document date
  kind: product
  actor: member
  entities:
  - {entity: document-date, effect: creates, to: Proposed, facts: [Date, Precision, Role, Document text, Model confidence, Reviewed by]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Extract dates from selected documents

## Trigger

The person selects documents and chooses Extract dates.

## Outcome

Dates found in the documents' text are added to Calendar or wait in Approvals, as the application mode says.

## Edge cases

- A date whose text does not appear exactly in the document is discarded.
- With Apply when confident, a date at or above the model confidence threshold goes to Calendar at once, as when a document is first read.
