---
appliesTo:
- type: entity
  id: group
  effect: creates
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/groups.go#CreateGroup
---

# Only administrators create groups

Groups are managed under People.
