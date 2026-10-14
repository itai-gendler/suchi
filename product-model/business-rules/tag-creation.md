---
appliesTo:
- type: entity
  id: tag
  effect: creates
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/tags.go#CreateTag
---

# Only administrators create tags

Document screens only assign existing tags.
