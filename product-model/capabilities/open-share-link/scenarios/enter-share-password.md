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
- text: The link has a password
  kind: condition
  actor: link-recipient
  entities:
  - {entity: share-link, effect: reads, facts: [Password]}
  contexts:
    web:
      place: share-page::shared-documents
- text: The Link recipient enters the password
  kind: actor
  actor: link-recipient
  entities: []
  contexts:
    web:
      place: share-page::shared-documents
- text: The Product admits the recipient to the shared Documents
  kind: product
  actor: link-recipient
  entities:
  - {entity: document, effect: reads, facts: [Title]}
  contexts:
    web:
      place: share-page::shared-documents
---

# Enter the share password

## Trigger

A recipient opens a link that has a password.

## Outcome

After the right password the documents are listed; a wrong one shows the form again.
