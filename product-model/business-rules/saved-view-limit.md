---
appliesTo:
- type: capability
  id: create-saved-view
- type: context
  context:
    place: suchi-web::archive::views
references:
- kind: code
  role: implementation
  target: core/api/saved_views.go#CreateSavedView
---

# A person keeps at most 50 saved views per filing system

A further view is refused.
