---
kind: primary
routes:
  web: Web
steps:
- text: The Member opens Calendar and picks a range
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::calendar
- text: The Product shows each accepted Document date with its Document and text
  kind: product
  actor: member
  entities:
  - {entity: document-date, effect: reads, facts: [Date, Precision, Role, Document text, Reviewed by]}
  - {entity: document, effect: reads, facts: [Title]}
  contexts:
    web:
      place: suchi-web::archive::calendar
---

# Browse document dates

## Trigger

The person opens Calendar.

## Outcome

The person sees the current month, or the next 30 days or 2 weeks, with each accepted date and how it was accepted.

## Edge cases

- Month- and year-only dates are not placed on a single day.
- Reviewed dates can be exported to another calendar as a local file.
