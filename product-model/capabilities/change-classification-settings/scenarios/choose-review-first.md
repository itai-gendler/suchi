---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator chooses Review first and saves the application mode
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::classification
- text: The Product saves the Archive settings application mode
  kind: product
  actor: administrator
  entities:
  - {entity: archive-settings, effect: changes, facts: [Application mode]}
  contexts:
    web:
      place: suchi-web::archive-configuration::classification
---

# Review everything first

## Trigger

An administrator chooses Review first and saves the application mode.

## Outcome

From now on every inferred change and date waits in Approvals, even at full confidence; explicit automations are unaffected.
