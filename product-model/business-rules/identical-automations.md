---
appliesTo:
- type: capability
  id: create-automation
- type: capability
  id: edit-automation
references:
- kind: code
  role: implementation
  target: core/automations/signature.go
---

# Two automations of a filing system never do the same thing

Automations with the same triggers and the same actions in the same order are one rule; a different name, position or enabled state does not make another.
