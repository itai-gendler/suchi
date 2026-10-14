---
kind: primary
routes:
  web: Web
steps:
- text: The Member selects Documents and chooses a category under File under
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  - {entity: category, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product files each selected Document the Member may change and reports how many succeeded
  kind: product
  actor: member
  entities:
  - {entity: document, effect: changes, from: Inbox, to: Filed, facts: [Filing category, Filing address]}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# File selected documents

## Trigger

The person selects documents in Documents or Inbox and chooses File under.

## Outcome

Every selected document the person may change is filed; the person sees how many succeeded and failed.
