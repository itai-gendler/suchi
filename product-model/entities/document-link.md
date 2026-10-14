---
references:
- kind: code
  role: implementation
  target: core/api/custom_field_values.go#SetCustomField
- kind: code
  role: implementation
  target: core/api/document_references.go#ListReferencedBy
---

# Document link

A named relationship from one exact document to another exact document, such as an invoice pointing to its payment receipt. Both documents stay peers; the link is shown on the document that holds it and, as a backlink, on the target.

## Information kept

- **Link type** — the Document link custom field that names the relationship
- **Target document** — the exact document version linked to
