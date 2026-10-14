---
type: web
actors:
- member
- administrator
entryPoints:
- web: /
- web: /login
references:
- kind: code
  role: implementation
  target: ui/src/App.svelte
- kind: code
  role: implementation
  target: core/ui/ui.go
- kind: doc
  role: intent
  target: docs/web-app.mdx
- kind: doc
  role: context
  target: docs/spa-architecture.mdx
---

# Suchi web app

The web app served by the Suchi server: the sign-in and first-run setup pages, the archive people work in, and Archive configuration for administrators.
