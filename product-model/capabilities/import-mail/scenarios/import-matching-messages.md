---
kind: primary
routes:
  web: Web
steps:
- text: A poll of an enabled Mailbox finds messages since its sync point
  kind: condition
  unattended: true
  entities:
  - {entity: mailbox, effect: reads, facts: [Source folder, Sync mail from, Enabled]}
- text: The Product adds each message an intake rule accepts as Documents owned by the Mailbox owner in the Inbox
  kind: product
  entities:
  - {entity: mailbox, effect: reads, facts: [Intake rules]}
  - {entity: document, effect: creates, to: Inbox, facts: [Title, Title ownership, Filing category, Original file, Added from, Added, Original file date]}
  contexts:
    web:
      place: suchi-web::archive::inbox
- text: The Product leaves, marks read or moves the imported messages as the Mailbox says and records the sync
  kind: product
  entities:
  - {entity: mailbox, effect: changes, facts: [Sync status]}
  contexts:
    web:
      place: suchi-web::archive::my-mailboxes
---

# Import messages that match an intake rule

## Trigger

A mailbox poll finds messages newer than its sync point.

## Outcome

Each accepted message is in the Inbox as the email with its files, or as its files only, and the mailbox shows its last successful sync.

## Edge cases

- When several rules match with different choices, the email is kept with its files.
- A message already imported for the same owner and filing system is not imported again.
- A message over the size bound is skipped and recorded so it cannot block the mailbox.
- A mailbox whose owner is disabled is not polled.
