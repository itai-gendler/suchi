---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/ingest/emailwatch/emailwatch.go
- kind: code
  role: implementation
  target: core/ingest/emailwatch/allowlist.go
- kind: doc
  role: intent
  target: docs/guides/mail-intake.mdx
---

# Import mail

Suchi polls each enabled mailbox on its interval and keeps the messages its intake rules accept, as documents owned by the mailbox owner. A message every rule rejects never enters the archive and is left untouched on the mail server.
