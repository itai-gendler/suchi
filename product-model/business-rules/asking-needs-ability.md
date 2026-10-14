---
appliesTo:
- type: capability
  id: ask-documents
- type: context
  context:
    place: suchi-web::archive::ask-your-documents
references:
- kind: code
  role: implementation
  target: core/api/chat.go#PostChat
---

# Only administrators and members with Ask your documents ask questions

Questions read only documents the person may open and are kept only in their browser.
