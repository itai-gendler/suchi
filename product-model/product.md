---
id: suchi
summary: 'A self-hosted document archive that files itself: capture scans, mail and photos, extract their text, file them into a numbered category tree and find them again.'
category: document-management
tags:
- document-archive
- document-management
- numbered-filing
- ocr
- self-hosted
authors:
- name: Ritesh Shrivastav
  url: https://github.com/johnnybravo-xyz
license: AGPL-3.0-or-later
languages:
- en
limitations:
- Original files are never modified; every later rendition, unlocked copy or replacement version sits beside the original.
- Every document is filed in exactly one category of its filing tree; a document the archive is not sure about waits in the System Inbox instead of being guessed.
- A language model is optional and only suggests or applies what an administrator allows; sending document text to a hosted model requires explicit acknowledgement.
- People sign in either with a local password or through the operator's identity provider, whichever the server is configured for.
- Revoking a share link stops later access but cannot take back a file the recipient already downloaded.
references:
- kind: code
  role: implementation
  target: README.md
- kind: doc
  role: context
  target: docs/index.mdx
- kind: doc
  role: intent
  target: docs/web-app.mdx
- kind: doc
  role: context
  target: docs/architecture.mdx
---

# Suchi

Suchi is a self-hosted document archive for a person, a household or a small team. Documents arrive by browser upload, from a watched folder or from the owners' own mailboxes; Suchi keeps each original unchanged, extracts its text, gives it a title and files what it is sure about into a numbered tree of areas and categories. Anything it is unsure about waits in the Inbox. People then find documents again by search, filters, saved views, a calendar of the dates found in them and questions answered with links to the documents used, and they keep each document's correspondents, tags, typed fields, notes, links to related documents and replacement versions.

Suggestions from local archive matching, from automations and from an optional language model either apply when confident or wait in Approvals for the document owner. Administrators configure the filing tree, the people and groups who may enter each filing system, the shared metadata catalogs, mail intake, automations and classification; document owners and administrators decide who else may open each document and may share documents with people who have no account through revocable links.

## Intent

A local archive that files itself: capture, extract, file, retrieve and review in one binary with SQLite and files on disk, where nothing calls out until an integration is enabled. One filing model of numbered categories replaces overlapping tags, types and storage paths; deterministic automations run before any model, and uncertain filing waits for a person instead of being guessed. Originals are never modified, human edits win over automatic suggestions, and optional integrations must preserve useful local operation.
