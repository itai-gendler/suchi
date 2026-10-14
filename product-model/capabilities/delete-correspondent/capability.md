---
domain: correspondents
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/taxonomy_crud.go#DeleteCorrespondent
---

# Delete a correspondent

A correspondent is deleted from the catalog and from every document that names it.
