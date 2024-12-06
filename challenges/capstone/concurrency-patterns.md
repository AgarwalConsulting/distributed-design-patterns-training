# Problem Statement: Distributed Task Processing System

## Objective

Design and implement a highly scalable and fault-tolerant distributed task processing system that supports concurrent execution of tasks, monitors system health, dynamically adjusts resource allocation, and ensures consistent ordering and synchronization of tasks across a distributed environment.

---

## Requirements

## Core Functionalities

1. **Task Submission and Execution**

- Allow clients to submit tasks to the system.

- Tasks may have dependencies (e.g., Task B can only execute after Task A completes).

- Tasks should execute concurrently where possible, while respecting dependencies.

2. **Dynamic Scheduling**

- Use a **dynamic scheduler** to prioritize tasks based on:

  - Deadlines

  - Resource availability

  - Task dependencies

- Adjust task priorities dynamically based on real-time system conditions (e.g., load, failures).

3. **Thread Pool for Execution**

- Use a **thread pool** to execute tasks concurrently while managing system resource usage effectively.

- Ensure that the number of threads scales with system load.

4. **Monitoring and Health Checks**

- Implement the **Proctor pattern** to monitor the health of worker nodes and the status of tasks.

- Detect and recover from failures (e.g., reschedule tasks if a worker node crashes).

5. **Event-Driven Task Management**

- Use the **Reactor pattern** to handle task events asynchronously (e.g., task completed, task failed, new task submitted).

- Events should trigger appropriate actions such as notifying dependent tasks or re-queuing failed tasks.

6. **Global Consistency in Task Ordering**

- Use a **Lamport clock** or **Hybrid clock** to ensure consistent task ordering and synchronization across distributed worker nodes.

7. **Asynchronous Completion Notification**

- Use the **Async Completion Token** pattern to notify clients about task completion, providing them with results or errors.

---

## Non-Functional Requirements

1. **Scalability**

- The system should handle a large number of tasks and scale horizontally by adding more worker nodes.

2. **Fault Tolerance**

- Recover gracefully from node failures without losing tasks.

- Ensure task execution completes even in the presence of partial failures.

3. **Low Latency**

- Tasks should execute with minimal delays, even under heavy load.

4. **Extensibility**

- The system should allow new features (e.g., new scheduling algorithms) to be added without major architectural changes.

---

## Example Use Case

A large-scale e-commerce company wants to build a **real-time order processing system**. Tasks include:

1. **Payment processing** Must complete before shipping starts.

2. **Inventory updates** Must be done after payment is confirmed.

3. **Shipping** Must start after inventory update is complete.

The system must handle thousands of concurrent orders, prioritize VIP customers, and recover from payment gateway failures or inventory update delays.

---

## Deliverables

1. **System Design Document**

- Describe how the patterns are integrated.

- Include architectural diagrams for components like task queues, worker nodes, and the scheduler.

2. **Implementation**

- Develop a prototype system in Go or another preferred language.

3. **Testing Plan**

- Simulate various scenarios like task dependencies, node failures, and high load.

- Validate system performance, scalability, and fault tolerance.

4. **Deployment Guide**

- Instructions for deploying the system in a distributed environment.
