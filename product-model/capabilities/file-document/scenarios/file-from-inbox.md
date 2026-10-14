---
kind: primary
routes:
  web: Web
steps:
- text: The Member chooses a category under Filed under
  kind: actor
  actor: member
  entities:
  - {entity: category, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
- text: The Product files the Document under that category and removes a classifier-added needs-review tag
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, from: Inbox, to: Filed, facts: [Filing category, Filing address, Tags]}
  - {entity: tag, effect: reads, facts: []}
  - {entity: category, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# File a document waiting in the Inbox

## Trigger

The person chooses a category under Filed under for a document in the Inbox.

## Outcome

The document leaves the Inbox, its filing address names the new category, and a review marker the classifier added is removed.

## Edge cases

- A pending filing Suggestion for the document becomes outdated.
