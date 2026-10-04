# greeter — PRD

## Problem Statement

Teams building and onboarding onto this platform need a minimal, known-good  
service to validate that the platform's conventions (build, deploy, contracts)  
work end to end. Without a small reference service, verifying those  
conventions means standing up something larger and harder to reason about. E2E marker s0p3-1004a.

## Solution

Greeter is a small Go HTTP service with a single endpoint that returns a JSON
greeting for a given name. It exists as a minimal, convention-compliant
reference implementation, following the patterns in the organization's
`app-factory-kaj/e2e-reference`.

## Actors

- **Client Application** — any caller (service, script, or tool) that sends
HTTP requests to the greeter service to retrieve a greeting.

## User Stories

1. As a client application, I want to send a GET request to `/hello` with a
 `name` parameter, so that I receive a JSON greeting personalized with that
 name.
2. As a client application, I want to receive a sensible default greeting
 when I omit the `name` parameter, so that the endpoint still returns a
 valid response rather than an error.

## Product Decisions

- Missing or empty `name`: the service returns a default greeting (e.g.
"Hello, World!") rather than an error. *assumed*
- Access: the `/hello` endpoint is public and requires no authentication,
consistent with its purpose as a minimal conventions-reference service.
*assumed*
- Response format: a JSON body carrying the greeting message (e.g.
`{"message": "Hello, <name>!"}`). *assumed*
- Implementation follows the conventions demonstrated in
`app-factory-kaj/e2e-reference`.

## Out of Scope

- Persistence or storage of any kind — the service is stateless.
- A user interface — this is an API-only service.
- Rate limiting, quotas, or API keys.
- Any endpoint beyond `/hello`.

## Open Questions

None at this time.