---
kind: validation
routes:
  web: Web
steps:
- text: The Member chooses a file larger than the upload limit
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::upload-queue
- text: The Product refuses the file without storing it
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::upload-queue
---

# Refuse an oversized file

## Trigger

The person uploads a file over the server's upload limit.

## Outcome

Nothing is stored; the queue asks for a smaller file or a changed server limit.
