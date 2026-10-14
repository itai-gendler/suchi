---
appliesTo:
- type: entity
  id: document-date
  effect: changes
  to: Accepted
permits:
- actors:
  - administrator
- actors:
  - member
  when:
  - entity: member
    fact: Manage document dates
    is: 'On'
references:
- kind: code
  role: implementation
  target: core/api/intelligence.go#ResolveIntelligence
---

# Only administrators and members with Manage document dates accept dates

The reviewer must also be able to read the document, and the date's text must still match it.
