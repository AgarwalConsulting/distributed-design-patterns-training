# Kafka Code Links

The Kafka codebase is vast and has many interdependent files, each responsible for a specific part of Kafka’s functionality. Here are some major files and approximate line numbers for interesting parts of the code base:

## 1. Replication Protocol

- **Leader-Follower Replication Logic**:
  - **File**: `kafka/server/ReplicaManager.scala`
  - **Lines**: The leader-follower replication logic begins around **line 800**, particularly in methods like `makeFollower`, `becomeLeader`, and `maybeShrinkIsr`.
- **In-Sync Replica (ISR) Management**:
  - **File**: `kafka/server/ReplicaManager.scala`
  - **Lines**: ISR updates are primarily managed in **lines 300–400**, with specific ISR changes in methods like `maybeExpandIsr` and `maybeShrinkIsr`.

## 2. Log Management

- **Segmented Log Files**:
  - **File**: `kafka/log/LogSegment.scala`
  - **Lines**: The code handling segmented logs, including reading and writing logic, starts around **line 100**.
- **Compaction and Retention Policies**:
  - **File**: `kafka/log/Log.scala`
  - **Lines**: The log compaction and retention logic is around **lines 1000–1200**. Methods like `maybeClean` and `deleteOldSegments` are crucial for this functionality.

## 3. ZooKeeper Dependency and Controller

- **Kafka Controller Logic**:
  - **File**: `kafka/controller/KafkaController.scala`
  - **Lines**: Kafka’s controller logic, including leader election and metadata propagation, is mainly around **lines 200–500**.
- **ZooKeeper Interactions**:
  - **File**: `kafka/zk/KafkaZkClient.scala`
  - **Lines**: Zookeeper client interactions like handling metadata and ISR updates are found around **lines 300–800**.

## 4. Kafka Streams API

- **Stream Processing Core Logic**:
  - **File**: `streams/src/main/java/org/apache/kafka/streams/processor/internals/StreamThread.java`
  - **Lines**: The main processing loop and thread management begin around **line 400**.
- **Stateful Transformations**:
  - **File**: `streams/src/main/java/org/apache/kafka/streams/kstream/internals/KStreamImpl.java`
  - **Lines**: Stateful operations like joins and aggregations are defined in **lines 100–300**.

## 5. Consumer and Producer API

- **Consumer Offset Management**:
  - **File**: `clients/src/main/java/org/apache/kafka/clients/consumer/KafkaConsumer.java`
  - **Lines**: Offset management functions are in **lines 1200–1300**, focusing on methods like `poll` and `commitSync`.
- **Producer Batching and Compression**:
  - **File**: `clients/src/main/java/org/apache/kafka/clients/producer/KafkaProducer.java`
  - **Lines**: Batching and compression settings are around **lines 500–800**, with methods like `send` and `accumulator`.

## 6. KRaft (Kafka Raft) Mode

- **Raft Implementation**:
  - **File**: `raft/src/main/java/org/apache/kafka/raft/KafkaRaftClient.java`
  - **Lines**: The core Raft consensus logic, including leader election, starts around **line 200**.
- **Metadata Management for KRaft**:
  - **File**: `metadata/src/main/java/org/apache/kafka/controller/QuorumController.java`
  - **Lines**: Metadata management and replication in KRaft mode are around **lines 300–700**.

## 7. Partitioner and Routing

- **Default and Sticky Partitioner**:
  - **File**: `clients/src/main/java/org/apache/kafka/clients/producer/internals/DefaultPartitioner.java`
  - **Lines**: Partitioning strategy, including sticky partitioning logic, is around **line 100**.
- **Custom Partitioner Interface**:
  - **File**: `clients/src/main/java/org/apache/kafka/clients/producer/Partitioner.java`
  - **Lines**: This file defines the `Partitioner` interface, starting at **line 20**.

## 8. Fault Tolerance and High Availability

- **Failure Detection**:
  - **File**: `kafka/server/ReplicaManager.scala`
  - **Lines**: The failure detection mechanisms, including heartbeats and ISR handling, are around **lines 600–700**.
- **Leader Failover**:
  - **File**: `kafka/controller/KafkaController.scala`
  - **Lines**: Failover and reassignment logic is located around **lines 300–500**.

## 9. Broker Configuration and Tuning

- **Dynamic Broker Configurations**:
  - **File**: `kafka/server/DynamicConfig.scala`
  - **Lines**: Code for dynamic reconfigurations can be found around **lines 100–300**.
- **Quota Management**:
  - **File**: `kafka/server/ClientQuotaManager.scala`
  - **Lines**: Quota management and enforcement logic are in **lines 200–500**.

These line numbers are approximate since line counts may vary across Kafka versions. You can explore the current state and locate exact line numbers in the codebase on GitHub in the [Apache Kafka repository](https://github.com/apache/kafka).
