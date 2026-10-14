---
appliesTo:
- type: capability
  id: upload-documents
- type: capability
  id: import-from-watched-folder
- type: capability
  id: import-mail
references:
- kind: code
  role: implementation
  target: core/api/documents.go#UploadDocument
- kind: doc
  role: intent
  target: docs/jd.mdx
---

# Every new document starts in the System Inbox

A document leaves the Inbox only when a person, an automation or a confident suggestion files it; anything uncertain waits there instead of being guessed.
