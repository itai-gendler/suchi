---
kind: primary
routes:
  web: Web
steps:
- text: A poll of the watched folder finds a file it has not imported
  kind: condition
  unattended: true
  entities: []
- text: The Product adds a Document owned by the watched-folder owner to the Inbox, applying any sidecar title, tags, correspondent and custom field values
  kind: product
  entities:
  - {entity: document, effect: creates, to: Inbox, facts: [Title, Title ownership, Filing category, Original file, Added from, Added, Original file date, Tags, Correspondents, Custom field values]}
  - {entity: tag, effect: reads, facts: []}
  - {entity: correspondent, effect: reads, facts: []}
  - {entity: custom-field, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::inbox
---

# Import a new file from the watched folder

## Trigger

A new or changed file appears in the watched folder.

## Outcome

The file is a new document in the Inbox of the configured filing system, owned by the watched-folder owner, and is read like an upload.

## Edge cases

- An unknown custom field or invalid value in a sidecar rejects that intake instead of adding a partly described document.
- Unchanged files, hidden files, links and special files are skipped.
