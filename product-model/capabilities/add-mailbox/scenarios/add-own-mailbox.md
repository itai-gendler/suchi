---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses the provider, name and username, signs in or enters an app password, and sets intake rules and what happens after import
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::my-mailboxes
- text: The Product saves the Mailbox owned by the Member and starts polling it
  kind: product
  actor: member
  entities:
  - {entity: mailbox, effect: creates, facts: [Name, Provider, Username, Credentials, Source folder, Intake rules, After import, Sync mail from, Poll interval, Enabled]}
  contexts:
    web:
      place: suchi-web::archive::my-mailboxes
---

# Add your own mailbox

## Trigger

A member chooses Add mailbox.

## Outcome

The mailbox is polled within seconds, from the chosen sync date onwards, and its documents belong to the member.

## Edge cases

- Sync mail from defaults to now, so years of history are not imported unexpectedly.
- Preview matches checks the intake rules against up to 25 recent messages.
