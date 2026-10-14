---
appliesTo:
- type: entity
  id: document-link
  effect: removes
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
  target: core/api/custom_field_values.go#DeleteCustomField
---

# Only someone who may change a document removes its links

Removing a link needs change access on the document that holds it.
