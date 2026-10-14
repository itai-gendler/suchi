---
appliesTo:
- type: entity
  id: tag
  effect: changes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/tags.go#UpdateTag
---

# Only administrators rename tags

The tag catalog is edited only under Metadata.
