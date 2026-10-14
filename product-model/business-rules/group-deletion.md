---
appliesTo:
- type: entity
  id: group
  effect: removes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/groups.go#DeleteGroup
---

# Only administrators delete groups

Groups are managed under People.
