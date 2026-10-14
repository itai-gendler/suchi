---
appliesTo:
- type: capability
  id: upload-documents
- type: capability
  id: process-document
- type: capability
  id: unlock-document
- type: capability
  id: upload-replacement-version
references:
- kind: code
  role: implementation
  target: core/blob/cas.go
- kind: doc
  role: intent
  target: README.md
---

# A document's original file is never modified

Text, searchable copies, unlocked copies, split parts and replacement versions are kept beside the original; none of them rewrites it.
