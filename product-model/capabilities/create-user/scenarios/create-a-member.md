---
kind: primary
routes:
  web: Web
steps:
- text: The Administrator enters the email, display name and password, the role and any additional access
  kind: actor
  actor: administrator
  entities: []
  contexts:
    web:
      place: suchi-web::archive-configuration::users
- text: The Product creates the Member with an active Account
  kind: product
  actor: administrator
  entities:
  - {entity: member, effect: creates, facts: [Ask your documents, Manage document dates, Manage mailboxes, Create share links, Share saved views]}
  - {entity: account, effect: creates, to: Active, facts: [Email, Display name, Password, Sign-in method], with: member}
  contexts:
    web:
      place: suchi-web::archive-configuration::users
---

# Create a member

## Trigger

An administrator fills in the new user form under People > Users.

## Outcome

The person can sign in; in a server with filing systems they still need membership before they can enter one.

## Edge cases

- An administrator created this way has every additional access implicitly.
