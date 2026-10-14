---
appliesTo:
- type: capability
  id: import-filing-tree
- type: context
  context:
    place: suchi-web::archive-configuration::filing-tree
references:
- kind: doc
  role: intent
  target: docs/jd.mdx
---

# Imports add to the current filing tree

Only a new archive with no filed or recoverable documents and no dependent configuration may have its first tree replaced; later imports never remove categories.
