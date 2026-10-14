---
kind: primary
routes:
  web: Web
steps:
- text: The Member deletes their Note and confirms
  kind: actor
  actor: member
  entities:
  - {entity: note, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product deletes the Note
  kind: product
  actor: member
  entities:
  - {entity: note, effect: removes}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Delete a note

## Trigger

The author chooses Delete note and confirms.

## Outcome

The note is permanently gone from the version.
