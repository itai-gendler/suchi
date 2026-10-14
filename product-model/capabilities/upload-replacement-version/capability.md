---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/versions.go#UploadNewVersion
- kind: doc
  role: intent
  target: docs/document-versions-and-links.mdx
---

# Upload a replacement version

A person adds a newer file for a document. The new file becomes a new document that replaces the previous one; both stay available, and discovery shows the newest version each person may see.
