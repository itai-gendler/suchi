---
availability:
- {place: 'suchi-web::archive'}
references:
- kind: code
  role: implementation
  target: ui/src/routes/Dashboard.svelte
- kind: code
  role: implementation
  target: core/api/stats.go#GetStats
---

# View the dashboard

The Dashboard summarizes the current filing system: totals, the Inbox and Approvals counts, failed jobs, recently added documents and saved views. Until a filing tree is chosen, administrators see the setup step instead of a finished archive.
