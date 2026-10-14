---
availability:
- {place: 'suchi-web::sign-in'}
references:
- kind: code
  role: implementation
  target: plugins/local-auth/handlers.go#LoginFormHandler
- kind: code
  role: implementation
  target: plugins/oidc/oidc.go#resolveOrProvisionUser
- kind: doc
  role: intent
  target: docs/permissions.mdx
---

# Sign in

A person signs in to reach the filing systems they may enter. Whether they type a password or use the operator's identity provider depends on how the server is configured.
