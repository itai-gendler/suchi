---
domain: filing-tree
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/setup.go#PreviewPresetChange
- kind: code
  role: implementation
  target: core/api/setup.go#ApplyPresetChange
- kind: code
  role: implementation
  target: ui/src/routes/FilingTreeSettings.svelte
- kind: doc
  role: intent
  target: docs/jd.mdx
---

# Choose a filing tree

An administrator chooses the filing tree of the current filing system: a ready-made tree (solo, household, freelance or small business), one composed from focused sets per decade, or Blank. Every change is previewed first and applied only as previewed. Choosing a tree completes archive setup.
