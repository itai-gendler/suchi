---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator turns matching on or off and sets the automatic and review thresholds
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::classification
- text: The Product saves the Archive settings matching options
  kind: product
  actor: administrator
  entities:
  - {entity: archive-settings, effect: changes, facts: [Similar documents matching, Automatic match threshold, Review match threshold]}
  contexts:
    web:
      place: suchi-web::archive-configuration::classification
---

# Adjust similar-document matching

## Trigger

An administrator opens Similar documents and adjusts it.

## Outcome

Archive matching suggests or applies details at the new thresholds for new documents.
