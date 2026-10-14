---
references:
- kind: code
  role: implementation
  target: core/api/acls.go#PutGrant
- kind: code
  role: implementation
  target: core/authz/store.go#Grant
---

# Document access

Access to one document granted to a person or a group beyond its owner and administrators.

## Information kept

- **Person or group** — who receives the access
- **Access level** — View, Edit or Full control; Full control includes moving the document to Trash
