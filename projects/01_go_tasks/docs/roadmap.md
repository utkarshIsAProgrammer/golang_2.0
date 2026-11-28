# Project 01 --- Task Management Platform

## Complete Project Plan: Phases, Tasks, Todos, and Completion Criteria

**Project type:** Multi-user task and project management backend\
**Primary language:** Go\
**HTTP:** Go standard library (`net/http`)\
**Database:** PostgreSQL\
**Cache / asynchronous work:** Redis, introduced only when needed\
**Deployment:** Docker, CI/CD, and a Linux server or cloud platform

**Current repository/module:**
`github.com/utkarshIsAProgrammer/go_task_manager`

**Current known progress:** Go project structure, HTTP server, router,
and `/health` handler have been created. The next small assignment is
the `/about` endpoint.

---

# 1. Product Definition

Build a backend for a collaborative task management application inspired
by Trello and Linear.

Users should be able to register, sign in, create projects, invite or
add collaborators, create tasks, assign tasks, update their status,
comment, and inspect activity history.

This is a backend-first project. A frontend is optional and should not
distract from completing the API and its engineering requirements.

## Core user journeys

1.  A new user registers and signs in.
2.  The user creates a project.
3.  The project owner adds collaborators.
4.  A project member creates a task.
5.  The task is assigned to a project member.
6.  A member changes its status, priority, or due date.
7.  Members discuss the task through comments.
8.  The system records important actions in an activity history.
9.  Users can only access projects and tasks they are authorized to
    access.

## Initial feature scope

- User registration, login, logout, and profile
- Projects and project memberships
- Tasks, assignments, statuses, priorities, and due dates
- Comments
- Activity history
- Authorization and roles
- Search, filtering, sorting, and pagination
- Production concerns: testing, security, logging, metrics,
  deployment, and documentation

## Explicitly out of scope for the first release

Do not build these until the core product is reliable:

- A full frontend
- Microservices
- Kubernetes
- Event streaming with Kafka
- Complex real-time collaboration
- File uploads
- Billing and subscriptions
- Mobile applications

These can become later extensions. The first version should be a
well-engineered modular monolith.

---

# 2. Target Architecture

Start simple and add layers when the requirements justify them.

```text
go_task_manager/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── handler/
│   ├── router/
│   ├── middleware/
│   ├── model/
│   ├── service/
│   ├── repository/
│   └── database/
├── migrations/
├── docs/
├── scripts/
├── Dockerfile
├── compose.yaml
├── Makefile
├── .env.example
├── .gitignore
├── README.md
└── go.mod
```

Do not create every directory immediately. Introduce each package when
the project reaches the phase that needs it.

## Responsibility boundaries

- `cmd/api`: application entry point and dependency wiring.
- `internal/config`: environment-based configuration.
- `internal/router`: URL patterns and route registration.
- `internal/handler`: HTTP input/output, status codes, and request
  parsing.
- `internal/middleware`: request logging, recovery, authentication,
  and other cross-cutting HTTP behavior.
- `internal/model`: domain and data structures.
- `internal/service`: business rules and use cases.
- `internal/repository`: persistence operations.
- `internal/database`: database connection and setup.
- `migrations`: versioned database schema changes.

## Request flow

```text
HTTP client
    |
    v
Router
    |
    v
Middleware
    |
    v
Handler
    |
    v
Service
    |
    v
Repository
    |
    v
PostgreSQL
```

The response travels back through the same application layers.

---

# 3. Phase 0 --- Requirements, Scope, and Project Setup

**Goal:** Establish what we are building and create a clean,
reproducible development environment.

## Tasks and todos

- [ ] Write a concise product description.
- [ ] Define the first-release features and explicitly defer
      out-of-scope features.
- [ ] Identify user roles: owner, admin, member.
- [ ] Write the main user journeys.
- [ ] Define initial API conventions:
  - [ ] JSON request and response bodies
  - [ ] Consistent error response format
  - [ ] Resource IDs
  - [ ] HTTP status code conventions
  - [ ] Pagination conventions
