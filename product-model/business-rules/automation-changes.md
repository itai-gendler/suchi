---
appliesTo:
- type: entity
  id: automation
  effect: changes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/automations.go#UpdateAutomation
---

# Only administrators change automations

Members see automations read-only; this covers editing, enabling, disabling and reordering.
