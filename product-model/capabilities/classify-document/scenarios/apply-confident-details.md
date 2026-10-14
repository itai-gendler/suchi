---
kind: primary
routes:
  web: Web
steps:
- text: Archive matching or the document model returns details at or above their thresholds while the mode is Apply when confident
  kind: condition
  unattended: true
  entities:
  - {entity: archive-settings, effect: reads, facts: [Application mode, Model confidence threshold, Automatic match threshold]}
  - {entity: document, effect: reads, facts: [Extracted text, Similar documents]}
- text: The Product files and names the Document with the confident details and adds its confident dates to Calendar
  kind: product
  entities:
  - {entity: document, effect: changes, from: Inbox, to: Filed, facts: [Filing category, Filing address, Title, Correspondents, Tags, Languages]}
  - {entity: document-date, effect: creates, to: Accepted, facts: [Date, Precision, Role, Document text, Model confidence, Reviewed by]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Apply confident details

## Trigger

A newly read document is still in the Inbox and the application mode is Apply when confident.

## Outcome

Results at or above their thresholds are applied: the document is filed and named, and its confident dates are on Calendar.

## Edge cases

- Text goes to a hosted model only after an administrator acknowledged it.
- Model filing applies only while the document is in the Inbox, and a correspondent only while none is set.
- A dated result is kept only when its exact text appears in the document.
