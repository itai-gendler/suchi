---
appliesTo:
- type: entity
  id: processing-issue
  effect: removes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/tasks.go#RetryDeadJob
- kind: code
  role: implementation
  target: core/api/tasks.go#DismissDeadJob
---

# Only administrators retry or dismiss processing issues

Members do not see processing issues.
