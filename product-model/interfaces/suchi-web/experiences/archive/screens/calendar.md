---
entities:
- entity: document-date
  shows:
  - Date
  - Precision
  - Role
  - Document text
  - Reviewed by
- entity: document
  shows:
  - Title
references:
- kind: code
  role: implementation
  target: ui/src/routes/Calendar.svelte
---

# Calendar

The accepted dates of the documents the person may read, by month or rolling range.
