---
appliesTo:
- type: capability
  id: upload-documents
- type: context
  context:
    place: suchi-web::archive::upload-queue
references:
- kind: code
  role: implementation
  target: core/api/documents.go#UploadDocument
---

# Uploading a file a person already keeps returns that document

Identical bytes from the same owner in the same filing system make no second document; a matching document in Trash comes back instead.
