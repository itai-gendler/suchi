---
appliesTo:
- type: capability
  id: file-document
- type: capability
  id: rename-document
- type: capability
  id: add-tag-to-document
- type: capability
  id: change-document-sensitivity
references:
- kind: doc
  role: intent
  target: docs/web-app.mdx
- kind: doc
  role: intent
  target: docs/llm-classifier.mdx
---

# A person's edit removes the needs-review tag the classifier added

Only the classifier's own marker goes; a needs-review tag a person, rule or import added stays until removed. Trash, restore and rescan do not count as review.
