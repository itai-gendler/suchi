---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/share_links.go#CreateShareLink
- kind: code
  role: implementation
  target: ui/src/lib/ShareLinkForm.svelte
- kind: doc
  role: intent
  target: docs/public-sharing.mdx
---

# Create a share link

A person with Create share links shares one of their documents, or a selected bundle, with someone who has no account. A Private network link uses the normal archive address and may never expire. A Public internet link exists only when the operator enabled public sharing; it uses the separate share address and always expires.