- [ ] Confirm the Go module path matches `go.mod` and all imports.
- [ ] Confirm the project runs with `go run ./cmd/api`.
- [ ] Initialize and inspect Git history.
- [ ] Add `.gitignore`.
- [ ] Add `.env.example` without real secrets.
- [ ] Create a README with project purpose and local run instructions.
- [ ] Record the Go version and required external dependencies.

## Exit criteria

- The project builds from a clean checkout.
- The module path and imports are consistent.
- No secrets or local environment files are committed.
- The product scope is written down.

---

# 4. Phase 1 --- HTTP Server and Application Foundation

**Goal:** Build a reliable HTTP application foundation before
implementing business features.

## 1.1 Server and routing

- [x] Create `cmd/api/main.go`.
- [x] Create `internal/router/router.go`.
- [x] Configure `http.Server`.
- [x] Make the router implement/return `http.Handler`.
- [x] Register `GET /health`.
- [ ] Create `internal/handler/about_handler.go`.
- [ ] Implement `GET /about`.
- [ ] Register `/about` in the router.
- [ ] Test `/health` and `/about` with `curl`.
- [ ] Verify an unknown route returns 404.
- [ ] Verify unsupported HTTP methods are handled appropriately.

## 1.2 HTTP behavior

- [ ] Understand the HTTP request lifecycle.
- [ ] Understand methods: GET, POST, PUT, PATCH, DELETE.
- [ ] Understand path parameters and query parameters.
- [ ] Understand headers, status codes, and request bodies.
- [ ] Set appropriate `Content-Type` headers.
- [ ] Return JSON consistently.
- [ ] Avoid writing a response twice.
- [ ] Handle malformed request bodies.

## 1.3 Server reliability

- [ ] Move server configuration out of hardcoded scattered values.
- [ ] Add configurable listen address and port.
- [ ] Add read, write, and idle timeouts.
- [ ] Implement graceful shutdown on SIGINT/SIGTERM.
- [ ] Ensure in-flight requests have time to finish.
- [ ] Log startup and shutdown.
- [ ] Return a non-zero process exit on startup failure.

## 1.4 Foundation tests

- [ ] Add a router test.
- [ ] Add a health handler test using `httptest`.
- [ ] Add an about handler test.
- [ ] Test unknown routes.
- [ ] Test method handling.

## Exit criteria

- The server starts and shuts down cleanly.
- `/health` and `/about` work.
- HTTP behavior is tested.
- The application has no business logic in `main.go`.

**Immediate assignment:** Finish `GET /about`, register it, test both
endpoints, and submit the code for review.

---

# 5. Phase 2 --- Task API: First Real Product Feature

**Goal:** Build task creation and retrieval, then complete CRUD
behavior.

Start with a clearly defined task domain. Do not add database
persistence before the API and domain behavior are understood. Use a
small in-memory implementation temporarily if needed, then replace it
with PostgreSQL in Phase 4.

## 2.1 Define the Task domain

- [ ] Define a `Task` model.
- [ ] Decide the fields:
  - [ ] ID
  - [ ] Title
  - [ ] Description
  - [ ] Status
  - [ ] Priority
  - [ ] CreatedAt
  - [ ] UpdatedAt
  - [ ] DueDate (optional)
- [ ] Define allowed task statuses: `todo`, `in_progress`,
      `completed`.
- [ ] Define allowed priorities: `low`, `medium`, `high`, `urgent`.
- [ ] Decide which fields are required.
- [ ] Separate API request structures from internal/domain structures
      where appropriate.
- [ ] Do not let clients set server-controlled fields such as IDs and
      timestamps.

## 2.2 Create tasks

- [ ] Implement `POST /tasks`.
- [ ] Decode JSON request bodies.
- [ ] Validate title and other input.
- [ ] Reject malformed JSON.
- [ ] Reject missing or invalid required fields.
- [ ] Generate the task ID.
- [ ] Set default status and priority.
- [ ] Set creation and update timestamps.
- [ ] Return `201 Created`.
- [ ] Return the created task in JSON.
- [ ] Consider the `Location` response header.

