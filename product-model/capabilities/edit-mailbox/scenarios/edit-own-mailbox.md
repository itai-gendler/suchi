---
kind: primary
routes:
  web: Web
steps:
- text: The Member changes the Mailbox settings and saves
  kind: actor
  actor: member
  entities:
  - {entity: mailbox, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::my-mailboxes
- text: The Product saves the Mailbox and reloads its polling
  kind: product
  actor: member
  entities:
  - {entity: mailbox, effect: changes, facts: [Name, Source folder, Intake rules, After import, Sync mail from, Poll interval, Enabled]}
  contexts:
    web:
      place: suchi-web::archive::my-mailboxes
---

# Edit your own mailbox

## Trigger

A member edits one of their mailboxes and saves.

## Outcome

The mailbox polls with the new settings; changing the connection, folder or sync date starts again from the sync date without importing messages twice.

## Edge cases

- Test connection checks connection and sign-in only and does not clear a sync error.
