---
entities:
- entity: saved-view
  shows:
  - Name
  - Filter
  - Shared
  collects:
  - Name
  - Filter
  - Shared
- entity: document
  shows:
  - Title
- entity: tag
  shows:
  - Name
- entity: correspondent
  shows:
  - Name
- entity: category
  shows:
  - Code
  - Name
- entity: custom-field
  shows:
  - Name
references:
- kind: code
  role: implementation
  target: ui/src/routes/Views.svelte
---

# Views

The person's saved views and the views shared with them, and the builder for one view.
