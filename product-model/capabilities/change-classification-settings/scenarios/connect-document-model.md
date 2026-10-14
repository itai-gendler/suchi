---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator enters the model endpoint and name and tests the connection
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::classification
- text: The endpoint is not on the local network
  kind: condition
  actor: administrator
  entities:
  - {entity: archive-settings, effect: reads, facts: [Document model]}
  contexts:
    web:
      place: suchi-web::archive-configuration::classification
- text: The Administrator acknowledges that extracted text will leave the server and enables the model
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::classification
- text: The Product saves and enables the Archive settings document model
  kind: product
  actor: administrator
  entities:
  - {entity: archive-settings, effect: changes, facts: [Document model, Hosted model acknowledgement]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::classification
---

# Connect a document model

## Trigger

An administrator configures a local or hosted model.

## Outcome

The tested model is enabled for suggestions, dates and Ask your documents.

## Edge cases

- A local model needs no acknowledgement.
- The connection test uses made-up text and checks connectivity and answer format, not accuracy.
