---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator moves an Automation up or down the list
  kind: actor
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
- text: The Product saves the new run order
  kind: product
  actor: administrator
  entities:
  - {entity: automation, effect: changes, facts: [Run order]}
  contexts:
    web:
      place: suchi-web::archive::automations
---

# Move an automation in the run order

## Trigger

An administrator drags an automation, or moves it with the keyboard.

## Outcome

Automations run top to bottom in the new order.