## 2.3 Read tasks

- [ ] Implement `GET /tasks`.
- [ ] Return all tasks for the initial single-user prototype.
- [ ] Implement `GET /tasks/{id}`.
- [ ] Parse and validate the task ID.
- [ ] Return `404 Not Found` for a missing task.
- [ ] Return a consistent JSON response.

## 2.4 Update tasks

- [ ] Implement `PATCH /tasks/{id}` for partial updates.
- [ ] Decide how omitted fields differ from explicitly supplied zero
      values.
- [ ] Validate updated values.
- [ ] Update `UpdatedAt`.
- [ ] Return the updated task.
- [ ] Return 404 for an unknown task.

## 2.5 Delete tasks

- [ ] Implement `DELETE /tasks/{id}`.
- [ ] Decide whether deletion is permanent or soft deletion.
- [ ] Return an appropriate success status.
- [ ] Return 404 for an unknown task.
- [ ] Ensure deleted tasks are no longer returned by normal reads.

## 2.6 Temporary storage and concurrency

If using in-memory storage before PostgreSQL:

- [ ] Encapsulate storage rather than using a package-global map.
- [ ] Protect shared mutable state with a mutex.
- [ ] Use `sync.RWMutex` where read/write access patterns justify it.
- [ ] Avoid returning mutable internal slices or maps directly.
- [ ] Run `go test -race ./...`.
- [ ] Test concurrent task creation and reads.

## 2.7 Task API tests

- [ ] Test successful creation.
- [ ] Test malformed JSON.
- [ ] Test validation failures.
- [ ] Test listing.
- [ ] Test retrieval by ID.
- [ ] Test unknown ID.
- [ ] Test updates and partial updates.
- [ ] Test deletion.
- [ ] Test concurrent access if using in-memory storage.

## Exit criteria

- All five task endpoints work.
- Invalid input produces predictable errors.
- Tests cover success and failure cases.
- Shared state is race-safe if stored in memory.

---

# 6. Phase 3 --- Application Architecture and Error Handling

**Goal:** Separate HTTP concerns, business rules, and data access
without overengineering.

## 3.1 Introduce service and repository boundaries

- [ ] Identify logic currently inside handlers.
- [ ] Move business rules into a task service.
- [ ] Introduce a task repository abstraction where it provides value.
- [ ] Keep HTTP request parsing and response writing in handlers.
- [ ] Keep persistence operations in repository implementations.
- [ ] Wire dependencies from `main.go`.
- [ ] Avoid global mutable dependencies.
- [ ] Use constructor functions where useful.
- [ ] Keep interfaces small and owned by the consumer where practical.

## 3.2 Define business rules

- [ ] Define valid status transitions.
- [ ] Decide whether completed tasks can be reopened.
- [ ] Decide whether title changes are allowed after completion.
- [ ] Define due-date rules.
- [ ] Define task deletion behavior.
- [ ] Ensure business rules are enforced in the service layer.
- [ ] Test business rules independently of HTTP.

## 3.3 Standardize errors

- [ ] Define application-level error types or sentinel errors where
      useful.
- [ ] Distinguish validation, not-found, conflict, unauthorized,
      forbidden, and internal errors.
- [ ] Map application errors to HTTP status codes in the handler
      layer.
- [ ] Define a consistent error JSON structure.
- [ ] Avoid exposing SQL errors, stack traces, or secrets to clients.
- [ ] Log internal details safely.
- [ ] Preserve useful error context with wrapping where appropriate.

## 3.4 Configuration and dependency wiring

- [ ] Add `internal/config`.
- [ ] Read configuration from environment variables.
- [ ] Validate required configuration at startup.
- [ ] Set sensible development defaults.
- [ ] Separate development and production settings.
- [ ] Keep secrets out of source control.
- [ ] Wire router, services, repositories, and database dependencies
      in the composition root.

