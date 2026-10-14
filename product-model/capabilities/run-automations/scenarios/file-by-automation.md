---
kind: primary
routes:
  web: Web
steps:
- text: An enabled Automation's triggers match a Document
  kind: condition
  unattended: true
  entities:
  - {entity: automation, effect: reads, facts: [Triggers, Run order]}
  - {entity: document, effect: reads, facts: [Extracted text, Added from]}
- text: The Product applies the Automation actions in order, filing the Document and setting its title, tags and correspondent
  kind: product
  entities:
  - {entity: automation, effect: reads, facts: [Actions]}
  - {entity: document, effect: changes, from: Inbox, to: Filed, facts: [Filing category, Filing address, Title, Title ownership, Tags, Correspondents]}
  - {entity: tag, effect: reads, facts: []}
  - {entity: correspondent, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::document
---

# File a document with a matching automation

## Trigger

A document is added or updated and an enabled automation's filters match it.

## Outcome

The document carries what the automation assigns, such as its category, tags, correspondent or title, and leaves the Inbox when a category is assigned.

## Edge cases

- A title a person set is never replaced; a title a rule set may be refreshed by a rule.
- An automation does not trigger itself again with its own changes.
- A failing action rolls back that automation only.
- An automation with a classifier question applies the action mapped to the model's answer only through the same automatic or review path as other model suggestions.
