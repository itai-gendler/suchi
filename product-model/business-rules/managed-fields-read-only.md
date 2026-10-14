---
appliesTo:
- type: capability
  id: edit-custom-field
- type: capability
  id: delete-custom-field
references:
- kind: code
  role: implementation
  target: core/api/custom_fields.go
---

# Custom fields that extraction manages cannot be changed or deleted

Their values on documents stay editable for corrections.
