---
appliesTo:
- type: entity
  id: custom-field
  effect: creates
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/custom_fields.go#CreateCustomFieldDef
---

# Only administrators create custom fields

Members only set values of existing fields.
