---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/ingest/fswatch/scanner.go
- kind: code
  role: implementation
  target: core/ingest/sidecar/sidecar.go
- kind: doc
  role: context
  target: docs/architecture.mdx
---

# Import from the watched folder

Suchi polls the folder the operator configured and adds new or changed files as documents owned by the configured watched-folder owner. It only reads the folder: it never moves, renames, deletes or writes files or sidecars. A sidecar file beside a document may supply its title, tags, correspondent and existing custom fields.
