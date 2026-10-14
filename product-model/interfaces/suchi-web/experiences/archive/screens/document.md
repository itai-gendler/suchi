---
entities:
- entity: document
  shows:
  - Title
  - Title ownership
  - Filing category
  - Filing address
  - Sensitivity
  - Tags
  - Correspondents
  - Custom field values
  - Languages
  - Extracted text
  - Original file
  - Searchable copy
  - Added from
  - Added
  - Original file date
  - Previous version
  - Previous archive number
  - Password protection
  - Similar documents
- entity: note
  shows:
  - Text
  - Author
  - Created
  - Edited
  collects:
  - Text
- entity: document-link
  shows:
  - Link type
  - Target document
- entity: document-access
  shows:
  - Person or group
  - Access level
- entity: share-link
  shows:
  - Surface
  - Expiry
  - Password
  - Created
- entity: custom-field
  shows:
  - Name
  - Type
- entity: tag
  shows:
  - Name
  - Slug
- entity: correspondent
  shows:
  - Name
- entity: category
  shows:
  - Code
  - Name
  - Area
- entity: group
  shows:
  - Name
- entity: filing-system
  shows:
  - Members
references:
- kind: code
  role: implementation
  target: ui/src/routes/DocumentDetail.svelte
---

# Document

One exact document version: its preview and metadata, extracted text, similar and linked documents, versions, notes, access and share links. A document in Trash opens here read-only.
