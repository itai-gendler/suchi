---
kind: primary
routes:
  web: Web
steps:
- text: The Member opens the Dashboard
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::dashboard
- text: The Product shows the totals, the six newest visible Documents and up to four Saved views
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title, Added]}
  - {entity: saved-view, effect: reads, facts: [Name, Filter]}
  contexts:
    web:
      place: suchi-web::archive::dashboard
---

# Open the dashboard

## Trigger

A signed-in person opens the archive.

## Outcome

The person sees the counts, the six newest documents they may open and up to four saved views; counts that could not load show a dash, never zero.
