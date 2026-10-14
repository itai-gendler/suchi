---
appliesTo:
- type: entity
  id: filing-system
  effect: changes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/systems.go#PutSystemMembers
---

# Only administrators change a filing system

This covers its name, its members and its filing tree choice.
