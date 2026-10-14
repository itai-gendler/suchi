---
availability:
- {place: 'suchi-web::sign-in'}
references:
- kind: code
  role: implementation
  target: plugins/local-auth/handlers.go#SetupFormHandler
- kind: code
  role: implementation
  target: plugins/local-auth/localauth.go#mintSetupToken
- kind: doc
  role: context
  target: README.md
---

# Set up the first administrator

On a new server that signs people in with a local password, the first visitor creates the first administrator with the one-time setup token the server writes to its log. The token expires after 24 hours. Setup exists only while the server signs people in with a local password; with an identity provider the operator names the first administrator's email instead.
