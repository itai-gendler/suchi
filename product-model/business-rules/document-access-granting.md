---
appliesTo:
- type: entity
  id: document-access
  effect: creates
permits:
- actors:
  - administrator
- related:
  - verb: has
    entity: document
  - verb: owns
    entity: member
references:
- kind: code
  role: implementation
  target: core/api/acls.go#PutGrant
- kind: code
  role: implementation
  target: core/authz/store.go#CanManage
---

# Only the owner or an administrator grants access to a document

Access granted to someone never lets them grant access further. The person receiving access must be active and able to enter the document's filing system.
