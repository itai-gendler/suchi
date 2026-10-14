---
entities:
- entity: filing-system
  shows:
  - Code
  - Name
  - Members
  collects:
  - Name
  - Members
references:
- kind: code
  role: implementation
  target: ui/src/lib/SystemSettings.svelte
---

# Filing system

The current filing system's code, name and the people admitted to it.
