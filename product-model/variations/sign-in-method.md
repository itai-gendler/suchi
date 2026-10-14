---
kind: configuration
of: capability-scenario
takesEffect: 'When the server starts: the operator''s identity-provider configuration is read at startup and applies to every sign-in.'
stability: Fixed while the server runs; a changed configuration applies from the next start. Each account keeps the sign-in method it is bound to.
alternatives:
- id: sign-in-with-password
  selectedWhen: The operator configured no identity provider. This is the default for a new server.
- id: sign-in-with-identity-provider
  selectedWhen: The operator configured an identity provider issuer for the server.
references:
- kind: code
  role: implementation
  target: distro/app/serve_routes.go
- kind: code
  role: implementation
  target: core/config/config.go
---

# Sign-in method

Operators either keep local passwords or hand sign-in to their identity provider; both are supported, one per server.
