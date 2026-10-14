---
appliesTo:
- type: entity
  id: suggestion
  effect: changes
  to: Dismissed
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
---

# Only the document owner or an administrator dismisses a suggestion

Suggestions are assigned to the document owner.
