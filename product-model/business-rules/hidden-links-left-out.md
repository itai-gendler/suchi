---
appliesTo:
- type: capability
  id: view-document
- type: capability
  id: restore-document
references:
- kind: code
  role: implementation
  target: core/api/document_references.go#ListReferencedBy
---

# Links to or from documents a person cannot see, or that are in Trash, are left out

Nothing about a hidden document is shown: no title, number or count. Restoring a trashed document shows its links again.
