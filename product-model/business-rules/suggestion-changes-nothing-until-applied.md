---
appliesTo:
- type: capability
  id: apply-suggestion
- type: capability
  id: dismiss-suggestion
- type: capability
  id: classify-document
references:
- kind: doc
  role: intent
  target: docs/approvals.mdx
---

# A suggestion changes nothing until it is applied

Applying records the decision; the change is written only after the document, the field and the reviewer's access are checked again.
