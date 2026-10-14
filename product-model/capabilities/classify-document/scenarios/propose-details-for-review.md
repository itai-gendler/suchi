---
kind: primary
routes:
  web: Web
steps:
- text: The mode is Review first or a result is below its threshold
  kind: condition
  unattended: true
  entities:
  - {entity: archive-settings, effect: reads, facts: [Application mode, Model confidence threshold, Review match threshold]}
  - {entity: document, effect: reads, facts: [Extracted text]}
- text: The Product creates a Suggestion for each proposed change and a Document date for each found date, waiting in Approvals
  kind: product
  entities:
  - {entity: suggestion, effect: creates, to: Proposed, facts: [Field, Current value, Suggested value, Why suggested, Model confidence, Suggested by]}
  - {entity: document-date, effect: creates, to: Proposed, facts: [Date, Precision, Role, Document text, Model confidence, Reviewed by]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Send details to Approvals

## Trigger

The application mode is Review first, or a result is below its threshold.

## Outcome

Each proposed change waits in Approvals for the document owner, and each found date waits for review.
