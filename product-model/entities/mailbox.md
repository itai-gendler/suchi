---
references:
- kind: code
  role: implementation
  target: core/api/emailaccounts.go#CreateEmailAccount
- kind: doc
  role: intent
  target: docs/guides/mail-intake.mdx
---

# Mailbox

A person's own IMAP mailbox that Suchi polls for documents; its owner owns every document it imports.

## Information kept

- **Name** — the display name
- **Provider** — one of the built-in mail provider presets, or custom IMAP
- **Username** — the mailbox sign-in name
- **Credentials** — an app password, or a sign-in at the provider for providers that require it
- **Source folder** — the folder polled
- **Intake rules** — one to twenty rules choosing which messages to keep and whether to keep the email with its files or the files only
- **After import** — leave messages unchanged, mark them read, or move them to another folder
- **Sync mail from** — the earliest message date imported
- **Poll interval** — minutes between polls, 1 to 1440
- **Enabled** — whether it is polled
- **Sync status** — never synced, the last successful sync, or the last error
