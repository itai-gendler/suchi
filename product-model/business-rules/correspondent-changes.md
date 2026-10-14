---
appliesTo:
- type: entity
  id: correspondent
  effect: changes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/taxonomy_crud.go#UpdateCorrespondent
---

# Only administrators rename correspondents

The correspondent catalog is edited under Metadata.
