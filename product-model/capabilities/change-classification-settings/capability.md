---
availability:
- {place: 'suchi-web::archive-configuration'}
references:
- kind: code
  role: implementation
  target: core/api/setup.go#SaveLLMSettings
- kind: code
  role: implementation
  target: core/api/setup.go#SavePreferences
- kind: code
  role: implementation
  target: ui/src/routes/ConfigurationSection.svelte
- kind: doc
  role: intent
  target: docs/llm-classifier.mdx
---

# Change classification settings

An administrator sets how Suchi suggests document details: the application mode, the model confidence threshold, similar-document matching and its thresholds, and the optional document model. Saving settings never accepts suggestions already waiting.
