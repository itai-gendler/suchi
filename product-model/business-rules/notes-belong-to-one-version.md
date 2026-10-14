---
appliesTo:
- type: capability
  id: add-note
- type: capability
  id: upload-replacement-version
references:
- kind: code
  role: implementation
  target: core/api/document_notes.go#CreateDocumentNote
---

# A note belongs to one exact document version

A replacement version starts with no notes, and the earlier version keeps its own.
