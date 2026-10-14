---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator switches the Automation on
  kind: actor
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
- text: The Product enables the Automation
  kind: product
  actor: administrator
  entities:
  - {entity: automation, effect: changes, from: Disabled, to: Enabled, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
---

# Enable an automation

## Trigger

An administrator switches a disabled automation on.

## Outcome

The automation runs again on matching documents.
