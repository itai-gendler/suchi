---
appliesTo:
- type: entity
  id: share-link
  effect: changes
  to: Revoked
permits:
- related:
  - verb: creates
    entity: member
- related:
  - verb: creates
    entity: administrator
references:
- kind: code
  role: implementation
  target: core/api/share_links.go#RevokeShareLink
---

# Only the person who created a share link revokes it

Administrators do not revoke other people's links one by one; a link also stops working when its creator is disabled, leaves the filing system or loses Create share links.
