---
kind: primary
routes:
  web: Web
steps:
- text: The Member opens Inbox
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::inbox
- text: The Product lists the Documents waiting in the System Inbox that the Member may open
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title, Added, Tags, Correspondents, Sensitivity]}
  contexts:
    web:
      place: suchi-web::archive::inbox
---

# Review the Inbox

## Trigger

The person opens Inbox.

## Outcome

The person sees the documents waiting for a filing decision, or Inbox zero.

## Edge cases

- When the Inbox cannot be looked up, no document list is shown instead.
