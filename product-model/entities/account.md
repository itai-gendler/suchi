---
references:
- kind: code
  role: implementation
  target: core/api/users.go#PatchUser
- kind: code
  role: implementation
  target: core/api/users_email.go#PostSelfEmail
---

# Account

The sign-in identity and profile of one person, whichever role they hold. Its number never changes, so documents, memberships, grants, mailboxes and share links stay attached when the email changes.

## Information kept

- **Email** — the canonical sign-in address, unique on the server
- **Display name** — the name shown to others
- **Photo** — an optional profile picture
- **Language** — the chosen interface language, or Automatic to follow the browser
- **Sign-in method** — local password or the identity provider the account is bound to
- **Password** — the local password, for an account that signs in with one

## States

### Active

The person may sign in and act.

### Disabled

The person cannot sign in; their mailboxes are not polled and their share links stop working.
