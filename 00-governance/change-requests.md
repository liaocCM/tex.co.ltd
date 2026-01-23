# Change Requests

## Open

## Closed
## CR: Separate TODO records by client uid
Requested by: user
Affected file: 01-requirements/user-stories.yaml
Proposed change: Add a user story to scope all CRUD operations by a uid stored in localStorage, with data persisted in the shared DB per uid.
Reason: Multiple users now share the same DB and need logical separation of records.
Risk: Requires schema/API changes and migration; may affect existing data visibility.
Status: Approved
CTO Notes: Approved to protect user data separation in a shared DB. Coordinate schema and API updates via architecture artifacts.
