---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator deletes the Automation and confirms
  kind: actor
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
- text: The Product deletes the Automation
  kind: product
  actor: administrator
  entities:
  - {entity: automation, effect: removes, from: Enabled}
  contexts:
    web:
      place: suchi-web::archive::automations
---

# Delete an automation

## Trigger

An administrator chooses Delete automation and confirms.

## Outcome

The automation no longer runs and is gone from the list.
