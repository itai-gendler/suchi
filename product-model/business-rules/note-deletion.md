---
appliesTo:
- type: entity
  id: note
  effect: removes
permits:
- actors:
  - administrator
- related:
  - verb: writes
    entity: member
references:
- kind: code
  role: implementation
  target: core/api/document_notes.go#DeleteDocumentNote
---

# Only a note's author or an administrator deletes it

Administrators may moderate any note; the author must still be able to change the document.
