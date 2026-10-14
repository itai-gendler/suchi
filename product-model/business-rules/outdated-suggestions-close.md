---
appliesTo:
- type: capability
  id: apply-suggestion
- type: capability
  id: rename-document
- type: capability
  id: file-document
references:
- kind: code
  role: implementation
  target: core/approvals/document_change.go#ReconcileDocumentChanges
---

# A suggestion closes as outdated when its document or field changes first

A person's edit, a rule or a new reading of the document makes a waiting suggestion for that field outdated; it leaves Approvals without anyone dismissing it.
