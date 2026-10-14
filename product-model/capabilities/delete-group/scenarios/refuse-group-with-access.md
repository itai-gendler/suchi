---
kind: validation
routes:
  web: Web
steps:
- text: The Administrator deletes a Group that still holds access
  kind: actor
  actor: administrator
  entities:
  - {entity: group, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
- text: The Product refuses the deletion because Document access still names the Group
  kind: product
  actor: administrator
  entities:
  - {entity: document-access, effect: reads, facts: [Person or group]}
  - {entity: document, effect: reads, facts: []}
  - {entity: group, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::groups
---

# Refuse to delete a group that holds access

## Trigger

Some document access still names the group.

## Outcome

Nothing is deleted; the administrator is told to revoke that access first.
