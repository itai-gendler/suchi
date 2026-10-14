---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects Documents and chooses Rescan
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product queues each selected Document the Member may change for reading again
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Rescan selected documents

## Trigger

The person selects documents and chooses Rescan.

## Outcome

The selected documents are queued to be read again; pending model suggestions are replaced while accepted and rejected dates are kept.

## Edge cases

- A rescan alone never counts as reviewing a document's review marker.
