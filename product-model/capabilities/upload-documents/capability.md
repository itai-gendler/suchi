---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/documents.go#UploadDocument
- kind: code
  role: implementation
  target: core/api/upload_prepare.go
- kind: code
  role: implementation
  target: ui/src/lib/UploadBox.svelte
- kind: doc
  role: intent
  target: docs/web-app.mdx
---

# Upload documents

A person adds files from the browser: chosen, dropped anywhere in the archive or pasted as a screenshot. Each file becomes a document they own in the current filing system, waiting in the Inbox while it is read.