## Exit criteria

- Handlers are thin.
- Business rules can be tested without an HTTP server.
- Persistence can be replaced without rewriting business logic.
- Errors are predictable and do not leak internals.

---

# 7. Phase 4 --- PostgreSQL and Persistent Storage

**Goal:** Replace temporary storage with a properly designed relational
database.

## 4.1 Database environment

- [ ] Start and verify the local PostgreSQL service.
- [ ] Create a development database and database user.
- [ ] Configure the connection string through environment variables.
- [ ] Add a PostgreSQL driver (for example, `pgx`).
- [ ] Create a database connection/pool during startup.
- [ ] Configure connection limits and timeouts.
- [ ] Verify connectivity at startup.
- [ ] Close the pool cleanly during shutdown.

## 4.2 Schema design

Create versioned migrations.

- [ ] Create the initial `users` table.
- [ ] Create the `projects` table.
- [ ] Create the `project_members` table.
- [ ] Create the `tasks` table.
- [ ] Define primary keys.
- [ ] Define foreign keys.
- [ ] Define uniqueness constraints.
- [ ] Define `NOT NULL` constraints.
- [ ] Define sensible defaults.
- [ ] Add created/updated timestamps.
- [ ] Decide deletion behavior and cascading rules.
- [ ] Add indexes based on actual query patterns.
- [ ] Write both up and down migrations.
- [ ] Test migrations against a fresh database.

Suggested initial relationships:

```text
users
  ├── project_members ── projects
  ├── tasks (created_by)
  └── tasks (assignee_id)

projects
  ├── project_members
  └── tasks

tasks
  ├── comments
  └── activity_events
```

## 4.3 Implement task persistence

- [ ] Implement task repository methods using SQL.
- [ ] Create a task in PostgreSQL.
- [ ] Fetch one task by ID.
- [ ] List tasks.
- [ ] Update a task.
- [ ] Delete or soft-delete a task.
- [ ] Handle `sql.ErrNoRows` or the driver's equivalent.
- [ ] Use parameterized queries.
- [ ] Never concatenate untrusted input into SQL.
- [ ] Add repository integration tests.
- [ ] Replace temporary in-memory storage with PostgreSQL.

## 4.4 SQL competence through the project

- [ ] Practice SELECT, INSERT, UPDATE, and DELETE.
- [ ] Use JOINs for related data.
- [ ] Understand primary and foreign keys.
- [ ] Understand unique and check constraints.
- [ ] Use transactions where multiple writes must succeed together.
- [ ] Inspect query plans with `EXPLAIN`.
- [ ] Add indexes only when justified by query patterns.
- [ ] Understand connection pooling.
- [ ] Understand transaction isolation and locking.

## 4.5 Database reliability

- [ ] Add database health checks.
- [ ] Set query and transaction timeouts.
- [ ] Handle database unavailability gracefully.
- [ ] Test rollback behavior.
- [ ] Test concurrent updates.
- [ ] Document backup and restore procedures.

## Exit criteria

- Data survives application restarts.
- Schema changes are managed through migrations.
- SQL is parameterized.
- Repository integration tests pass.
- Database failures do not crash the application unpredictably.

---

# 8. Phase 5 --- Users, Authentication, and Security

**Goal:** Make the application multi-user and protect private resources.

## 5.1 User registration

- [ ] Define the user model.
- [ ] Add email and password input validation.
- [ ] Normalize email addresses consistently.
- [ ] Enforce unique email addresses in PostgreSQL.
- [ ] Hash passwords with a suitable password-hashing algorithm (for
      example, bcrypt).
- [ ] Never store plaintext passwords.
- [ ] Return safe user fields only.
- [ ] Handle duplicate registration.

## 5.2 Login and tokens

