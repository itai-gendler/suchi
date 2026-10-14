---
appliesTo:
- type: capability
  id: upload-replacement-version
- type: capability
  id: move-document-to-trash
references:
- kind: code
  role: implementation
  target: core/api/versions.go#UploadNewVersion
---

# Only the latest version of a document can be replaced

A newer version in Trash still holds its place until it is restored or deleted, so version history never branches.
