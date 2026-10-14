---
kind: validation
routes:
  web: Web
steps:
- text: The Administrator saves an automation with an invalid pattern or action
  kind: actor
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
- text: The Product refuses the Automation
  kind: product
  actor: administrator
  entities:
  - {entity: automation, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::automations
---

# Refuse a malformed automation

## Trigger

The administrator saves a filter pattern or action the rules do not accept.

## Outcome

Nothing is saved and the builder shows the error, so no rule sits idle unnoticed.
