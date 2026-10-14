---
domain: automations
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/automations.go#CreateAutomation
- kind: code
  role: implementation
  target: core/automations/registry.go
- kind: doc
  role: intent
  target: docs/automations.mdx
---

# Create an automation

An administrator writes a filing rule for the current filing system: triggers with filters, ordered actions and, optionally, one bounded question for the document model with an action per answer. Members see the list read-only.
