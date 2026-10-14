---
appliesTo:
- type: capability
  id: set-up-first-administrator
- type: context
  context:
    place: suchi-web::sign-in::setup
references:
- kind: code
  role: implementation
  target: plugins/local-auth/localauth.go#mintSetupToken
---

# The setup token works once and for 24 hours

An expired token is replaced in the server log when someone tries it.
