---
appliesTo:
- type: entity
  id: document-link
  effect: creates
permits:
- actors:
  - administrator
- related:
  - verb: links
    entity: document
  - verb: owns
    entity: member
- related:
  - verb: links
    entity: document
  - verb: has
    entity: document-access
  - verb: holds
    entity: member
  when:
  - entity: document-access
    fact: Access level
    is-not: View
- related:
  - verb: links
    entity: document
  - verb: has
    entity: document-access
  - verb: holds
    entity: group
  - verb: includes
    entity: member
  when:
  - entity: document-access
    fact: Access level
    is-not: View
references:
- kind: code
  role: implementation
  target: core/api/custom_field_values.go#SetCustomField
---

# Only someone who may change a document adds a link from it

Adding a link needs change access on the document that holds it and view access on the live target; a target in another filing system needs entry to both.
