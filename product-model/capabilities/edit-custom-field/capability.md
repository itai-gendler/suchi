---
domain: custom-fields
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/custom_fields.go#UpdateCustomFieldDef
---

# Edit a custom field

An administrator renames a custom field or adds choices to it. Its type is fixed once documents have values, and a choice in use cannot be removed.
