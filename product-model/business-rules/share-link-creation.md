---
appliesTo:
- type: entity
  id: share-link
  effect: creates
permits:
- actors:
  - administrator
- actors:
  - member
  when:
  - entity: member
    fact: Create share links
    is: 'On'
references:
- kind: code
  role: implementation
  target: core/api/share_links.go#CreateShareLink
- kind: code
  role: implementation
  target: core/api/share_links.go#assertShareable
---

# Only administrators and members with Create share links create share links

A member shares only live documents they own, from one filing system per link; an administrator may share any live document of the filing system.
