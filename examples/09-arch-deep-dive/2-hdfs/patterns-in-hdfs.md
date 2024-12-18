# Patterns in HDFS

Hadoop Distributed File System (HDFS) applies several distributed systems patterns to achieve its scalability, fault tolerance, and high performance.

---

## 1. Master-Slave Architecture

- **Pattern:** Centralized Coordination

- **Application in HDFS:**
  - HDFS uses a master-slave architecture, where the **NameNode** acts as the master and **DataNodes** as slaves.
  - The NameNode manages the filesystem metadata (e.g., namespace, file-to-block mappings), while the DataNodes store the actual data blocks.
  - This central coordination simplifies metadata management and global consistency.

---

## 2. Data Partitioning

- **Pattern:** Sharding

- **Application in HDFS:**
  - Large files are split into fixed-size blocks (default: 128 MB or 256 MB).
  - These blocks are distributed across multiple DataNodes.
  - Partitioning allows parallel processing and efficient resource utilization in distributed environments.

---

## 3. Replication

- **Pattern:** Replication

- **Application in HDFS:**
  - Each data block is replicated (default replication factor: 3) and stored on multiple DataNodes.
  - This ensures fault tolerance—if one DataNode fails, data is still accessible from other replicas.

---

## 4. Fault Tolerance and Failure Detection

- **Pattern:** Redundancy and Heartbeat Monitoring

- **Application in HDFS:**
  - DataNodes send periodic heartbeats and block reports to the NameNode.
  - If a DataNode fails to send a heartbeat, the NameNode marks it as failed and initiates replication to maintain the desired replication factor.
  - This ensures the system is resilient to node failures.

---

## 5. Write-Once, Read-Many

- **Pattern:** Immutable Data

- **Application in HDFS:**
  - HDFS files are designed to be written once and read multiple times.
  - This simplifies consistency and concurrency issues, enabling efficient distributed storage and retrieval.

---

## 6. Data Locality

- **Pattern:** Compute-Storage Co-location

- **Application in HDFS:**
  - HDFS is designed to run on the same nodes as the computation frameworks (like Hadoop MapReduce).
  - Tasks are scheduled on nodes where the data resides (data locality), reducing network overhead and improving performance.

---

## 7. High Availability

- **Pattern:** Active-Passive Failover

- **Application in HDFS:**
  - HDFS supports high availability using a primary NameNode (active) and a standby NameNode (passive).
  - Shared storage (e.g., using a journal) keeps the metadata synchronized between the two NameNodes.
  - If the active NameNode fails, the standby NameNode takes over.

---

## 8. Distributed Consensus

- **Pattern:** Leader Election and Coordination

- **Application in HDFS:**
  - In high-availability setups, ZooKeeper is often used for leader election and coordination between the active and standby NameNodes.
  - This ensures a consistent and reliable failover process.

---

## 9. Client-Side Metadata Caching

- **Pattern:** Caching

- **Application in HDFS:**
  - Clients cache metadata from the NameNode (e.g., block locations).
  - This reduces the load on the NameNode and improves performance by minimizing metadata queries.

---

## 10. Batch Processing

- **Pattern:** Bulk Data Processing

- **Application in HDFS:**
  - HDFS is optimized for large-scale batch processing workloads where high throughput is prioritized over low latency.
  - This makes it ideal for systems like Hadoop MapReduce and Spark.

---

## 11. Eventual Consistency

- **Pattern:** Relaxed Consistency Model

- **Application in HDFS:**
  - HDFS ensures eventual consistency for block replicas.
  - Temporary inconsistencies (e.g., during replica creation or recovery) are resolved over time, maintaining system availability.

---

These distributed systems patterns collectively make HDFS a robust and efficient distributed storage solution for large-scale data processing.
