---
kind: primary
routes:
  web: Web
steps:
- text: The Member opens a Document
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product shows the preview, metadata, extracted text, similar documents, linked documents, versions and notes
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title, Filing category, Filing address, Sensitivity, Tags, Correspondents, Custom field values, Extracted text, Added from, Similar documents]}
  - {entity: document-link, effect: reads, facts: [Link type, Target document]}
  - {entity: note, effect: reads, facts: [Text, Author, Created]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Open a document

## Trigger

The person opens a document from a list, search, a link or its address.

## Outcome

The person sees the document and everything kept with it that they may see.

## Edge cases

- A filing address opens the document only while it is filed at that exact address.
- Linked documents the person cannot see, or that are in Trash, are left out.
