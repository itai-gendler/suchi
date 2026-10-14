---
kind: primary
routes:
  web: Web
steps:
- text: The Member filters by text, category, tag, correspondent, sensitivity, sharing or added date and chooses a sort
  kind: actor
  actor: member
  entities:
  - {entity: tag, effect: reads, facts: []}
  - {entity: correspondent, effect: reads, facts: []}
  - {entity: category, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::documents
- text: The Product lists the newest visible version of each matching Document, page by page
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title, Filing category, Filing address, Tags, Correspondents, Sensitivity, Added]}
  contexts:
    web:
      place: suchi-web::archive::documents
---

# Filter and sort documents

## Trigger

The person opens Documents or a category from the Index.

## Outcome

The list shows only matching documents the person may open, and the filters stay in the address for reload and sharing.

## Edge cases

- Shared by me lists each document once while one of the person's own share links for it is active.
- No correspondent matches documents with no correspondent in any role.
- The version: query term lists every version or only earlier ones.
