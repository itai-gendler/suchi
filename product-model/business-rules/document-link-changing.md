---
appliesTo:
- type: entity
  id: document-link
  effect: changes
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

# Only someone who may change a document changes its links

Changing a link's target needs change access on the document that holds it and view access on the new target.
