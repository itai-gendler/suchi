---
kind: validation
routes:
  web: Web
steps:
- text: The Administrator tries to disable the last active Administrator
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product refuses the change and leaves the Account active
  kind: product
  actor: administrator
  entities:
  - {entity: account, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Keep at least one active administrator

## Trigger

The change would leave no active administrator, or targets the administrator's own account.

## Outcome

Nothing changes; the administrator's own Active switch is read-only.
