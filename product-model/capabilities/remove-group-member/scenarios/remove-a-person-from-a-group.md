---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator removes a person from the Group
  kind: actor
  actor: administrator
  entities:
  - {entity: group, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
- text: The Product removes the person from the Group
  kind: product
  actor: administrator
  entities:
  - {entity: group, effect: changes, facts: [Members]}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
---

# Remove a person from a group

## Trigger

An administrator removes a person from a group's members.

## Outcome

The person loses the access the group gave them.
