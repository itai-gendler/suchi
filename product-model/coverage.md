---
scope: 'The Suchi server and its web app: document capture, processing, filing, retrieval, review, sharing, permissions and archive configuration.'
method: Static inspection of source, tests and documentation; no code was run.
covered:
- description: HTTP handlers behind the web app and share pages
  paths:
  - core/api/
- description: Approval runs for document-change suggestions
  paths:
  - core/approvals/
- description: Authorization, groups, access grants and member abilities
  paths:
  - core/authz/
  - core/auth/
- description: Automation rules and their actions
  paths:
  - core/automations/
- description: Watched-folder, sidecar and mailbox intake
  paths:
  - core/ingest/
  - core/emailaccounts/
- description: Post-ingest processing and job dispatch
  paths:
  - core/pipeline/postingest/
  - core/jobs/
- description: Filing trees, filing systems and taxonomy import and export
  paths:
  - core/jd/
  - core/taxonomy/
- description: Trash retention and purge
  paths:
  - core/trash/
- description: Extracted dates, query language, typed custom fields, archive matching and inference guards
  paths:
  - core/intelligence/
  - core/searchquery/
  - core/customfield/
  - core/similar/
  - core/documentstate/
- description: Server settings and setup state
  paths:
  - core/settings/
- description: Password and identity-provider sign-in
  paths:
  - plugins/local-auth/
  - plugins/oidc/
- description: Document model classification
  paths:
  - plugins/llm-classifier/
- description: The web app and server-rendered pages
  paths:
  - ui/src/
  - core/ui/
- description: Server assembly and route registration
  paths:
  - distro/app/
exclusions:
- description: Documentation, developer tooling, fixtures and benchmarks
  paths:
  - docs/
  - hack/
  - hooks/
- description: Packaging, deployment templates and CI
  paths:
  - deploy/
  - packaging/
  - .github/
  - Dockerfile
  - Makefile
  - compose.yaml
- description: Extension interfaces for third-party plugins
  paths:
  - plugin-api/
- description: Storage, schema, logging, HTTP plumbing, translations and assets
  paths:
  - core/db/
  - core/blob/
  - core/httpx/
  - core/logx/
  - core/i18n/
  - core/uiassets/
  - core/config/
  - core/crypto/
  - core/netutil/
  - core/slug/
  - core/mimeutil/
  - core/sandbox/
  - core/lang/
- description: Browser end-to-end tests
  paths:
  - ui/tests/
unmapped:
- description: Folder layouts and the rendered folder tree
  paths:
  - core/render/
  - ui/src/routes/FolderLayoutsSettings.svelte
- description: API tokens, mobile app pairing and the client handshake
  paths:
  - core/api/tokens.go
  - core/api/mobile_pairing.go
  - core/api/handshake.go
  - ui/src/lib/MobilePairing.svelte
- description: Rescan proposals for stale processing results
  paths:
  - core/rescan/
- description: Refiling documents against the current tree
  paths:
  - core/refile/
  - core/api/refile.go
- description: Format-specific extraction, text recognition, splitting and e-invoice reading
  paths:
  - core/pipeline/
- description: Backups, restore, takeout export, third-party imports and storage cleanup
  paths:
  - core/backup/
  - core/importer/
  - core/gc/
- description: Command-line tools and the MCP server
  paths:
  - distro/cmd/
  - distro/internal/
- description: Public demo mode
  paths:
  - distro/demo/
  - core/api/demo.go
  - ui/src/routes/Demo.svelte
- description: Activity events and audit retention
  paths:
  - core/audit/
  - core/api/events.go
limitations: []
---

# Coverage
