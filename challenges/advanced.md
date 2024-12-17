# Advanced System Design

## Exercise 1: Design a Global News Feed System

### Problem Statement

You are tasked with designing a **global news feed system** for a social media platform like Facebook or Twitter. The news feed must support:
1. Billions of users worldwide.
2. Real-time updates when a user posts new content.
3. Fast reads when a user fetches their news feed.

### Requirements

- **High Write Throughput:** Users frequently post updates that need to propagate quickly.
- **High Read Scalability:** Millions of users may read their feeds simultaneously.
- **Fault Tolerance:** System should be available even if some servers fail.
- **Consistency:** While strong consistency isn't mandatory, eventual consistency is required to ensure feeds are up-to-date.

### Tasks

1. **Design the Architecture:** Identify which distributed system patterns (e.g., **sharding**, **leader-follower**, **replication**, etc.) are needed to solve the problem.
2. **State Assumptions:** Outline assumptions about data volume, user behavior, and latency.
3. **Discuss Trade-offs:** For example:
   - Should you compute feeds in real time or precompute them?
   - How will sharding impact system scalability?
   - How can you balance between strong consistency and availability?

---

## Exercise 2: Design a Multi-Region E-Commerce Platform

### Problem Statement

Design the backend for a global e-commerce platform like Amazon that serves users across multiple regions. The platform must support:
1. Product catalog browsing.
2. Placing and processing orders.
3. Handling inventory across warehouses.

### Requirements

- **Low Latency:** Users in any region must experience minimal latency.
- **Scalable Reads and Writes:** Product searches, orders, and inventory updates must scale with traffic.
- **Consistency for Inventory:** Inventory counts need to be accurate to prevent overselling.
- **Fault Tolerance:** The system should remain available during server or regional failures.

### Tasks

1. **Choose and Justify Patterns:** Decide how to shard data (e.g., product catalog, orders, inventory) and how to replicate services across regions.
2. **Design High-Level Components:** Include components for product search, order processing, and inventory management.
3. **Address Fault Tolerance:** How will the system handle regional failures?
4. **Trade-offs to Consider:**
   - Should inventory data prioritize consistency (e.g., strong consistency) or availability?
   - How can you optimize for read-heavy product catalog queries?

---

## Exercise 3: Design a Real-Time Chat System

### Problem Statement

You are designing a real-time chat system like WhatsApp or Slack. The system must support:
1. Sending and receiving messages in real time.
2. Maintaining chat history for all users.
3. Ensuring availability even during server failures.

### Requirements

- **Low Latency:** Messages should appear in real time for users.
- **Message Storage:** Chat history must be stored persistently and retrievable at any time.
- **Scalability:** The system must scale to millions of concurrent users.
- **Fault Tolerance:** Ensure messages are not lost and the system remains available.

### Tasks

1. **Identify Patterns to Apply:**
   - How will you handle real-time messaging (e.g., via load balancing, leader-follower replication)?
   - How will you store and shard chat history to ensure scalability?
2. **Design the Architecture:** Include message queues, storage services, and user connection handling.
3. **Handle Failures:** Propose solutions for leader failure, message duplication, or dropped messages.
4. **Discuss Trade-offs:**
   - Should chat history be strongly consistent across all nodes?
   - How do you ensure low latency while guaranteeing message delivery?

---

## Exercise 4: Design a Distributed Video Processing System

### Problem Statement

Build a video processing pipeline for a platform like YouTube or TikTok where users upload videos, and the system processes them into different resolutions/formats.

### Requirements

- **Scalable Upload Handling:** Support millions of simultaneous uploads.
- **Parallel Video Processing:** Videos need to be processed quickly and efficiently into multiple formats.
- **Fault Tolerance:** Ensure processing can recover from server failures.
- **Metadata Storage:** Video metadata (e.g., title, resolution links) must be accessible at scale.

### Tasks

1. **Propose the System Design:**
   - Identify which patterns (e.g., **master-worker**, **sharding**, **load balancing**) to use for video processing and storage.
2. **Fault Tolerance and Recovery:** Explain how the system ensures no uploaded videos are lost during processing.
3. **Scalability:** Discuss how the system handles spikes in uploads.
4. **Discuss Trade-offs:**
   - Should video metadata use leader-follower replication to optimize reads?
   - How can you distribute processing tasks evenly across worker nodes?

---

## Exercise 5: Design a Global Payment Processing System

### Problem Statement

You are tasked with building a payment processing system for a financial platform like Stripe or PayPal. The system must handle:
1. Payment transactions from users globally.
2. Ensuring no double payments or data loss.
3. High availability with minimal downtime.

### Requirements

- **Consistency:** Payments must be strongly consistent (no double charges).
- **Scalability:** The system must scale to process millions of payments per minute.
- **Fault Tolerance:** Payments must not be lost even during failures.
- **Low Latency:** Payments should be processed quickly.

### Tasks

1. **Identify Suitable Patterns:**
   - How will you ensure strong consistency while scaling horizontally?
   - How will you replicate data for fault tolerance?
   - Do you need **leader election** for critical processes like transaction coordination?
2. **Propose System Architecture:** Design components for payment validation, transaction logs, and balance updates.
3. **Fault Handling:** How will you handle server crashes during ongoing transactions?
4. **Trade-offs:**
   - How do you balance latency and consistency?
   - Should reads (e.g., account balances) be served from follower replicas?

---

## Exercise 6: Design a Large-Scale Online Exam Platform

### Problem Statement

Design a platform like Google Forms or ProctorU for conducting large-scale online exams.

### Requirements

- **Scalability:** The platform must handle millions of concurrent users taking exams.
- **Real-Time Results:** Exams must be scored in real time.
- **Fault Tolerance:** Ensure exams are not disrupted if some servers fail.
- **Consistency:** User answers must be stored reliably without loss.

### Tasks

1. **Choose Patterns to Apply:**
   - How will you distribute exam traffic across multiple servers?
   - How will you shard and replicate exam data (e.g., questions, answers)?
2. **Design High-Level Architecture:** Include components for user management, answer storage, and scoring.
3. **Handle Failures:** Propose a strategy for recovering data if a server fails mid-exam.
4. **Discuss Trade-offs:**
   - Should scoring be done synchronously (real time) or asynchronously?
   - How will you replicate answers to ensure durability and consistency?
