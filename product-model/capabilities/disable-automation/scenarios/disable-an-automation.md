---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator switches the Automation off
  kind: actor
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
- text: The Product disables the Automation
  kind: product
  actor: administrator
  entities:
  - {entity: automation, effect: changes, from: Enabled, to: Disabled, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
---

# Disable an automation

## Trigger

An administrator switches an automation off.

## Outcome

The automation is kept but no longer runs.
