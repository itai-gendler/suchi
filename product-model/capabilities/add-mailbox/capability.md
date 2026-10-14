---
availability:
- {place: 'suchi-web::archive'}
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/emailaccounts.go#CreateEmailAccount
- kind: code
  role: implementation
  target: ui/src/lib/MailAccountForm.svelte
- kind: doc
  role: intent
  target: docs/guides/mail-intake.mdx
---

# Add a mailbox

A member with Manage mailboxes connects their own mailbox under My account > Mailboxes; an administrator adds one for any person under Email intake and chooses who owns the documents it imports.
