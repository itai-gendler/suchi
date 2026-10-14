---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: core/api/documents_list.go#ListDocuments
- kind: code
  role: implementation
  target: ui/src/routes/Documents.svelte
- kind: doc
  role: context
  target: docs/query-language.mdx
---

# Browse documents

A person lists the documents they may open in the current filing system, in Documents, in the Inbox or through a saved view, and narrows them with filters, sorting and pages. By default each version family shows only its newest version the person may see.
