---
kind: edge
routes:
  web: Web
steps:
- text: A processing step for a Document fails after its retries
  kind: condition
  unattended: true
  entities:
  - {entity: document, effect: reads, facts: []}
- text: The Product records a Processing issue with the step, error and retries
  kind: product
  entities:
  - {entity: processing-issue, effect: creates, facts: [Processing step, Error, Retries]}
  contexts:
    web:
      place: suchi-web::archive::approvals
---

# Report a processing step that stopped

## Trigger

A processing step keeps failing.

## Outcome

Administrators see the step under Processing issues with its error and retry count.
