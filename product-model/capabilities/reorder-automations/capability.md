---
domain: automations
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: ui/src/routes/Automations.svelte
- kind: code
  role: implementation
  target: core/api/automations.go#UpdateAutomation
---

# Reorder automations

An administrator changes the order in which custom automations run.
