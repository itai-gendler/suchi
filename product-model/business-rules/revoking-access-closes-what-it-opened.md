---
appliesTo:
- type: capability
  id: revoke-additional-access
- type: capability
  id: grant-additional-access
references:
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Taking away an additional access closes what it opened

Create share links revokes the member's live links, Share saved views makes their shared views private, and Manage mailboxes disables their mailboxes; granting it again restores none of them.
