---
kind: primary
routes:
  web: Web
steps:
- text: The Member reviews the current and suggested values and why it was suggested, and chooses File document
  kind: actor
  actor: member
  entities:
  - {entity: suggestion, effect: reads, facts: [Current value, Suggested value, Why suggested]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Document and the field are unchanged since the Suggestion was made and the Member may still change it
  kind: condition
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Filing category]}
  - {entity: suggestion, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Product applies the Suggestion and files the Document
  kind: product
  actor: member
  entities:
  - {entity: suggestion, effect: changes, from: Proposed, to: Accepted, facts: []}
  - {entity: document, effect: changes, from: Inbox, to: Filed, facts: [Filing category, Filing address]}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Apply a filing suggestion

## Trigger

The owner reviews a filing suggestion and chooses File document.

## Outcome

The document is filed under the suggested category and the card leaves Approvals.

## Edge cases

- Applying a title suggestion makes the title the person's own.
- Applying a correspondent or tag that does not exist yet also needs the authority to create it.
- Supporting documents the reviewer may not see are never revealed.
