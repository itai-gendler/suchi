---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/pipeline/postingest/postingest.go
- kind: code
  role: implementation
  target: core/jobs/jobs.go
- kind: doc
  role: context
  target: docs/formats.mdx
---

# Process a document

After a document is added, Suchi reads it without changing the original: it extracts the text, recognizes text in scans, detects languages and makes a searchable copy and a thumbnail. A scan with QR separator sheets is split into one document per part.
