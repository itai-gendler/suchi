---
appliesTo:
- type: entity
  id: saved-view
  effect: changes
  facts:
  - Shared
permits:
- actors:
  - administrator
- actors:
  - member
  when:
  - entity: member
    fact: Share saved views
    is: 'On'
references:
- kind: code
  role: implementation
  target: core/api/saved_views.go#UpdateSavedView
- kind: code
  role: implementation
  target: core/api/saved_views.go#CreateSavedView
---

# Only administrators and members with Share saved views share a view

Sharing a view shows it to everyone who may enter the filing system; each still sees only the documents they may open.
