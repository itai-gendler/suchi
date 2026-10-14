---
appliesTo:
- type: capability
  id: run-automations
- type: capability
  id: classify-document
- type: capability
  id: rename-document
- type: capability
  id: release-document-title
references:
- kind: code
  role: implementation
  target: core/api/documents.go#PatchDocument
- kind: doc
  role: context
  target: docs/architecture.mdx
---

# Suggestions and automations never change a title a person set

Renaming a document or applying a title suggestion makes the title the person's own; an automation may refresh a title a rule set; Let Suchi rename this hands the title back.
