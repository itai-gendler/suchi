---
appliesTo:
- type: capability
  id: open-share-link
- type: capability
  id: disable-user
- type: capability
  id: remove-filing-system-member
- type: capability
  id: revoke-additional-access
references:
- kind: code
  role: implementation
  target: core/api/share_links.go#loadShareByToken
---

# A share link works only while its creator may still share its documents

Disabling the creator, removing them from the filing system or taking away Create share links ends the link; documents later trashed are left out.
