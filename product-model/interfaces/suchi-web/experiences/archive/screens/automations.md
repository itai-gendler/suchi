---
entities:
- entity: automation
  shows:
  - Name
  - Triggers
  - Actions
  - Classifier question
  - Run order
references:
- kind: code
  role: implementation
  target: ui/src/routes/Automations.svelte
---

# Automations

The automations of the current filing system in run order, and the builder for one automation; read-only for members.
