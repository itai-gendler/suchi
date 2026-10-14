---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/documents.go#PatchDocument
- kind: code
  role: implementation
  target: core/approvals/document_change.go#ReconcileDocumentChanges
- kind: doc
  role: intent
  target: docs/automations.mdx
---

# Rename a document

A person gives a document a title of their own. A title a person set is never changed by suggestions or automations.
