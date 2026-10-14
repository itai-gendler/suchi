---
relations:
- {entity: category, verb: contains, cardinality: one-to-many}
references:
- kind: code
  role: implementation
  target: core/api/systems.go#PutSystemMembers
- kind: doc
  role: intent
  target: docs/jd.mdx
---

# Filing system

One filing cabinet of the archive with its own filing tree, metadata catalogs, rules and views. A server starts with one hidden filing system; the first taxonomy import that names a system code introduces codes and lets administrators add more.

## Information kept

- **Code** — the permanent code, such as S01
- **Name** — the editable display name, at most 80 characters
- **Members** — the people admitted to enter it besides administrators
- **Filing tree choice** — the ready-made tree, composed tree, import or Blank chosen for it, if any
