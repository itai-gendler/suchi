---
relations:
- {entity: document-access, verb: holds, cardinality: one-to-many}
- {entity: member, verb: includes, cardinality: many-to-many}
references:
- kind: code
  role: implementation
  target: core/api/groups.go#CreateGroup
---

# Group

A named set of people that one document access grant can cover at once.

## Information kept

- **Name** — the group name
- **Members** — the people in the group
