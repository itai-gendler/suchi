---
domain: filing-system
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/systems.go#PutSystemMembers
---

# Admit a person to a filing system

An administrator lets a person enter the current filing system. Entry grants no access to anyone's documents by itself.
