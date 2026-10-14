---
entities:
- entity: account
  collects:
  - Email
  - Password
references:
- kind: code
  role: implementation
  target: ui/src/routes/Login.svelte
---

# Sign in

Where a person signs in with email and password, or starts sign-in at the identity provider.
