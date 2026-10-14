---
appliesTo:
- type: entity
  id: correspondent
  effect: removes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/taxonomy_crud.go#DeleteCorrespondent
---

# Only administrators delete correspondents

The correspondent catalog is edited under Metadata.
