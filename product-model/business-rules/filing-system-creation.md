---
appliesTo:
- type: entity
  id: filing-system
  effect: creates
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/taxonomy.go#ImportTaxonomy
---

# Only administrators create filing systems

A filing system is created by importing a tree that names a new code.
