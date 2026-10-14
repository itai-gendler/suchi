---
appliesTo:
- type: entity
  id: suggestion
  effect: changes
  to: Outdated
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
- unattended: true
references:
- kind: code
  role: implementation
  target: core/approvals/document_change.go#ReconcileDocumentChanges
---

# A suggestion closes as outdated only through a change by someone who may change its document

A person's edit, a rule or a new reading of the document closes a waiting suggestion for the same field.
