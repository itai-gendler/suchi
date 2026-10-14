---
appliesTo:
- type: entity
  id: note
  effect: creates
permits:
- actors:
  - administrator
- related:
  - verb: has
    entity: document
  - verb: owns
    entity: member
- related:
  - verb: has
    entity: document
  - verb: has
    entity: document-access
  - verb: holds
    entity: member
  when:
  - entity: document-access
    fact: Access level
    is-not: View
- related:
  - verb: has
    entity: document
  - verb: has
    entity: document-access
  - verb: holds
    entity: group
  - verb: includes
    entity: member
  when:
  - entity: document-access
    fact: Access level
    is-not: View
references:
- kind: code
  role: implementation
  target: core/api/document_notes.go#CreateDocumentNote
---

# Only someone who may change a document adds a note to it

Adding a note needs the same change access as editing the document.
