---
domain: groups
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/groups.go#DeleteGroup
- kind: code
  role: implementation
  target: core/authz/store.go#DeleteGroup
---

# Delete a group

An administrator deletes a group that no document access names.
