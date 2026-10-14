---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: plugins/llm-classifier/handler.go
- kind: code
  role: implementation
  target: core/automations/archive_classifier.go
- kind: code
  role: implementation
  target: core/approvals/document_change.go
- kind: doc
  role: intent
  target: docs/llm-classifier.mdx
---

# Suggest document details

After a document is read, local archive matching compares it with already-filed documents and, when an administrator connected one, the document model is asked for a title, category, correspondent, tags, language and up to three important dates. Confident results apply when the application mode is Apply when confident; everything else waits in Approvals for the document owner.
