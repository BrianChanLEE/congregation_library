# PRD: Backend/API/Frontend Stability Roadmap

## Problem Statement

The project currently works in narrow happy paths, but authentication, API contracts, schema changes, and frontend data fetching are not stable enough for repeatable development. Recent fixes restored local behavior, but several implementation details still violate `GEMINI.md` standards and create recurring runtime errors.

## Solution

Stabilize the system in five ordered steps:

1. Remove the `AuthService` logger panic and make tests reliable.
2. Align authentication behavior with the OpenAPI contract.
3. Replace ad hoc runtime schema edits with versioned, idempotent migrations.
4. Move frontend server-state fetching to TanStack Query.
5. Remove noisy educational comments from production code.

## User Stories

1. As a developer, I want backend tests to pass consistently, so that I can safely change authentication code.
2. As a user, I want expired or missing auth to route me clearly to login, so that protected screens do not fail with noisy console errors.
3. As a frontend developer, I want API docs to declare auth requirements, so that UI behavior matches backend constraints.
4. As an operator, I want schema changes to be versioned and idempotent, so that deploys are predictable and auditable.
5. As a developer, I want frontend data fetching to use React Query, so that loading, error, cache, and retry behavior are consistent.
6. As a maintainer, I want production code comments to explain intent rather than syntax, so that files stay readable.

## Implementation Decisions

- Initialize or guard the project logger so service-level tests cannot panic when logging paths execute.
- Document bearer authentication for protected endpoints, especially history, inventory, profile, transactions, and admin APIs.
- Keep migrations idempotent, but track applied migration IDs in a schema version table.
- Use React Query for movement history fetching before expanding the same pattern to other screens.
- Preserve current routing decisions and existing uncommitted work unless directly required for the stability sequence.

## Testing Decisions

- Backend validation uses `go test ./...` with `GOCACHE` inside the workspace when sandbox access blocks the default Go cache.
- Frontend validation uses `npm test --silent` and `npm run build`.
- Tests should assert behavior at service/hook/screen boundaries instead of implementation details.

## Out of Scope

- Full redesign of login/session refresh.
- Full migration framework adoption such as goose or golang-migrate.
- Complete conversion of every screen to React Query in one pass.
- Visual parity work against `stitch_ref`.

## Further Notes

- Browser `contentscript.js` warnings are extension-originated and are not treated as application defects.
- Runtime schema repair remains useful for local/legacy databases, but it must be versioned and visible.
