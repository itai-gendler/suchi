---
appliesTo:
- type: capability
  id: change-classification-settings
- type: capability
  id: classify-document
- type: capability
  id: ask-documents
references:
- kind: doc
  role: intent
  target: docs/llm-classifier.mdx
- kind: doc
  role: context
  target: docs/privacy.mdx
---

# Document text goes to a hosted model only after an administrator acknowledged it

A model on the local network needs no acknowledgement. Only extracted text and bounded names are sent, never original files.
