---
kind: primary
result: achieved
routes:
  web: Web
steps:
- text: The Administrator enters the setup token with their email, display name and password
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::sign-in::setup
  capability: set-up-first-administrator
- text: The Product creates the Administrator with an active Account, signs them in and asks them to choose a filing tree
  kind: product
  actor: administrator
  entities:
  - {entity: administrator, effect: creates, facts: []}
  - {entity: account, effect: creates, to: Active, facts: [Email, Display name, Password, Sign-in method], with: administrator}
  contexts:
    web:
      place: suchi-web::sign-in::setup
  capability: set-up-first-administrator
- text: The Administrator chooses Continue setup, picks a ready-made tree and applies the preview
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
  capability: choose-filing-tree
- text: The Product adds the categories and records the choice for the Filing system
  kind: product
  actor: administrator
  entities:
  - {entity: category, effect: creates, facts: [Code, Name, Description, Area]}
  - {entity: filing-system, effect: changes, facts: [Filing tree choice]}
  contexts:
    web:
      place: suchi-web::archive-configuration::filing-tree
  capability: choose-filing-tree
---

# Set up with a ready-made tree

## Trigger

Someone opens a new server for the first time.

## Outcome

The first administrator is signed in and the filing system has the chosen tree.
