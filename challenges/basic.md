# Basic Problems

1. **Problem Statement**

Design a system that listens for log messages from multiple sources and writes them to a file asynchronously. Ensure that the system can handle a high volume of log messages without blocking producers.

- **Key Features**

  - Event-driven logging system.
  - Supports multiple log levels (e.g., info, error, debug).
  - Ensures logs are written in the order they are received.

---

2. **Problem Statement**

Develop a simple watchdog service that monitors the health of an application. If the application becomes unresponsive or unhealthy, the watchdog should restart it and log the event.

- **Key Features**

  - Periodically check application health via a heartbeat mechanism.
  - Restart the application if it stops responding.
  - Maintain a log of all health checks and restarts.

---

3. **Problem Statement**

Create a system that resizes images in parallel. Clients should be able to submit a batch of images, and the system should process them concurrently using a fixed number of worker threads.

- **Key Features**

  - Resize images to a specific dimension.
  - Use a thread pool to limit resource usage.
  - Provide progress updates to the client.

---

4. **Problem Statement**

Design a system that schedules tasks based on their priority. The system should dynamically adjust task priorities based on real-time conditions (e.g., task deadline or system load).

- **Key Features**

  - Dynamically re-prioritize tasks.
  - Ensure high-priority tasks are always executed first.
  - Allow tasks to be rescheduled in case of failure.

---

5. **Problem Statement**

Build a distributed counter system where multiple nodes increment a shared counter. Ensure that the increments are ordered correctly, even if some nodes update the counter concurrently.

- **Key Features**

  - Allow multiple nodes to increment the counter.
  - Maintain a consistent order of updates.
  - Resolve conflicts if two nodes update simultaneously.

---

6. **Problem Statement**

Implement a file downloader that allows clients to download multiple files concurrently. Each download should notify the client upon completion or failure, including a unique token for tracking.

- **Key Features**

  - Support for downloading multiple files at the same time.
  - Provide completion notifications with download status.
  - Retry downloads in case of failure.

---

7. **Problem Statement**

Design a system that manages access to a shared resource (e.g., a database connection) among multiple threads. Ensure that only one thread can access the resource at a time.

- **Key Features**

  - Support multiple threads requesting access to the shared resource.
  - Queue requests and ensure fair access.
  - Prevent deadlocks and race conditions.

---

8. **Problem Statement**

Create a simple chat server that supports multiple clients. The server should handle incoming messages asynchronously and broadcast them to all connected clients.

- **Key Features**

  - Handle multiple clients simultaneously.
  - Asynchronously receive and broadcast messages.
  - Efficiently manage resources for high concurrency.
