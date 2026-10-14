---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator picks a person to add to the Group
  kind: actor
  actor: administrator
  entities:
  - {entity: group, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
- text: The Product adds the person to the Group
  kind: product
  actor: administrator
  entities:
  - {entity: group, effect: changes, facts: [Members]}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
---

# Add a person to a group

## Trigger

An administrator opens a group's members and adds a person.

## Outcome

The person gets every access the group holds.
