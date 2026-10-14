---
kind: primary
routes:
  web: Web
steps:
- text: The Member opens Access and chooses a person or Group with View, Edit or Full control
  kind: actor
  actor: member
  entities:
  - {entity: group, effect: reads, facts: [Name]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product records the Document access
  kind: product
  actor: member
  entities:
  - {entity: document-access, effect: creates, facts: [Person or group, Access level]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Grant access to a person or group

## Trigger

The owner opens Access and chooses a person or group with a level.

## Outcome

That person, or everyone in the group, may open the document at the chosen level.

## Edge cases

- Access never lets anyone into a filing system they may not enter.
