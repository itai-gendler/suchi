---
kind: edge
routes:
  web: Web
steps:
- text: The Administrator saves a new automation identical to an existing one
  kind: actor
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
- text: An Automation with the same triggers and actions already exists
  kind: condition
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: [Triggers, Actions]}
  contexts:
    web:
      place: suchi-web::archive::automations
- text: The Product refuses the copy and offers to open or enable the existing Automation
  kind: product
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
---

# Point to an identical automation instead of duplicating it

## Trigger

The administrator saves a new automation identical to an existing one.

## Outcome

No copy is created; the administrator can open the existing automation or enable it if it is disabled.
