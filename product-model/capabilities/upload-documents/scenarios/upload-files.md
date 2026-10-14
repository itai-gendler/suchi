---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses or drops one or more files to upload
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::upload-queue
- text: The Product keeps each original unchanged and adds a Document owned by the Member to the Inbox
  kind: product
  actor: member
  entities:
  - {entity: document, effect: creates, to: Inbox, facts: [Title, Title ownership, Filing category, Original file, Added from, Added, Original file date]}
  contexts:
    web:
      place: suchi-web::archive::upload-queue
- text: The Product queues each Document for reading and shows its progress until it is ready
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title]}
  contexts:
    web:
      place: suchi-web::archive::upload-queue
---

# Upload files into the Inbox

## Trigger

The person chooses Upload, drops files or pastes a screenshot.

## Outcome

Each accepted file is a new document in the Inbox, owned by the person, and its row in the queue links to it once processing is ready.

## Edge cases

- A network or server failure keeps the file in the queue for Retry.
- Accepted uploads keep the filing system they started in, even after the person switches systems.
- Uploading a file whose identical copy the same person has in Trash brings that document back from Trash.
