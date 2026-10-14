---
kind: validation
routes:
  web: Web
steps:
- text: The Administrator previews a tree file with errors
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
- text: The Product rejects the file with its validation errors
  kind: product
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
---

# Refuse an invalid tree file

## Trigger

The file has unknown keys, wrong types or an unsupported format.

## Outcome

Nothing is imported and the preview shows the errors.
