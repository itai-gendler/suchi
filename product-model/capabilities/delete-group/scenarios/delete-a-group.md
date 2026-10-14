---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator deletes the Group and confirms
  kind: actor
  actor: administrator
  entities:
  - {entity: group, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
- text: The Product deletes the Group
  kind: product
  actor: administrator
  entities:
  - {entity: group, effect: removes}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
---

# Delete a group

## Trigger

An administrator deletes a group and confirms.

## Outcome

The group is gone.
