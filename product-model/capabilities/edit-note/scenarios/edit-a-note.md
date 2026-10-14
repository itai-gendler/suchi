---
kind: primary
routes:
  web: Web
steps:
- text: The Member edits the text of their Note
  kind: actor
  actor: member
  entities:
  - {entity: note, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product saves the Note text and edited time
  kind: product
  actor: member
  entities:
  - {entity: note, effect: changes, facts: [Text, Edited]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Edit a note

## Trigger

The author chooses Edit on their note.

## Outcome

The note shows the new text and its edited time.
