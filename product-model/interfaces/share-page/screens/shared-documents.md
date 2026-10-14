---
entities:
- entity: share-link
  shows:
  - Expiry
  - Label
  - Created
- entity: document
  shows:
  - Title
  - Original file
references:
- kind: code
  role: implementation
  target: core/api/share_links.go#renderShareHTML
---

# Shared documents

The documents a share link serves, its expiry and, when it has one, the password form.
