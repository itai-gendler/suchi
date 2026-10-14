---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator changes the triggers, actions or question and saves
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive::automations
- text: The Product validates and saves the Automation
  kind: product
  actor: administrator
  entities:
  - {entity: automation, effect: changes, facts: [Name, Triggers, Actions, Classifier question]}
  contexts:
    web:
      place: suchi-web::archive::automations
---

# Edit an automation

## Trigger

An administrator chooses Edit on an automation.

## Outcome

The automation runs with its new definition.
