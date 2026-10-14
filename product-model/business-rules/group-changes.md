---
appliesTo:
- type: entity
  id: group
  effect: changes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/groups.go#AddGroupMember
---

# Only administrators change group membership

Groups are managed under People.
