---
appliesTo:
- type: capability
  id: import-from-watched-folder
- type: context
  context:
    place: suchi-web::archive::inbox
references:
- kind: code
  role: implementation
  target: core/ingest/fswatch/scanner.go
---

# Suchi never writes to the watched folder

It does not move, rename or delete files or sidecars there.
