---
entities:
- entity: account
  collects:
  - Email
  - Display name
  - Password
references:
- kind: code
  role: implementation
  target: core/ui/ui.go
---

# First administrator setup

Where the first visitor of a new server creates the first administrator with the setup token.
