---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects Locked Documents and tries one password against them
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product unlocks each selected Document the password opens and reports how many opened
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: [Password protection]}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Try one password on a selection

## Trigger

The person selects Locked documents and chooses Bulk unlock.

## Outcome

Every selected document the password opens is unlocked, and the person sees how many opened.
