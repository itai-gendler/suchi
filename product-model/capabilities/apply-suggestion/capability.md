---
domain: approvals
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/approvals.go#ApprovalResolveTask
- kind: code
  role: implementation
  target: core/approvals/runner.go
- kind: code
  role: implementation
  target: ui/src/routes/Tasks.svelte
- kind: doc
  role: intent
  target: docs/approvals.mdx
---

# Apply a suggestion

The document owner, or an administrator, applies a suggested change from Approvals. The card shows the current and suggested values, why it was suggested and, on request, the model confidence and who suggested it; its button names the change, such as File document, Set correspondent, Add tag or Change title. Applying records the decision and queues the change, which is checked again before it is written.
