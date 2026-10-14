---
type: web
actors:
- link-recipient
entryPoints:
- web: /s/{token}
references:
- kind: code
  role: implementation
  target: core/api/share_links.go#RegisterShareSurface
- kind: doc
  role: intent
  target: docs/public-sharing.mdx
---

# Share page

The page a share link opens, on the normal archive address for a Private network link or on the operator's separate public share address for a Public internet link. It needs no account.
