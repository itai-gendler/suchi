---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator enters a name and creates the Group
  kind: actor
  actor: administrator
  entities:
  - {entity: group, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
- text: The Product creates the Group
  kind: product
  actor: administrator
  entities:
  - {entity: group, effect: creates, facts: [Name, Members]}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
---

# Create a group

## Trigger

An administrator enters a group name under Groups.

## Outcome

The group exists, empty.
