---
kind: primary
routes:
  web: Web
steps:
- text: The Member writes a note of up to 16 KiB and adds it
  kind: actor
  actor: member
  entities:
  - {entity: note, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product keeps the Note on this version with its author
  kind: product
  actor: member
  entities:
  - {entity: note, effect: creates, facts: [Text, Author, Created]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Add a note to a version

## Trigger

The person writes in the Notes tab and chooses Add note.

## Outcome

The note shows on this version with its author and time.

## Edge cases

- A document in Trash shows its notes read-only.
- A replacement version starts with no notes.
