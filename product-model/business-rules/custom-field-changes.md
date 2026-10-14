---
appliesTo:
- type: entity
  id: custom-field
  effect: changes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/custom_fields.go#UpdateCustomFieldDef
---

# Only administrators change custom fields

Members only set values of existing fields.
