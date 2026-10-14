---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator ticks an active person under Members and saves
  kind: actor
  actor: administrator
  entities:
  - {entity: member, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-system
- text: The Product admits the person to the Filing system
  kind: product
  actor: administrator
  entities:
  - {entity: filing-system, effect: changes, facts: [Members]}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-system
---

# Admit a person

## Trigger

An administrator ticks a person under Members and saves.

## Outcome

The person may enter the filing system.

## Edge cases

- Administrators always enter every filing system.
- Access granted earlier on documents of this system becomes usable again.
