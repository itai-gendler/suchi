---
appliesTo:
- type: entity
  id: note
  effect: changes
permits:
- actors:
  - administrator
- related:
  - verb: writes
    entity: member
references:
- kind: code
  role: implementation
  target: core/api/document_notes.go#UpdateDocumentNote
---

# Only a note's author or an administrator edits it

Administrators may moderate any note; the author must still be able to change the document.
