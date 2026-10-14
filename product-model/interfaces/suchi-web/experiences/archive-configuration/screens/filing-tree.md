---
entities:
- entity: category
  shows:
  - Code
  - Name
  - Description
  - Area
- entity: filing-system
  shows:
  - Filing tree choice
references:
- kind: code
  role: implementation
  target: ui/src/routes/FilingTreeSettings.svelte
---

# Filing tree

Ready-made and composed trees, Blank, and import or export of the current filing system's tree, each previewed before it applies.
