---
appliesTo:
- type: entity
  id: correspondent
  effect: creates
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/taxonomy_crud.go#CreateCorrespondent
---

# Only administrators create correspondents

Document screens only choose existing correspondents.
