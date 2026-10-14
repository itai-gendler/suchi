---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator adds triggers with filters, ordered actions and an optional classifier question, and saves
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive::automations
- text: The Product validates the filters and actions and saves the Automation enabled at the end of the run order
  kind: product
  actor: administrator
  entities:
  - {entity: automation, effect: creates, to: Enabled, facts: [Name, Triggers, Actions, Classifier question, Run order]}
  contexts:
    web:
      place: suchi-web::archive::automations
---

# Create an automation

## Trigger

An administrator chooses New automation or a starter pattern.

## Outcome

The automation is enabled at the end of the run order and applies to documents from now on.

## Edge cases

- Existing documents are only affected when an administrator refiles or they are updated.
