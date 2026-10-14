---
entities:
- entity: suggestion
  shows:
  - Field
  - Current value
  - Suggested value
  - Why suggested
  - Model confidence
  - Suggested by
- entity: document-date
  shows:
  - Date
  - Precision
  - Role
  - Document text
  - Model confidence
- entity: processing-issue
  shows:
  - Processing step
  - Error
  - Retries
- entity: document
  shows:
  - Title
  - Filing category
references:
- kind: code
  role: implementation
  target: ui/src/routes/Tasks.svelte
---

# Approvals

Decisions waiting for the person: suggested document changes and dates, and for administrators the processing issues.
