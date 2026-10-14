---
appliesTo:
- type: capability
  id: edit-custom-field
- type: capability
  id: set-custom-field-value
references:
- kind: code
  role: implementation
  target: core/api/custom_fields.go#UpdateCustomFieldDef
---

# A custom field's type is fixed once documents have values, and a choice in use stays

Adding choices and renaming the field remain possible.
