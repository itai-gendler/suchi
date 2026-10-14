---
kind: primary
routes:
  web: Web
steps:
- text: The Link recipient opens the Share link
  kind: actor
  actor: link-recipient
  entities:
  - {entity: share-link, effect: reads, facts: []}
  contexts:
    web:
      place: share-page::shared-documents
- text: The link is active, unexpired and its creator may still share these documents
  kind: condition
  actor: link-recipient
  entities:
  - {entity: share-link, effect: reads, facts: [Expiry, Documents]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: share-page::shared-documents
- text: The Product lists the shared Documents that are still live and offers their download
  kind: product
  actor: link-recipient
  entities:
  - {entity: document, effect: reads, facts: [Title, Original file]}
  contexts:
    web:
      place: share-page::shared-documents
- text: The Link recipient downloads a Document
  kind: actor
  actor: link-recipient
  entities:
  - {entity: document, effect: reads, facts: [Title]}
  contexts:
    web:
      place: share-page::shared-documents
---

# Download shared documents

## Trigger

A recipient opens an active share link.

## Outcome

The recipient sees the expiry and the shared documents that are still live and downloads their original files.

## Edge cases

- Documents moved to Trash since the link was made are left out.
