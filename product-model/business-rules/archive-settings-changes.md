---
appliesTo:
- type: entity
  id: archive-settings
  effect: changes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/setup.go#SaveLLMSettings
---

# Only administrators change classification settings

The settings are server-wide.
