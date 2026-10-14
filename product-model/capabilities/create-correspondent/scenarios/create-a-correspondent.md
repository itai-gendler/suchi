---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator enters a new Correspondent name
  kind: actor
  actor: administrator
  entities:
  - {entity: correspondent, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive-configuration::correspondents
- text: The Product creates the Correspondent
  kind: product
  actor: administrator
  entities:
  - {entity: correspondent, effect: creates, facts: [Name]}
  contexts:
    web:
      place: suchi-web::archive-configuration::correspondents
---

# Create a correspondent

## Trigger

An administrator enters a new correspondent name.

## Outcome

People can choose the correspondent on documents.
