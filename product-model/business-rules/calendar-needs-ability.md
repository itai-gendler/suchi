---
appliesTo:
- type: capability
  id: browse-calendar
- type: capability
  id: extract-dates
references:
- kind: code
  role: implementation
  target: core/api/intelligence.go
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Only administrators and members with Manage document dates use Calendar and extract dates

Taking the ability away hides Calendar and extraction; dates already accepted stay stored.
