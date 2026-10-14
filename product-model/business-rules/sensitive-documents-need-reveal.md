---
appliesTo:
- type: capability
  id: view-document
- type: capability
  id: change-document-sensitivity
references:
- kind: code
  role: implementation
  target: ui/src/routes/DocumentDetail.svelte
- kind: doc
  role: intent
  target: docs/web-app.mdx
---

# Confidential and Restricted documents show their preview and text only after Reveal

Their thumbnails are blurred in lists. Sensitivity never grants or removes access.
