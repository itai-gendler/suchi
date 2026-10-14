---
kind: primary
routes:
  web: Web
steps:
- text: The Member changes their display name, photo or language and saves
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::profile
- text: The Product saves the Account profile
  kind: product
  actor: member
  entities:
  - {entity: account, effect: changes, facts: [Display name, Photo, Language]}
  contexts:
    web:
      place: suchi-web::archive::profile
---

# Save display name, photo and language

## Trigger

The person opens Profile and changes their details.

## Outcome

The new display name and photo show wherever the person appears, and the chosen language applies across their signed-in browsers.

## Edge cases

- Unsupported languages and missing messages fall back to English.
