---
domain: filing-system
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/systems.go#PutSystemMembers
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Remove a person from a filing system

An administrator stops a person entering the current filing system.
