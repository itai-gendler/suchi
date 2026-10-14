---
relations:
- {entity: document, verb: shares, cardinality: many-to-many}
references:
- kind: code
  role: implementation
  target: core/api/share_links.go#CreateShareLink
- kind: doc
  role: intent
  target: docs/public-sharing.mdx
---

# Share link

A revocable link that lets anyone holding it download one document or a bundle of documents from one filing system without an account.

## Information kept

- **Documents** — the documents the link serves
- **Surface** — Private network, on the normal archive address, or Public internet, on the operator's separate share address
- **Expiry** — when the link stops working; required for a Public internet link
- **Password** — an optional password the recipient must enter
- **Label** — an optional label
- **Created** — when it was created and by whom

## States

### Active

The link serves its documents until it expires.

### Revoked

The link no longer serves anything.
