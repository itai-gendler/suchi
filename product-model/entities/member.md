---
kind: person
acts: external
relations:
- {entity: account, verb: has, cardinality: one-to-one}
- {entity: document, verb: owns, cardinality: one-to-many}
- {entity: note, verb: writes, cardinality: one-to-many}
- {entity: share-link, verb: creates, cardinality: one-to-many}
- {entity: saved-view, verb: owns, cardinality: one-to-many}
- {entity: mailbox, verb: owns, cardinality: one-to-many}
- {entity: document-access, verb: holds, cardinality: one-to-many}
references:
- kind: doc
  role: intent
  target: docs/permissions.mdx
- kind: code
  role: implementation
  target: core/authz/capabilities.go
---

# Member

A person who keeps and finds documents in the filing systems they may enter. A member owns the documents they add, opens other documents only through access granted to them or their groups, and uses the optional abilities an administrator grants them under Additional access.

## Information kept

- **Ask your documents** — On when the member may ask questions across the documents they can read
- **Manage document dates** — On when the member may extract, review and use document dates and open Calendar
- **Manage mailboxes** — On when the member may connect and manage their own mailboxes
- **Create share links** — On when the member may create share links for their own documents
- **Share saved views** — On when the member may share their saved views with others in the filing system
