# Default Message delivery guarantees

## 1. RabbitMQ

- **Default Delivery Guarantee:** **At-Least-Once**
  - Messages are acknowledged when the consumer confirms receipt.
  - If no acknowledgment is received, RabbitMQ will re-deliver the message.
- **Optional Guarantees:**
  - **At-Most-Once:** Disable acknowledgments or use `auto-ack`.
  - No built-in **Exactly-Once**, but can be approximated with idempotent consumers and careful setup.

---

## 2. ZeroMQ

- **Default Delivery Guarantee:** **Best-Effort** (No guarantees by default)
  - Messages are sent over a socket, and there is no built-in persistence or acknowledgment mechanism.
- **Optional Guarantees:**
  - Delivery guarantees can be implemented at the application layer (e.g., retries, acknowledgments, persistence).

---

## 3. Apache Kafka

- **Default Delivery Guarantee:** **At-Least-Once**
  - Messages are persisted to disk on the broker and will remain available until consumed and acknowledged.
  - Consumers may see duplicate messages in failure scenarios unless deduplication is handled.
- **Optional Guarantees:**
  - **At-Most-Once:** Consumers commit offsets before processing messages.
  - **Exactly-Once:** Supported natively with transactional producers and consumers.

---

## 4. Amazon SQS

- **Default Delivery Guarantee:**
  - **Standard Queues:** **At-Least-Once**, with the possibility of occasional duplicates.
  - **FIFO Queues:** **Exactly-Once** (with deduplication enabled) and in-order delivery.
- **Optional Guarantees:**
  - No further options; guarantees depend on the type of queue chosen.

---

## 5. ActiveMQ

- **Default Delivery Guarantee:** **At-Least-Once**
  - Messages are acknowledged by default when successfully delivered.
  - If no acknowledgment is received, the broker retries delivery.
- **Optional Guarantees:**
  - **At-Most-Once:** Configurable with acknowledgment modes.
  - No built-in **Exactly-Once**, but can be approximated with external deduplication or transactional messaging.

---

## 6. NATS

- **Default Delivery Guarantee:**
  - **Core NATS:** **At-Most-Once**
    - Messages are sent to subscribers without acknowledgment or persistence. If a subscriber is unavailable, the message is dropped.
  - **NATS JetStream:** **At-Least-Once**
    - Enables persistence and message acknowledgments.
- **Optional Guarantees:**
  - **Exactly-Once:** Not supported even with JetStream. Applications must handle deduplication if needed.

---

## Comparison Table

| Feature/Requirement  | RabbitMQ          | ZeroMQ           | Kafka            | SQS                     | ActiveMQ         | NATS             |
|-----------------------|-------------------|------------------|------------------|-------------------------|------------------|------------------|
| **Default Guarantee** | At-Least-Once    | Best-Effort      | At-Least-Once    | At-Least-Once (Standard) / Exactly-Once (FIFO) | At-Least-Once    | At-Most-Once (Core) / At-Least-Once (JetStream) |
| **At-Least-Once**     | Default          | Application-Defined | Default          | Standard Queues         | Default          | JetStream Option |
| **At-Most-Once**      | Configurable     | Default          | Configurable     | Not Applicable          | Configurable     | Default (Core)   |
| **Exactly-Once**      | No, Approximate  | No, Custom Logic | Yes, with Transactions | FIFO Queues          | No, Approximate  | No, Custom Logic |

---

## Summary

- **At-Least-Once:** Default for RabbitMQ, Kafka, ActiveMQ, SQS Standard, and NATS JetStream.
- **At-Most-Once:** Default for ZeroMQ and Core NATS.
- **Exactly-Once:** Native in Kafka (with transactions) and SQS FIFO; other systems require custom logic.
