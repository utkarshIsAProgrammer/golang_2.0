# Go Engineering Journey --- From Basic to God Level

**Primary language:** Go\
**Learning method:** Build real software, learn concepts when the
project requires them, test, debug, optimize, and deploy.

**Long-term goal:** Become a strong backend and distributed-systems
engineer with deep understanding of databases, networking, operating
systems, programming languages, and infrastructure.

This is a progressive project roadmap. The projects are the
curriculum---not a requirement to rush through a checklist. Each project
should be completed to an appropriate engineering standard before moving
on.

---

## Level 1 --- Backend Engineering

**Goal:** Become capable of designing, building, testing, deploying, and
maintaining production-style backend applications.

### 01. Task Management Platform --- Current Project

Build a multi-user Trello/Linear-style application.

- Users, registration, login, and profiles
- Projects, tasks, assignments, statuses, comments, and activity
  history
- REST API using Go's `net/http`
- PostgreSQL, SQL, migrations, and transactions
- Authentication, authorization, and role-based access control
- Validation, error handling, testing, Docker, deployment, and
  observability

### 02. E-Commerce and Order Management System

Build a commerce backend with a complete order lifecycle.

- Product catalog, categories, carts, and inventory
- Orders, payment integration using a test provider, refunds, and
  webhooks
- Inventory consistency and concurrent purchases
- Database transactions and isolation levels
- Idempotency, background processing, and audit history

### 03. URL Shortener and Analytics Platform

Build a Bitly-like link management and redirect service.

- Short-code generation and redirect handling
- Link ownership, expiration, and custom aliases
- Click analytics and aggregation
- Redis caching and rate limiting
- Database indexing, benchmarking, and load testing

### 04. Real-Time Chat and Notification System

Build a Slack-like messaging backend.

- Direct messages, group conversations, and message history
- WebSockets and connection lifecycle management
- Goroutines, channels, mutexes, and concurrency safety
- Online presence, delivery acknowledgements, and retries
- Redis Pub/Sub and notification delivery

### 05. Background Job Processing Platform

Build a task queue with workers and scheduling.

- Job submission, queues, and worker pools
- Scheduling, retries, backoff, and dead-letter queues
- Job state, cancellation, and observability
- At-least-once delivery and idempotent job execution
- Persistence and graceful worker shutdown

### 06. API Gateway and Reverse Proxy

Build a gateway that routes traffic to backend services.

- Reverse proxying and route configuration
- Load balancing and health checks
- Rate limiting, authentication, and request policies
- Connection pooling, timeouts, and TLS termination
- Metrics, graceful shutdown, and failure handling

**Level 1 outcome:** You can independently build and operate complete
backend applications and understand the tradeoffs behind their
architecture.

---

## Level 2 --- Systems Engineering

**Goal:** Stop treating infrastructure as a black box. Build simplified
versions of the systems backend engineers depend on.

### 07. Build Your Own Redis-Like In-Memory Database

- TCP server and a custom request/response protocol
- Key-value data structures and commands such as `GET`, `SET`, and
  `DEL`
- Expiration and TTL management
- Concurrent clients and synchronization
- Persistence, snapshots, and eviction policies

### 08. Build Your Own Container Runtime

- Process creation and lifecycle management
- Linux namespaces and cgroups
- Filesystem isolation and mount namespaces
- Networking and resource limits
- Container configuration and image/layer concepts

### 09. Build Your Own Database Engine

- Page and record formats
- Disk storage and buffer management
- B-tree or LSM-tree indexes
- Write-ahead logging (WAL)
- Query execution, transactions, and crash recovery

### 10. Build Your Own Message Broker

- Producers, consumers, and topics
- Append-only logs and partitions
- Consumer groups and offsets
- Persistence and delivery semantics
- Retry handling and replication concepts

### 11. Build Your Own Programming Language

- Language syntax and grammar
- Lexer/tokenizer and parser
- Abstract syntax tree (AST)
- Interpreter and runtime
- Type checking, bytecode, and eventually compilation

### 12. Build a Mini Operating System

- Boot process and kernel entry
- C and assembly fundamentals
- Interrupts and system calls
- Memory management
- Scheduling, processes, and a basic filesystem

**Level 2 outcome:** You understand the internal mechanisms behind
databases, runtimes, brokers, languages, and operating systems.

---

## Level 3 --- Distributed Systems Engineering

**Goal:** Build systems that span multiple machines and continue working
through failures.

### 13. Distributed Key-Value Store

- Multiple nodes and cluster membership
- Leader/follower replication
- Raft consensus and leader election
- Log replication and snapshots
- Consistency, failure detection, and recovery

### 14. Service Discovery and Coordination System

Build a simplified Consul/ZooKeeper-style coordination service.

- Service registration and discovery
- Heartbeats and health monitoring
- Leases and distributed locks
- Membership and failure detection
- Coordination and recovery after node failures

### 15. Complete Microservices Platform

Build a multi-service business platform.

- Independently deployable Go services
- API Gateway and gRPC/Protobuf
- Event-driven communication and Kafka
- Saga workflows and distributed transaction challenges
- Authentication, authorization, tracing, and observability

### 16. Distributed Object Storage

Build a simplified S3-like storage system.

