---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses Upload replacement on the latest version and picks the new file
  kind: actor
  actor: member
  entities:
  - {entity: document, as: previous, effect: reads, facts: [Title]}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product adds the file as a new Document that replaces the previous one, carrying its filing system, owner, category, sensitivity, title, access, tags a person or rule assigned and correspondents
  kind: product
  actor: member
  entities:
  - {entity: document, as: previous, effect: reads, facts: [Filing category, Sensitivity, Title, Title ownership, Tags, Correspondents]}
  - {entity: document, as: replacement, effect: creates, to: Filed, facts: [Title, Title ownership, Filing category, Filing address, Sensitivity, Tags, Correspondents, Previous version, Original file, Added from, Added]}
  - {entity: document-access, effect: creates, facts: [Person or group, Access level], with: replacement}
  - {entity: tag, effect: reads, facts: []}
  - {entity: correspondent, effect: reads, facts: []}
  - {entity: category, effect: reads, facts: []}
  - {entity: filing-system, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product queues the new version for reading with its version-specific details starting fresh
  kind: product
  actor: member
  entities:
  - {entity: document, as: replacement, effect: reads, facts: [Previous version]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Replace the latest version with a newer file

## Trigger

The person chooses Upload replacement on the latest version of a document.

## Outcome

A new version exists with the continuity metadata of the previous one and is read like an upload; the previous version stays available as Earlier.

## Edge cases

- A replacement for a document in the Inbox stays in the Inbox.
- Extracted text, model-added tags, dates, custom field values, links to other documents and notes do not carry over.
- A failed upload keeps the chosen file and its original date for a safe retry.