- [ ] Implement `POST /auth/login`.
- [ ] Verify password hashes.
- [ ] Issue short-lived access tokens.
- [ ] Design refresh-token behavior.
- [ ] Store refresh tokens safely (prefer storing a hash).
- [ ] Implement token refresh.
- [ ] Implement logout/revocation.
- [ ] Validate token signature, expiry, issuer, and relevant claims.
- [ ] Keep signing keys in environment/secret management, never in
      source control.

## 5.3 Authentication middleware

- [ ] Implement authentication middleware.
- [ ] Read bearer tokens from the Authorization header.
- [ ] Validate the token.
- [ ] Put authenticated user identity into request context.
- [ ] Reject missing or invalid credentials with 401.
- [ ] Ensure public routes and protected routes are explicitly
      separated.
- [ ] Test expired, malformed, and invalid-signature tokens.

## 5.4 Authorization

- [ ] Ensure users can only access their own permitted resources.
- [ ] Enforce authorization in application use cases, not only by
      hiding UI actions.
- [ ] Distinguish unauthenticated (401) from unauthorized (403).
- [ ] Prevent insecure direct object reference (IDOR) vulnerabilities.
- [ ] Test cross-user access attempts.
- [ ] Avoid trusting user IDs supplied by clients when identity is
      available from the token.

## 5.5 Security hardening

- [ ] Set request body size limits.
- [ ] Validate all untrusted input.
- [ ] Configure CORS only for known frontend origins if browser access
      is needed.
- [ ] Add rate limiting to sensitive endpoints.
- [ ] Avoid leaking whether a password or account exists where
      inappropriate.
- [ ] Use HTTPS in deployed environments.
- [ ] Protect secrets and rotate keys when necessary.
- [ ] Review dependency vulnerabilities.
- [ ] Add security-focused tests.

## Exit criteria

- Users can register, log in, refresh, and log out.
- Passwords are hashed.
- Protected endpoints require valid authentication.
- Users cannot access another user's private data without permission.
- Security-sensitive behavior has tests.

---

# 9. Phase 6 --- Projects, Memberships, Roles, and Assignments

**Goal:** Turn the task API into a collaborative project management
product.

## 6.1 Projects

- [ ] Define the project model.
- [ ] Implement `POST /projects`.
- [ ] Implement `GET /projects`.
- [ ] Implement `GET /projects/{id}`.
- [ ] Implement `PATCH /projects/{id}`.
- [ ] Implement `DELETE /projects/{id}`.
- [ ] Ensure project lists are scoped to the authenticated user.
- [ ] Add pagination to project listing.

## 6.2 Project memberships

- [ ] Define project membership records.
- [ ] Add owner/admin/member roles.
- [ ] Implement adding a member.
- [ ] Implement removing a member.
- [ ] Implement listing project members.
- [ ] Prevent duplicate memberships.
- [ ] Prevent unauthorized membership changes.
- [ ] Define what happens when a project owner leaves.
- [ ] Test membership and permission edge cases.

## 6.3 Task/project relationships

- [ ] Associate tasks with projects.
- [ ] Ensure tasks belong to a valid project.
- [ ] Ensure only project members can view project tasks.
- [ ] Ensure only authorized members can create or update tasks.
- [ ] Assign tasks to project members.
- [ ] Prevent assigning a task to a non-member unless product rules
      explicitly allow it.
- [ ] Filter tasks by project, assignee, status, and priority.

## 6.4 Roles and permissions

- [ ] Define permissions for owner, admin, and member.
- [ ] Define who can edit project settings.
- [ ] Define who can manage members.
- [ ] Define who can delete tasks.
- [ ] Define who can change task assignments.
- [ ] Centralize permission checks in appropriate application logic.
- [ ] Add tests for allowed and denied operations.

## Exit criteria

- Multiple users can collaborate on projects.
- Project access is enforced server-side.
- Roles and permissions are consistent.
- Task assignment respects project membership.

---

# 10. Phase 7 --- Comments, Activity History, and Product Completeness

**Goal:** Add collaboration and traceability.

## 7.1 Comments

