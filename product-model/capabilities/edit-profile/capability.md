---
domain: profile
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchSelf
- kind: code
  role: implementation
  target: core/api/users_avatar.go#PostSelfAvatar
- kind: code
  role: implementation
  target: ui/src/routes/AccountSettings.svelte
---

# Edit profile

A person changes the display name, photo and interface language of their own account under My account > Profile. Choosing Automatic follows the browser's language again.
