---
appliesTo:
- type: capability
  id: classify-document
- type: capability
  id: change-classification-settings
references:
- kind: code
  role: implementation
  target: core/settings/settings.go
- kind: doc
  role: intent
  target: docs/llm-classifier.mdx
---

# Inferred changes apply by themselves only in Apply when confident and at or above their threshold

Review first sends every inferred change and date to Approvals, even at full confidence. Archive matching applies at 0.90 and suggests from 0.50 by default; the document model applies at 0.70 by default; each date uses its own score. Explicit automations are independent of the mode.
