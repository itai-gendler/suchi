---
kind: validation
routes:
  web: Web
steps:
- text: The Member submits a query with an invalid term
  kind: actor
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::search
- text: The Product rejects the query and shows its error
  kind: product
  actor: member
  entities: []
  contexts:
    web:
      place: suchi-web::archive::search
---

# Reject an invalid query

## Trigger

The person submits a query the language does not accept.

## Outcome

The query stays visible with its error and no results replace the list.