- [ ] Define the comment model.
- [ ] Add a comments table and migration.
- [ ] Implement creating a comment.
- [ ] Implement listing task comments.
- [ ] Implement editing a comment with ownership rules.
- [ ] Implement deleting a comment with permission rules.
- [ ] Add pagination for comment history.
- [ ] Validate comment length and content.
- [ ] Test access control.

## 7.2 Activity history

- [ ] Define activity event types.
- [ ] Record task creation.
- [ ] Record task updates.
- [ ] Record status changes.
- [ ] Record assignment changes.
- [ ] Record comments and relevant project membership changes.
- [ ] Store actor, target, event type, timestamp, and useful change
      metadata.
- [ ] Make activity records append-only from normal application paths.
- [ ] Implement listing project/task activity.
- [ ] Test that failed operations do not create misleading activity
      records.

## 7.3 Product usability features

- [ ] Add task search.
- [ ] Add filtering by status, priority, assignee, and due date.
- [ ] Add sorting.
- [ ] Add pagination with stable ordering.
- [ ] Add overdue-task queries.
- [ ] Add project summaries and task counts.
- [ ] Define empty-state and not-found behavior in the API.
- [ ] Document API examples.

## Exit criteria

- Users can collaborate through comments.
- Important changes are auditable.
- Search, filtering, sorting, and pagination work predictably.

---

# 11. Phase 8 --- Testing and Concurrency Engineering

**Goal:** Build confidence that the application behaves correctly under
normal and concurrent use.

## 8.1 Unit tests

- [ ] Test validation functions.
- [ ] Test task business rules.
- [ ] Test project permission rules.
- [ ] Test token validation logic.
- [ ] Test service behavior using controlled dependencies.
- [ ] Test error mapping.

## 8.2 HTTP tests

- [ ] Test every endpoint's success path.
- [ ] Test malformed JSON.
- [ ] Test missing fields.
- [ ] Test invalid IDs.
- [ ] Test authentication failures.
- [ ] Test authorization failures.
- [ ] Test response status codes and JSON shapes.
- [ ] Test pagination and filters.

## 8.3 Database integration tests

- [ ] Run tests against a dedicated test database.
- [ ] Apply migrations before tests.
- [ ] Clean up test data safely.
- [ ] Test constraints and foreign keys.
- [ ] Test transaction rollback.
- [ ] Test concurrent updates and conflicting operations.

## 8.4 Concurrency and race detection

- [ ] Identify shared mutable state.
- [ ] Avoid unnecessary shared state.
- [ ] Use synchronization only where required.
- [ ] Test concurrent requests.
- [ ] Run `go test -race ./...`.
- [ ] Test graceful shutdown with active requests.
- [ ] Review goroutine lifetimes and cancellation.
- [ ] Use contexts and deadlines for request-scoped work.

## 8.5 Test quality

- [ ] Add table-driven tests where useful.
- [ ] Avoid tests that depend on execution order.
- [ ] Avoid relying on arbitrary sleeps.
- [ ] Make failure output useful.
- [ ] Add coverage measurement, but do not optimize for coverage
      percentage alone.
- [ ] Add repeatable test commands to the README or Makefile.

## Exit criteria

- Critical business rules and authorization paths are tested.
- Database integration tests are repeatable.
- Race detector passes.
- Tests do not rely on fragile timing or shared external state.

---

# 12. Phase 9 --- Performance and Scalability

**Goal:** Measure the system, identify bottlenecks, and improve them
with evidence.

## 9.1 Establish a baseline

- [ ] Create representative test data.
- [ ] Benchmark key service operations.
- [ ] Load-test important API endpoints.
- [ ] Measure throughput, latency percentiles, and error rate.
- [ ] Measure CPU and memory usage.
- [ ] Record baseline results.

## 9.2 Database performance

- [ ] Inspect slow queries.
- [ ] Use `EXPLAIN ANALYZE` on representative queries.
- [ ] Add or adjust indexes based on evidence.
- [ ] Avoid N+1 query patterns.
- [ ] Review pagination strategy.
- [ ] Review connection-pool configuration.
- [ ] Measure transaction contention.

