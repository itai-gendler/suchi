---
appliesTo:
- type: entity
  id: account
  effect: changes
  facts:
  - Email
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
  target: core/api/users_email.go#PostSelfEmail
---

# Only the person themselves changes their sign-in email

Administrators never change another person's email, bind an identity or merge accounts through an email change.
