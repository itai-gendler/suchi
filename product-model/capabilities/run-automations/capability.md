---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/automations/apply.go
- kind: code
  role: implementation
  target: core/automations/builtin_actions.go
- kind: doc
  role: intent
  target: docs/automations.mdx
---

# Run automations

Enabled automations run top to bottom on documents their triggers select: when a document is added, after its text is read, or after a person updates its metadata. Each automation applies its actions together or not at all.
