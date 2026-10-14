---
entities:
- entity: account
  shows:
  - Email
  - Display name
  - Photo
  - Language
  - Sign-in method
  collects:
  - Display name
  - Photo
  - Language
  - Email
references:
- kind: code
  role: implementation
  target: ui/src/routes/AccountSettings.svelte
---

# Profile

The person's own account under My account: display name, photo, language and sign-in email.
