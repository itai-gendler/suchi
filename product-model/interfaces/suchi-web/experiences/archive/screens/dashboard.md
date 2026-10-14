---
entities:
- entity: document
  shows:
  - Title
  - Added
- entity: saved-view
  shows:
  - Name
  - Filter
- entity: filing-system
  shows:
  - Filing tree choice
references:
- kind: code
  role: implementation
  target: ui/src/routes/Dashboard.svelte
---

# Dashboard

The overview of the current filing system and, until a filing tree is chosen, the setup step.