## 9.3 Caching and background work

Only introduce Redis when a measured or clear product requirement
justifies it.

- [ ] Identify cacheable reads.
- [ ] Define cache keys and expiration.
- [ ] Define invalidation behavior.
- [ ] Prevent stale or cross-user cached data.
- [ ] Decide whether background work is necessary.
- [ ] Move non-critical work out of request paths when justified.
- [ ] Make background jobs retry-safe and idempotent.
- [ ] Test cache outages and job failures.

## 9.4 Performance verification

- [ ] Repeat the same load test after changes.
- [ ] Compare latency and throughput to baseline.
- [ ] Confirm correctness has not regressed.
- [ ] Document tradeoffs and remaining bottlenecks.

## Exit criteria

- Performance claims are supported by measurements.
- Bottlenecks have been identified and prioritized.
- Caching or asynchronous processing has a clear justification.

---

# 13. Phase 10 --- Production Readiness and Operations

**Goal:** Package, deploy, observe, and maintain the application.

## 10.1 Logging and observability

- [ ] Add structured logging.
- [ ] Include request IDs or correlation IDs.
- [ ] Log request method, route, status, and duration.
- [ ] Avoid logging passwords, tokens, or sensitive payloads.
- [ ] Expose application metrics.
- [ ] Measure request count, latency, and error rate.
- [ ] Measure database pool health.
- [ ] Add health and readiness checks with distinct meanings.
- [ ] Add tracing when the system's complexity justifies it.

## 10.2 Docker and local environment

- [ ] Create a multi-stage Dockerfile.
- [ ] Run the application as a non-root user.
- [ ] Add a `compose.yaml` for the app and PostgreSQL.
- [ ] Add persistent database storage for local development.
- [ ] Add health checks.
- [ ] Ensure configuration is supplied through environment variables.
- [ ] Ensure secrets are not baked into images.
- [ ] Test a clean build and startup.

## 10.3 CI pipeline

- [ ] Add automated formatting checks.
- [ ] Run `go vet ./...`.
- [ ] Run unit tests.
- [ ] Run integration tests.
- [ ] Run race tests where practical.
- [ ] Build the application.
- [ ] Build the container image.
- [ ] Fail the pipeline when checks fail.

## 10.4 Deployment

- [ ] Choose a deployment target.
- [ ] Configure production environment variables and secrets.
- [ ] Provision PostgreSQL securely.
- [ ] Configure HTTPS.
- [ ] Apply migrations safely during release.
- [ ] Deploy the application.
- [ ] Verify health and readiness.
- [ ] Test core user journeys against the deployed service.
- [ ] Document rollback steps.

## 10.5 Operations and recovery

- [ ] Define backup frequency and retention.
- [ ] Perform a database restore drill.
- [ ] Document incident diagnosis steps.
- [ ] Document common operational commands.
- [ ] Define basic service objectives.
- [ ] Review logs and metrics after deployment.
- [ ] Test application behavior when PostgreSQL is unavailable.
- [ ] Test graceful restart and deployment.

## Exit criteria

- The application is deployed and reachable over HTTPS.
- CI validates changes automatically.
- Logs and metrics support diagnosis.
- Backups and restoration have been tested.
- A documented rollback path exists.

---

# 14. Phase 11 --- Final Engineering Review

**Goal:** Prove that the project is understood and maintainable---not
merely functional.

## Final review checklist

### Product behavior

- [ ] Registration and login work.
- [ ] Users can create and manage projects.
- [ ] Project memberships and roles work.
- [ ] Users can create, assign, update, filter, and delete tasks.
- [ ] Comments and activity history work.
- [ ] Access control is enforced for every protected resource.

### Engineering quality

- [ ] Code is formatted.
- [ ] Packages have clear responsibilities.
- [ ] Business rules are not buried in HTTP handlers.
- [ ] SQL is parameterized.
- [ ] Migrations are versioned.
- [ ] Errors are handled consistently.
- [ ] Tests cover critical behavior.
- [ ] Race detection has been run.
- [ ] No secrets are committed.
- [ ] Dependencies are reviewed.

