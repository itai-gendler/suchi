---
domain: filing-tree
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/taxonomy.go#ImportTaxonomy
- kind: code
  role: implementation
  target: core/jd/importer/importer.go
- kind: code
  role: implementation
  target: ui/src/lib/TaxonomyImport.svelte
- kind: doc
  role: intent
  target: docs/jd.mdx
---

# Import a filing tree

An administrator imports a filing tree from a HuML or TOML file or pasted text. Imports add to the current tree; the first import that names a filing system code introduces filing systems.
