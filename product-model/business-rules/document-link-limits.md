---
appliesTo:
- type: capability
  id: add-document-link
- type: capability
  id: change-document-link
- type: capability
  id: upload-replacement-version
references:
- kind: code
  role: implementation
  target: core/api/custom_field_values.go#SetCustomField
- kind: doc
  role: intent
  target: docs/document-versions-and-links.mdx
---

# A document never links to itself or to its own versions, and a replacement never copies a link

Distinct documents may link both ways or in a cycle; only direct links are shown.
