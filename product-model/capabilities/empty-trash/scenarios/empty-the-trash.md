---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Empty Trash and confirms
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::trash
- text: The Product permanently deletes each Document in the Member's Trash they may delete
  kind: product
  actor: member
  entities:
  - {entity: document, effect: removes, from: Trashed}
  contexts:
    web:
      place: suchi-web::archive::trash
---

# Empty Trash

## Trigger

The person chooses Empty Trash and confirms.

## Outcome

Every document in the person's Trash they may delete is gone; the person sees how many.

## Edge cases

- A member's Trash lists only the documents they own; an administrator's lists every trashed document of the filing system.
