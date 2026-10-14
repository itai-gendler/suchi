---
appliesTo:
- type: entity
  id: document-date
  effect: creates
permits:
- actors:
  - administrator
- actors:
  - member
  when:
  - entity: member
    fact: Manage document dates
    is: 'On'
- unattended: true
references:
- kind: code
  role: implementation
  target: core/api/intelligence.go#ExtractIntelligence
---

# Only administrators and members with Manage document dates extract dates

Dates found while a document is processed are kept for its owner to review.
