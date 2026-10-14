---
kind: primary
routes:
  web: Web
steps:
- text: The Member asks a question about the current documents or the whole archive
  kind: actor
  actor: member
  entities:
  - {entity: document, effect: reads, facts: []}
  contexts:
    web:
      place: suchi-web::archive::ask-your-documents
- text: The Product finds up to six readable Documents, sends their selected text and the question to the document model and shows the answer with links to them
  kind: product
  actor: member
  entities:
  - {entity: document, effect: reads, facts: [Title, Sensitivity, Extracted text]}
  contexts:
    web:
      place: suchi-web::archive::ask-your-documents
---

# Ask a question

## Trigger

The person opens Ask your documents and asks a question.

## Outcome

The person sees an answer with links to the documents it used and may save those documents as a view.

## Edge cases

- Confidential or Restricted text is sent only after the person consents for that conversation.
- Changing the document scope or the model clears the conversation.
