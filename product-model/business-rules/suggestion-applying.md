---
appliesTo:
- type: entity
  id: suggestion
  effect: changes
  to: Accepted
permits:
- actors:
  - administrator
- related:
  - verb: has
    entity: document
  - verb: owns
    entity: member
references:
- kind: code
  role: implementation
  target: core/approvals/runner.go
- kind: code
  role: implementation
  target: core/approvals/document_change.go
---

# Only the document owner or an administrator applies a suggestion

Suggestions are assigned to the document owner; whoever applies one must still be able to change the document, and the change is checked again before it is written.
