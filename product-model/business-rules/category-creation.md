---
appliesTo:
- type: entity
  id: category
  effect: creates
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/setup.go#ApplyPresetChange
---

# Only administrators change the filing tree

Categories come only from a chosen, composed or imported tree.
