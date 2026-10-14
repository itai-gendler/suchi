---
singleton: true
references:
- kind: code
  role: implementation
  target: core/api/setup.go#SaveLLMSettings
- kind: doc
  role: intent
  target: docs/llm-classifier.mdx
---

# Archive settings

The server-wide processing settings an administrator keeps under Classification.

## Information kept

- **Application mode** — Apply when confident, the default, or Review first
- **Model confidence threshold** — the score from 0.50 to 0.95 at which model metadata and dates may apply; 0.70 by default
- **Similar documents matching** — whether local archive matching suggests metadata
- **Automatic match threshold** — the matching score at which archive matches may apply; 0.90 by default
- **Review match threshold** — the minimum matching score for a review suggestion; 0.50 by default
- **Document model** — the optional local or hosted model, its endpoint and model name
- **Hosted model acknowledgement** — the administrator's acknowledgement that document text leaves the server for a hosted model
