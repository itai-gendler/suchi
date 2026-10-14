---
relations:
- {entity: document-link, verb: names, cardinality: one-to-many}
references:
- kind: code
  role: implementation
  target: core/api/custom_fields.go#CreateCustomFieldDef
---

# Custom field

A typed field an administrator defines for the documents of one filing system: text, URL, number, monetary, date, yes/no, select, multi-choice or Document link.

## Information kept

- **Name** — the field name
- **Type** — the declared value type
- **Choices** — the allowed choices of a select or multi-choice field
- **Managed by extraction** — whether invoice extraction owns the definition, making it read-only
