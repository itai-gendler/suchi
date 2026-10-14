---
references:
- kind: code
  role: implementation
  target: core/api/tasks.go#RetryDeadJob
- kind: code
  role: implementation
  target: core/jobs/jobs.go
---

# Processing issue

A processing step for a document that stopped after its retries, shown to administrators under Processing issues.

## Information kept

- **Processing step** — which step stopped, such as text recognition or classification
- **Error** — why it stopped
- **Retries** — how many times it was tried
