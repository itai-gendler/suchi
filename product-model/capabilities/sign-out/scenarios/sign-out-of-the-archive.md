---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Sign out
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive
- text: The Product ends the session and returns to the sign-in page
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive
---

# Sign out of the archive

## Trigger

The person chooses Sign out.

## Outcome

The session ends and the browser returns to the sign-in page.
