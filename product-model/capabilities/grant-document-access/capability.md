---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/acls.go#PutGrant
- kind: code
  role: implementation
  target: core/authz/store.go#CanManage
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Grant document access

A document's owner or an administrator lets another person, or a group, open the document with View, Edit or Full control.
