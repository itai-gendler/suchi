---
appliesTo:
- type: capability
  id: add-note
- type: capability
  id: edit-note
references:
- kind: code
  role: implementation
  target: core/api/document_notes.go
---

# A note holds at most 16 KiB

Longer text is refused.
