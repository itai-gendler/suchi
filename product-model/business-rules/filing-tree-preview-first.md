---
appliesTo:
- type: capability
  id: choose-filing-tree
- type: capability
  id: import-filing-tree
references:
- kind: code
  role: implementation
  target: core/jd/preset_change.go
- kind: doc
  role: intent
  target: docs/jd.mdx
---

# A filing tree change applies only exactly as previewed

If the tree or its documents changed after the preview, the change is refused and must be previewed again; a failed change leaves nothing half-applied.
