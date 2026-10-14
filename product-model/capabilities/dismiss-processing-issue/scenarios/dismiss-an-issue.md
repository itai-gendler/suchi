---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator chooses Dismiss on a Processing issue
  kind: actor
  actor: administrator
  entities:
  - {entity: processing-issue, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Product removes the Processing issue
  kind: product
  actor: administrator
  entities:
  - {entity: processing-issue, effect: removes}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Dismiss a processing issue

## Trigger

An administrator chooses Dismiss on a processing issue.

## Outcome

The issue leaves the list and the step is not run again.
