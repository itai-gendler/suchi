---
references:
- kind: code
  role: implementation
  target: core/api/document_notes.go#CreateDocumentNote
---

# Note

Attributed working context kept beside one exact document version: a decision, a follow-up or background, without changing the document.

## Information kept

- **Text** — the note itself, at most 16 KiB
- **Author** — who wrote it
- **Created** — when it was written
- **Edited** — when it was last edited
