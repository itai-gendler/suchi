---
relations:
- {entity: note, verb: has, cardinality: one-to-many}
- {entity: document-access, verb: has, cardinality: one-to-many}
- {entity: suggestion, verb: has, cardinality: one-to-many}
- {entity: document-date, verb: has, cardinality: one-to-many}
- {entity: document-link, verb: links, cardinality: one-to-many}
- {entity: processing-issue, verb: has, cardinality: one-to-many}
- {entity: tag, verb: carries, cardinality: many-to-many}
- {entity: correspondent, verb: involves, cardinality: many-to-many}
references:
- kind: code
  role: implementation
  target: core/api/documents.go#GetDocument
- kind: doc
  role: intent
  target: docs/document-versions-and-links.mdx
---

# Document

One exact file kept in the archive, with its extracted text and the metadata people and automations give it. A replacement version is a new document that names the one it replaces; both stay available. Each document is addressed by its installation-wide number and, once filed, by its filing address such as `S01.13.147`.

## Information kept

- **Title** — the name shown everywhere; starts as the file name or a generated title
- **Title ownership** — Automatic, set by a rule, or set by a person; suggestions and rules never change a title a person set
- **Filing category** — the one category of the filing tree the document is filed under
- **Filing address** — the filing system code, category code and document number, such as S01.13.147
- **Sensitivity** — Unset, Public, Internal, Confidential or Restricted
- **Tags** — the existing tags assigned to the document
- **Correspondents** — the correspondents attached to the document, each with its role: sender, recipient, cc or other
- **Custom field values** — the typed values of ordinary custom fields set on the document
- **Languages** — the detected or chosen languages of the text
- **Extracted text** — the text read from the file, by direct extraction or text recognition
- **Original file** — the unchanged original bytes, their format and size
- **Searchable copy** — the searchable PDF or preview rendition derived from the original
- **Added from** — how the file entered the archive: upload, watched folder, mailbox or import
- **Added** — when the document was added
- **Original file date** — the date carried from the original file, when known
- **Previous version** — the document this one replaces, for a replacement version
- **Previous archive number** — a read-only number retained from an import
- **Password protection** — None, Locked while it waits for its password, or Password unlocked
- **Similar documents** — already-filed documents whose words match this one

## States

### Inbox

Filed under the protected System Inbox, waiting for a filing decision.

### Filed

Filed under a category of the filing tree.

### Trashed

In Trash: read-only, recoverable for 30 days, then permanently deleted.
