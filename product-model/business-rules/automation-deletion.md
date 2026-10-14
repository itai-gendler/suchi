---
appliesTo:
- type: entity
  id: automation
  effect: removes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/automations.go#DeleteAutomation
---

# Only administrators delete automations

Members see automations read-only.
