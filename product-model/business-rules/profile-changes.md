---
appliesTo:
- type: entity
  id: account
  effect: changes
  facts:
  - Display name
  - Photo
  - Language
permits:
- related:
  - verb: has
    entity: member
- related:
  - verb: has
    entity: administrator
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchSelf
---

# Only the person themselves changes their profile

Display name, photo and language belong to the account holder.
