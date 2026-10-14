---
relations:
- {entity: document, verb: files, cardinality: one-to-many}
references:
- kind: code
  role: implementation
  target: core/jd/preset_change.go
- kind: doc
  role: intent
  target: docs/jd.mdx
---

# Category

One numbered category of a filing tree, inside a ten-code area, that documents are filed under. The System area 40–49 and its 49 Inbox are generated and protected.

## Information kept

- **Code** — the two-digit category code
- **Name** — the category name
- **Description** — what belongs there and where near neighbours go instead
- **Area** — the area the category belongs to
