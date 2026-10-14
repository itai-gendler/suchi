---
kind: primary
routes:
  web: Web
steps:
- text: A newly added Document is waiting to be read
  kind: condition
  unattended: true
  entities:
  - {entity: document, effect: reads, facts: [Original file]}
- text: The Product extracts or recognizes the text, detects the languages and makes a searchable copy
  kind: product
  entities:
  - {entity: document, effect: changes, facts: [Extracted text, Languages, Searchable copy]}
  contexts:
    web:
      place: suchi-web::archive::document
---

# Read a newly added document

## Trigger

A document is added by any intake path.

## Outcome

The document has searchable text, detected languages, a preview and a thumbnail.

## Edge cases

- Text the phone app recognized is provisional: text read directly from a PDF replaces it, and an empty extraction never erases it.
- E-invoice data in a PDF fills the invoice custom fields that extraction manages.
- A document with a separator sheet becomes one new document per part, each inheriting the parent's filing system, owner, category and sensitivity.
