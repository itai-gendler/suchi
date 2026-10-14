---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/decrypt.go#DecryptDocument
- kind: code
  role: implementation
  target: core/api/decrypt.go#DecryptBatch
---

# Unlock a document

A person opens a password-protected document with its password. Suchi keeps an unlocked working copy beside the protected original and reads it like any other document.
