---
appliesTo:
- type: entity
  id: tag
  effect: removes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/tags.go#DeleteTags
---

# Only administrators delete tags

The tag catalog is edited only under Metadata.
