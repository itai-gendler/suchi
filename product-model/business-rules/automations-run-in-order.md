---
appliesTo:
- type: capability
  id: run-automations
- type: capability
  id: reorder-automations
references:
- kind: code
  role: implementation
  target: core/automations/apply.go
- kind: doc
  role: intent
  target: docs/automations.mdx
---

# Automations run top to bottom, each all or nothing

A failing action rolls back only its own automation; an automation does not re-trigger itself.