### Operations

- [ ] Docker build works.
- [ ] CI passes.
- [ ] Production deployment works.
- [ ] HTTPS is configured.
- [ ] Logs and metrics are available.
- [ ] Database backup and restore are documented and tested.
- [ ] Rollback procedure is documented.

### Independent mastery challenge

Without following a tutorial or copying an existing implementation:

- [ ] Add a new feature from requirements through database migration,
      service logic, handler, and tests.
- [ ] Diagnose a deliberately introduced bug.
- [ ] Explain the request lifecycle through the codebase.
- [ ] Explain why each architectural layer exists.
- [ ] Explain the main database schema and indexes.
- [ ] Explain the authentication and authorization model.
- [ ] Demonstrate how the application behaves during database failure.
- [ ] Identify the current performance bottleneck using measurements.

## Project completion definition

Project 01 is complete when the core product works, critical behavior is
tested, access control is reliable, data is persistent, the application
is deployed, and you can independently explain and extend the system.

A large checklist is not a reason to implement every advanced feature
immediately. Work phase by phase. Some tasks may be adjusted as we
discover better designs during implementation.

---

# 15. Suggested API Resource Map

This is the intended API surface. It is a target contract to refine
during implementation, not a requirement to implement every endpoint at
once.

## Health and metadata

Method Endpoint Purpose

---

GET `/health` Liveness check
GET `/ready` Readiness check
GET `/about` API metadata

## Authentication and users

Method Endpoint Purpose

---

POST `/auth/register` Register
POST `/auth/login` Login
POST `/auth/refresh` Refresh access token
POST `/auth/logout` Revoke refresh session
GET `/users/me` Current user profile
PATCH `/users/me` Update current user profile

## Projects and memberships

Method Endpoint Purpose

---

POST `/projects` Create project
GET `/projects` List accessible projects
GET `/projects/{projectID}` Get project
PATCH `/projects/{projectID}` Update project
DELETE `/projects/{projectID}` Delete project
GET `/projects/{projectID}/members` List members
POST `/projects/{projectID}/members` Add member
DELETE `/projects/{projectID}/members/{userID}` Remove member

## Tasks

Method Endpoint Purpose

---

POST `/projects/{projectID}/tasks` Create task
GET `/projects/{projectID}/tasks` List project tasks
GET `/tasks/{taskID}` Get task
PATCH `/tasks/{taskID}` Update task
DELETE `/tasks/{taskID}` Delete task
PATCH `/tasks/{taskID}/assignee` Change assignee

## Comments and activity

Method Endpoint Purpose

---

POST `/tasks/{taskID}/comments` Add comment
GET `/tasks/{taskID}/comments` List comments
PATCH `/comments/{commentID}` Edit comment
DELETE `/comments/{commentID}` Delete comment
GET `/tasks/{taskID}/activity` Task activity
GET `/projects/{projectID}/activity` Project activity

---

# 16. Working Agreement for This Project

- We work on one phase and a small number of tasks at a time.
- You write the implementation; I explain concepts, review code,
  identify problems, and give the next assignment.
- I will not dump a complete finished application for you to copy.
- When a concept is needed, we learn enough to use it correctly and
  then apply it.
- We will revisit earlier code when later requirements expose a design
  limitation.
- We will not add Redis, microservices, or other infrastructure merely
  because they appear in the roadmap.
- We will keep a clear record of completed tasks, current work,
  blockers, and the next assignment.

## Immediate next step

Finish **Phase 1, task 1.1: `GET /about`**.

1.  Create `internal/handler/about_handler.go`.
2.  Return API metadata as JSON.
3.  Register the route in `internal/router/router.go`.
4.  Run the server.
5.  Test both `/health` and `/about`.
6.  Share your implementation for review.

Then continue with the remaining Phase 1 tasks before starting the Task
API.
