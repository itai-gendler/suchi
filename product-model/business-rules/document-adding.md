---
appliesTo:
- type: entity
  id: document
  effect: creates
permits:
- actors:
  - administrator
- actors:
  - member
- unattended: true
references:
- kind: code
  role: implementation
  target: core/api/documents.go#UploadDocument
- kind: code
  role: implementation
  target: core/api/versions.go#UploadNewVersion
---

# Members and administrators add documents to the filing systems they may enter

A member or administrator adds documents only to a filing system they may enter, and becomes their owner; the watched folder and mailboxes add documents for the owner configured for them. A replacement version also needs change access to the version it replaces.
