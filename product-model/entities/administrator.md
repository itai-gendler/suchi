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
references:
- kind: doc
  role: intent
  target: docs/permissions.mdx
- kind: code
  role: implementation
  target: core/api/authz_helpers.go#requireAdmin
---

# Administrator

A person who administers the archive: the first account created on a new server, or anyone an administrator creates with that role. Administrators have every additional access implicitly, may enter every filing system, open every document and configure the archive in Archive configuration.
