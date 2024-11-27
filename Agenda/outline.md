# Design Patters for Resilient Distributed Systems

## Objectives

1. To understand the different patterns and tradeoffs while system designing
2. To understand the different patterns and tradeoffs for communication between services including achieving consensus
3. Overview of key implementation details of Kafka/MongoDB/Hadoop Distributed File System (HDFS)/Google File System (GFS)

## Agenda

### Day 1

- Application Architecture Overviews
  - Monoliths vs Modular vs Microservices vs Legacy
    - Database per service vs Shared Database
  - Client/Server
    - SOA
    - ReST
    - gRPC
    - Messaging
  - Distributed vs Cloud vs Serverless
    - Tradeoffs
      - Scalability
      - Resilience
      - Stability
      - Ability to Grow
      - Service Integration
      - Costs
  - CAP vs BASE vs PACELC

- Application Design for Scalability
  - 12 Factor App Principles
    - For Host / VM / Container
  - Strangler for migrating from a legacy application
  - Chassis for easier service creation
  - Observability
    - Log & Metrics Aggregation
    - Distributed Tracing
    - Heartbeat

- Patterns for Concurrent Systems
  - Active Object
  - Monitor Object
  - Lock
  - Thread-Specific Storage
  - Scheduler
  - Thread Pool
  - Async Completion Token
  - Lamport Clock
  - Generation Clock
  - Hybrid Clock
  - Reactor
  - Proctor

## Day 2

- Data Management Patterns within a single service
  - Event Sourcing
  - Domain Event
  - Two-Phase Commit
  - Saga
  - Write-Ahead Log

- Data Management Patterns across services
  - API Composition
  - CQRS
  - Transactional Outbox
  - Orchestration vs Choreography
  - Service Registry
  - Service Discovery
  - Self Registration
  - Circuit Breaker
  - Adapter vs Sidecar vs Ambassador

- Messaging Patterns across services
  - Point-to-Point
  - Publish/Subscribe
  - Message Bus
  - Command
  - Document
  - Event Notification
  - Return Address
  - Correlation
  - Message Sequence
  - Channel Adapter
  - Dead Letter Channel
  - Filters
  - Routers
  - Translators
  - Pipes
  - Endpoints

## Day 3

- Replication Patterns for Scalable Distributed Systems
  - Replicated Load-Balanced Services
  - Sharded Services
  - Master-Worker
  - Consistent Core
  - Master Election
  - Leaders and Followers
  - Scatter and Gather
  - Follower Reads

- Consistency Patterns for Sharded Systems
  - Idempotent Consumer
  - Write-Ahead Log
  - Replicated Log
  - Follower Reads

- Consensus Protocols
  - Paxos
  - Raft
  - Comparison of Consensus Protocols
  - Linearizability vs. Serializability

- Code & feature walkthrough of OSS projects
  - MongoDB
  - Apache Kafka
  - NATS Server
  - Hadoop Distributed File System
