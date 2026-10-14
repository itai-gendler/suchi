---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator chooses Retry on a Processing issue
  kind: actor
  actor: administrator
  entities:
  - {entity: processing-issue, effect: reads, facts: [Processing step, Error]}
  contexts:
    web:
      place: suchi-web::archive::approvals
- text: The Product queues the step again and removes the Processing issue
  kind: product
  actor: administrator
  entities:
  - {entity: processing-issue, effect: removes}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Try a processing step again

## Trigger

An administrator chooses Retry on a processing issue.

## Outcome

The step runs again and the issue leaves the list.
