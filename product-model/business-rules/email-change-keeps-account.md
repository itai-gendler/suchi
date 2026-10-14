---
appliesTo:
- type: capability
  id: change-email
- type: context
  context:
    place: suchi-web::archive::profile
references:
- kind: code
  role: implementation
  target: core/api/users_email.go#PostSelfEmail
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Changing email keeps the same account and signs out every other browser

Documents, memberships, access, tokens and connected phones stay attached; administrators never change another person's email.
