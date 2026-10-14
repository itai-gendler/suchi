---
appliesTo:
- type: entity
  id: saved-view
  effect: removes
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
  target: core/api/saved_views.go#DeleteSavedView
---

# Only a view's owner deletes it

Shared views cannot be deleted by anyone else, administrators included.