- Object upload, download, and metadata
- Chunking and storage across nodes
- Replication and checksums
- Data placement and consistency
- Node failure, repair, and recovery

### 17. Distributed Search Engine

Build a simplified Elasticsearch-like search platform.

- Document ingestion and tokenization
- Inverted indexes
- Query parsing and relevance ranking
- Indexing pipelines and storage
- Sharding, distributed queries, and result merging

**Level 3 outcome:** You can reason about distributed state,
consistency, replication, coordination, and failure recovery.

---

## Level 4 --- Infrastructure and Developer Platforms

**Goal:** Build platforms that manage, observe, and automate software
systems.

### 18. Build a Mini Kubernetes-Like Orchestrator

- Cluster nodes and workload scheduling
- Desired state and reconciliation loops
- Container lifecycle management
- Health checks and service networking
- Control-plane design and failure recovery

### 19. Build an Observability Platform

Build a simplified metrics, logging, and tracing platform.

- Metrics collection and time-series storage
- Structured log ingestion and search
- Distributed tracing and context propagation
- Dashboards, alert rules, and notifications
- Data retention and collection pipelines

### 20. Build Your Own CI/CD Platform

Build a simplified GitHub Actions/Jenkins-style runner.

- Pipeline definitions and job scheduling
- Build and test execution
- Isolated job execution
- Artifact storage and build logs
- Secrets handling and deployment automation

### 21. Build a Mini Cloud Platform

Build a control plane that provisions and manages application instances.

- Compute instance lifecycle
- Networking and storage provisioning
- Resource scheduling, quotas, and isolation
- Control-plane APIs and state management
- Monitoring, failure recovery, and automation

**Level 4 outcome:** You understand how developer platforms and cloud
infrastructure provision, operate, and monitor software.

---

## Level 5 --- God-Level Engineering Challenges

**Goal:** Combine the previous skills into ambitious, research-grade
engineering systems. These are long-term challenges, not quick portfolio
projects.

### 22. Globally Distributed Database

- Multi-region architecture
- Replication and consensus
- Consistency models and conflict handling
- Failover, partition tolerance, and recovery
- Latency, durability, and availability tradeoffs

### 23. High-Performance Language Runtime

- Virtual machine and bytecode execution
- Memory management and garbage collection
- JIT compilation
- Runtime profiling and optimization
- Concurrency and runtime scheduling

### 24. Distributed Computing Framework

Build a simplified MapReduce/Spark-like system.

- Distributed task execution
- Data partitioning and shuffling
- Worker scheduling and coordination
- Fault recovery and task re-execution
- Parallel computation and performance measurement

### 25. Cloud Operating Platform

Combine compute, networking, storage, orchestration, and coordination.

- Resource provisioning and scheduling
- Isolation and multi-tenancy
- Distributed control plane
- Storage and networking services
- Observability, security, and recovery

### 26. Large-Scale Internet Service

Design, build, and operate a complete service against demanding scale
and reliability targets.

- End-to-end architecture and capacity planning
- Multi-region deployment
- Caching, partitioning, and replication
- Load, stress, and soak testing
- Fault injection, disaster recovery, and security
- Operational runbooks and measurable service objectives

**Level 5 outcome:** You can take on complex systems problems, make
explicit engineering tradeoffs, and validate your designs through
implementation and testing.

---

## How We Will Execute Every Project

Each project follows an engineering lifecycle. The exact milestones vary
by project.

1.  **Requirements:** Define the problem, users, use cases, and
    constraints.
2.  **Design:** Choose architecture, data models, APIs, and protocols.
3.  **Build:** Implement the system incrementally; write the code
    yourself.
4.  **Understand:** Learn the underlying concepts when they become
    necessary.
5.  **Test:** Add unit, integration, concurrency, and failure tests as
    appropriate.
6.  **Debug:** Reproduce defects and understand their root causes.
7.  **Optimize:** Profile CPU, memory, I/O, latency, and throughput.
8.  **Scale:** Identify bottlenecks and introduce appropriate
    distributed design.
9.  **Ship:** Package, deploy, document, monitor, and maintain the
    software.
10. **Prove mastery:** Rebuild or extend important components
    independently and explain the tradeoffs.

### Engineering principles

- Tutorials are references, not solutions to copy.
- Go is the primary language; use C, assembly, SQL, or other tools
  when the problem calls for them.
- Introduce abstractions when they solve a real problem, not just to
  make the architecture look sophisticated.
- Do not skip fundamentals that are necessary to understand the
  system.
- A working demo is not automatically a finished project.
- Measure performance instead of guessing.
- Treat security, testing, reliability, and operations as part of
  engineering---not optional polish.

## Current Position

**Project 01: Task Management Platform**

Current milestone: application foundation.

- [x] Go project structure
- [x] HTTP server
- [x] Router
- [x] Health endpoint
- [ ] About endpoint
- [ ] Task CRUD API
- [ ] PostgreSQL integration
- [ ] Authentication and authorization
- [ ] Testing and production deployment

**Immediate next assignment:** Implement and register the `GET /about`
handler, test it alongside `GET /health`, and review the code before
moving to task functionality.

---

## The Standard

The objective is not to collect 26 repositories or finish projects as
quickly as possible.

The objective is to become capable of independently designing,
implementing, testing, debugging, scaling, and operating complex
software systems---and explaining why they work.
