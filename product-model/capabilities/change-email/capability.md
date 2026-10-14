---
domain: profile
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/users_email.go#PostSelfEmail
- kind: code
  role: implementation
  target: plugins/oidc/oidc.go
- kind: doc
  role: intent
  target: docs/web-app.mdx
---

# Change email

A person changes the email they sign in with. The account keeps its number, so everything attached to it stays attached. A local-password account types the new address; an account bound to the identity provider takes the provider's current verified email instead.
