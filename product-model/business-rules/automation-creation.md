---
appliesTo:
- type: entity
  id: automation
  effect: creates
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/automations.go#CreateAutomation
---

# Only administrators create automations

Members see automations read-only.
