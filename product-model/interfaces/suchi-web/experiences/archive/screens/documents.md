---
entities:
- entity: document
  shows:
  - Title
  - Filing category
  - Filing address
  - Tags
  - Correspondents
  - Sensitivity
  - Added
  - Password protection
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
references:
- kind: code
  role: implementation
  target: ui/src/routes/Documents.svelte
---

# Documents

The list of documents the person may open in the current filing system, with filters, selection and bulk actions.
