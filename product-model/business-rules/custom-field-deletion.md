---
appliesTo:
- type: entity
  id: custom-field
  effect: removes
permits:
- actors:
  - administrator
references:
- kind: code
  role: implementation
  target: core/api/custom_fields.go#DeleteCustomFieldDef
---

# Only administrators delete custom fields

Members only set values of existing fields.
