---
appliesTo:
- type: capability
  id: upload-replacement-version
- type: capability
  id: view-document
references:
- kind: code
  role: implementation
  target: core/api/versions.go#UploadNewVersion
- kind: doc
  role: intent
  target: docs/document-versions-and-links.mdx
---

# A replacement version keeps the document's filing and access but not its version-specific details

Filing system, owner, category, sensitivity, title, access, tags a person or rule assigned and correspondents carry over; extracted text, model tags, dates, custom field values, links, notes and the previous archive number start fresh.
