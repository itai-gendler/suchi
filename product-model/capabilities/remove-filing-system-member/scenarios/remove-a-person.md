---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator unticks a person under Members and saves
  kind: actor
  actor: administrator
  entities:
  - {entity: member, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-system
- text: The Product removes the person from the Filing system and revokes their Share links there
  kind: product
  actor: administrator
  entities:
  - {entity: filing-system, effect: changes, facts: [Members]}
  - {entity: share-link, effect: changes, from: Active, to: Revoked, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-system
---

# Remove a person

## Trigger

An administrator unticks a person under Members and saves.

## Outcome

The person can no longer enter the filing system; their share links for it are revoked and do not come back if they are admitted again.
