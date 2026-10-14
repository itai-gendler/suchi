---
appliesTo:
- type: entity
  id: saved-view
  effect: changes
permits:
- related:
  - verb: owns
    entity: member
- related:
  - verb: owns
    entity: administrator
references:
- kind: code
  role: implementation
  target: core/api/saved_views.go#UpdateSavedView
---

# Only a view's owner changes it

Shared views are read-only for everyone else, administrators included.
