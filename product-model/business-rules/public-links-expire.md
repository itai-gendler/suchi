---
appliesTo:
- type: capability
  id: create-share-link
- type: capability
  id: open-share-link
references:
- kind: code
  role: implementation
  target: core/api/share_links.go#CreateShareLink
- kind: doc
  role: intent
  target: docs/public-sharing.mdx
---

# A Public internet link always expires, within one year

Private network links may also have an expiry or none. A link's surface never changes; revoke it and create another.
