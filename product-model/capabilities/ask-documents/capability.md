---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/chat.go#PostChat
- kind: code
  role: implementation
  target: ui/src/lib/ArchiveChat.svelte
- kind: doc
  role: intent
  target: docs/archive-chat.mdx
---

# Ask your documents

A person with Ask your documents asks a question across the documents they may read; the document model answers with links to the documents it used. Questions change nothing and the conversation is kept only in the browser.
