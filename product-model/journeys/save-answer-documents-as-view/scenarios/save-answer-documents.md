---
kind: primary
result: achieved
routes:
  web: Web
steps:
- text: The Member asks a question about their documents
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::ask-your-documents
  capability: ask-documents
- text: The Product shows the answer with links to the Documents it used
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title]}
  contexts:
    web:
      place: suchi-web::archive::ask-your-documents
  capability: ask-documents
- text: The Member chooses Save these documents as a view
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::ask-your-documents
  capability: ask-documents
- text: The Product opens the view builder holding exactly those documents
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::views
  capability: ask-documents
- text: The Member names the view and saves it
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::views
  capability: create-saved-view
- text: The Product saves the Saved view of those exact documents
  kind: product
  actor: member
  entities:
  - {entity: saved-view, effect: creates, facts: [Name, Filter, Shared]}
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::views
  capability: create-saved-view
---

# Save the documents of an answer

## Trigger

The person has an answer from Ask your documents.

## Outcome

A saved view keeps the exact documents of the answer.
