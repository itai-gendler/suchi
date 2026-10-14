---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/documents_bulk.go#BulkEdit
- kind: code
  role: implementation
  target: core/rescan/enqueue.go
- kind: doc
  role: context
  target: docs/pipeline-versions.mdx
---

# Rescan documents

A person asks Suchi to read and process selected documents again, for example after a better text-recognition setup. Originals are not rewritten.
