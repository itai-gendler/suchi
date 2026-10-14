---
appliesTo:
- type: capability
  id: add-filing-system-member
- type: capability
  id: remove-filing-system-member
- type: capability
  id: grant-document-access
references:
- kind: code
  role: implementation
  target: core/jd/systems/systems.go
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Only administrators and admitted people enter a filing system

Entry grants no access to documents, and access to a document never admits anyone to its filing system.
