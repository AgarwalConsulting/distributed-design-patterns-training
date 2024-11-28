layout: true

.signature[@algogrit]

---

class: center, middle

# Distributed Design Patterns

Gaurav Agarwal

---

# Agenda

- Distributed Systems are an *Engineering Marvel* <sup>`*`</sup>

.caveat[`*` *when done right*]

---

class: center, middle

![Me](assets/images/me.png)

Software Engineer & Product Developer

Director of Engineering & Founder @ https://codermana.com

ex-Tarka Labs, ex-BrowserStack, ex-ThoughtWorks

---

## As an instructor

- I promise to

  - make this class as interactive as possible

  - use as many resources as available to keep you engaged

  - ensure everyone's questions are addressed

---

## What I need from you

- Be vocal

  - Let me know if there are any audio/video issues ASAP

  - Feel free to interrupt me and ask me questions

- Be punctual

- Give feedback

- Work on the exercises

- Be *on mute* unless you are speaking

---
class: center, middle

## Class progression

![Learning Curve](assets/images/learning-curve.jpg)

---
class: center, middle

Here you are trying to *learn* something, while here your *brain* is doing you a favor by making sure the learning doesn't stick!

---

### Some tips (1/2)

- Slow down => stop & think
  - listen for the questions and answer

- Do the exercises
  - not add-ons; not optional

- There are no dumb questions!

- Drink water. Lots of it!

.caveat[continued...]

---

### Some tips (2/2)

- Take notes
  - Try: *Repetitive Spaced Out Learning*

- Talk about it out loud

- Listen to your brain

- *Experiment!*

.caveat[...continued]

---
class: center, middle

### 📚 Content ` > ` 🕒 Time

---
class: center, middle

## Show of hands

*Yay's - in Chat*

---
class: center, middle

## Application Architecture Overview

---
class: center, middle

Let's begin with...

---
class: center, middle

### Monoliths

---
class: center, middle

Monolithic architecture is a traditional software development approach where an application is built as a single, unified unit.

---
class: center, middle

In this architecture, all components of the application, such as the user interface (UI), business logic, and database access, are **tightly** coupled and operate as a single program.

---
class: center, middle

#### What exactly is a monolithic application?

---

Is it an application with a single codebase, deployed as a single service with a database?

![Traditional Monolith](assets/images/monolith-traditional-wt.png)

---
class: center, middle

or

---

Is it a modular application, deployed as a single service with a single database?

![Modular Monolith](assets/images/modular-monolith-wt.png)

---
class: center, middle

or

---

Is it a multi-services based application using an event-driven architecture and a database?

![Modular Distributed Microservice](assets/images/modular-distributed-monolith-wt.png)

---

Which of these is **NOT** a monolithic application?

![Monolith the Question](assets/images/monolith-the-question.png)

---

Well, **ALL** of them are!

![Monolith is everywhere](assets/images/monolith-is-everywhere.png)

.content-credits[https://tech.tamara.co/monolith-architecture-5f00270f384e]

---
class: center, middle

A single component failure can disrupt the entire application.

---
class: center, middle

With a distributed monolith, the *single component* is the **database**!

---
class: center, middle

They are all sharing the same database instance, with probable tight coupling between the data models.

---

#### Key Characteristics of Monolithic Architecture

- **Unified Codebase**

  All functionalities are contained within a single codebase and are compiled together into a single executable or deployable unit.

- **Tight Coupling**

  All components are interconnected and interdependent, making it difficult to change one component without affecting others.

- **Single Deployment Unit**

  The entire application is deployed as a single artifact, such as a WAR, JAR, or executable file.

- **Shared Database**

  Typically uses a single, centralized database to manage data for the entire application.

---

#### Advantages of Monolithic Architecture

- **Simplicity**

  Easier to develop, test, and deploy for smaller applications.

- **Performance**

  No inter-service communication overhead since everything runs in a single process.

- **Straightforward Debugging**

  Debugging and logging are simpler because the application runs as a single entity.

- **Development Tools**

  A wide variety of tools and frameworks are available to support monolithic architectures.

---

#### Challenges with Monolithic Architecture

- **Scalability Challenges**

  Difficult to scale individual components since the entire application must scale together.

- **Reduced Agility**

  Harder to adopt new technologies or make changes without risking the entire system.

- **Slow Deployment**

  Even small updates require redeploying the entire application.

- **Complexity with Size**

  As the application grows, the codebase can become large and difficult to manage (the "monolith monster").

- **Reliability Issues**

  A failure in one component can bring down the entire application.

---
class: center, middle

### Microservices

---
class: center, middle

A microservice is an architectural style that structures an application as a collection of small, independent, and **loosely** coupled services.

---
class: center, middle

Each microservice focuses on a specific business capability and operates autonomously.

---
class: center, middle

These services communicate with each other via well-defined APIs, typically using lightweight protocols such as HTTP/REST, gRPC, or messaging systems like Kafka.

---

An **e-commerce platform** can use microservices like:

- **Catalog Service**: Handles product listings.

- **Cart Service**: Manages shopping cart functionality.

- **Payment Service**: Processes payments.

- **User Service**: Manages user profiles.

- **Order Service**: Tracks order processing and history.

---
class: center, middle

Each of these services operates independently but collaborates to provide a seamless user experience.

---
class: center, middle

![E-commerce Microservice](assets/images/ecommerce-microservice.png)

.image-credits[https://medium.com/design-microservices-architecture-with-patterns/design-e-commerce-applications-with-microservices-architecture-c69e7f8222e7]

---

#### Key Characteristics of Microservices (1/2)

- **Single Responsibility**

  Each microservice is designed to perform a specific function or business capability (e.g., user management, payment processing).

- **Independence**

  Services can be developed, deployed, and scaled independently without affecting other components of the application.

- **Decentralized Data Management**

  Each microservice often manages its own database, promoting data autonomy and reducing bottlenecks.

.caveat[continued...]

---

#### Key Characteristics of Microservices (2/2)

- **Technology Agnostic**

  Different microservices can use different programming languages, frameworks, and data storage technologies as per their requirements.

- **Resilience**

  A failure in one service does not necessarily bring down the entire system due to fault-tolerance mechanisms and service isolation.

- **Continuous Delivery and Deployment**

  Enables faster development cycles and easier integration of new features or updates.

.caveat[...continued]

---

#### Advantages of Microservices

- **Scalability**

  Individual services can scale independently based on demand.

- **Flexibility**

  Teams can work on different services simultaneously, speeding up development.

- **Resilience**

  Isolated failures reduce the risk of a complete system outage.

- **Improved Maintainability**

  Smaller codebases are easier to understand, test, and maintain.

- **Technology Diversity**

  Allows teams to choose the best tools for each service.

---

#### Challenges with Microservices

- **Complexity**

  Managing many services increases the complexity of the system.

- **Deployment and Monitoring**

  Requires sophisticated tools for orchestration (e.g., Kubernetes) and monitoring.

- **Inter-Service Communication**

  Ensuring reliable and efficient communication is critical.

- **Data Consistency**

  Maintaining consistency across services can be challenging, especially with distributed databases.

---
class: center, middle

### Client/Server Application

---
class: center, middle

A client/server application is a distributed system where the application is divided into two main components: the client and the server.

---
class: center, middle

The client initiates requests, and the server processes those requests and provides responses. This model separates the user interface and user interactions (client-side) from the data storage and processing (server-side), making it a fundamental architecture for many modern applications.

---
class: center, middle

Generally, implemented over the internet (or intranet)

---
class: center, middle

#### OSI Layers - Quick Refresher

![OSI Layers](assets/images/osi-layers.png)

.image-credits[https://www.bmc.com/blogs/osi-model-7-layers/]

---
class: center, middle

### SOA

---
class: center, middle

Service-Oriented Architecture (SOA) is an architectural style where an application is composed of loosely coupled, reusable services that communicate over a network.

---
class: center, middle

Each service is designed to perform a specific business function and is accessible through a standard interface, making it easy to integrate services from different platforms, technologies, or organizations.

---
class: center, middle

## Communication

---

Microservices/SOA interact with each other primarily through:

- **Synchronous communication**: Using APIs like SOAP, REST or gRPC.

- **Asynchronous communication**: Using message brokers like RabbitMQ, Apache Kafka, or Amazon SQS.

---
class: center, middle

### SOAP

---
class: center, middle

SOAP (Simple Object Access Protocol) is a protocol for exchanging structured information between systems over a network.

---
class: center, middle

It is a widely used messaging protocol in Service-Oriented Architectures (SOA) and enables communication between applications, regardless of platform or programming language.

---

#### Key Features of SOAP

- **XML-Based**

  SOAP messages are encoded in XML, making them human-readable and platform-independent.

- **Protocol-Agnostic**

  SOAP can use various protocols for communication, including HTTP, HTTPS, SMTP, and more.

- **Standards-Based**

  It is defined by the **W3C** and follows strict standards for message formatting and communication.

- **Extensibility**

  Supports additional functionalities through headers, such as security, transactions, and routing.

- **Error Handling**

  Provides a structured way to return error information via fault messages.

---
class: center, middle

#### Structure of a SOAP Message

.caveat[~optional content~]

---

A SOAP message is composed of the following elements:

- **Envelope**

  - The root element that defines the message structure and namespaces.
  - Example:

    ```xml
    <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
      ...
    </soap:Envelope>
    ```

- **Header** (Optional)

  - Contains metadata, such as authentication information or transaction context.
  - Example:

    ```xml
    <soap:Header>
      <authToken>ABC123</authToken>
    </soap:Header>
    ```

.caveat[~optional content~]

---

- **Body**

  - Contains the actual message or data being exchanged.
  - Example:

    ```xml
    <soap:Body>
      <getUserDetails>
        <userId>123</userId>
      </getUserDetails>
    </soap:Body>
    ```

- **Fault** (Optional)

  - A sub-element of the body used to return error details if something goes wrong.
  - Example:

    ```xml
    <soap:Fault>
      <faultcode>Client</faultcode>
      <faultstring>Invalid user ID</faultstring>
    </soap:Fault>
    ```

.caveat[~optional content~]

---

#### Example of a SOAP Request

A SOAP request to fetch user details might look like this:

```xml
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <getUserDetails xmlns="http://example.com/user">
      <userId>123</userId>
    </getUserDetails>
  </soap:Body>
</soap:Envelope>
```

---

#### Example of a SOAP Response

The server might respond with:

```xml
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <getUserDetailsResponse xmlns="http://example.com/user">
      <user>
        <id>123</id>
        <name>John Doe</name>
        <email>john.doe@example.com</email>
      </user>
    </getUserDetailsResponse>
  </soap:Body>
</soap:Envelope>
```

---

#### Advantages of SOAP

- **Platform and Language Agnostic**

  Works with any platform and programming language.

- **Extensive Standards**

  Comprehensive standards for security (e.g., WS-Security), transactions, and reliability.

- **Support for Complex Operations**

  Handles complex operations like distributed transactions and stateful interactions.

- **Error Handling**

  Well-defined fault messages for robust error reporting.

---

#### Disadvantages of SOAP

- **Verbosity**

  SOAP messages can be large due to XML formatting, leading to increased network usage.

- **Complexity**

  The strict standards and extensive features make SOAP more complex than other protocols like REST.

- **Performance**

  Parsing XML can be slower compared to other lightweight formats like JSON.

---
class: center, middle

### ReST

**Re**presentational **S**tate **T**ransfer

---
class: center, middle

ReST is an architectural style for designing networked applications.

---
class: center, middle

It uses standard HTTP methods to interact with resources identified by URLs, and it emphasizes simplicity, scalability, and stateless communication.

---

#### HTTP Methods in REST

| **Method**   | **Description**                              | **Example**               |
|--------------|----------------------------------------------|---------------------------|
| **GET**      | Retrieve a list of resources                 | `GET /users`              |
| **GET**      | Retrieve a resource-based                    | `GET /users/123`          |
| **POST**     | Create a new resource                        | `POST /users`             |
| **PUT**      | Update an existing resource or create it     | `PUT /users/123`          |
| **PATCH**    | Partially update a resource                  | `PATCH /users/123`        |
| **DELETE**   | Delete a resource                            | `DELETE /users/123`       |

---

#### Example Request to Retrieve a User

**GET** `https://api.example.com/users/123`

**Response**:
```json
{
  "id": 123,
  "name": "John Doe",
  "email": "john.doe@example.com"
}
```

---

#### Example Request to Create a New User

**POST** `https://api.example.com/users`

**Request Body**:

```json
{
  "name": "Jane Smith",
  "email": "jane.smith@example.com"
}
```

**Response**:

```json
{
  "id": 124,
  "name": "Jane Smith",
  "email": "jane.smith@example.com"
}
```

---

#### Key Principles of REST (1/2)

- **Client-Server Architecture**

  Separation of concerns between the client (user interface) and the server (data storage and processing).

- **Statelessness**

  Each request from a client to a server must contain all the information the server needs to process it. No client context is stored on the server between requests.

- **Cacheability**

  Responses should indicate whether they are cacheable to improve performance and scalability.

---

#### Key Principles of REST (2/2)

- **Uniform Interface**

  REST APIs should follow standard conventions, making them easy to understand and use. Key components include:

  - **Resources**: Identified by URIs.
  - **HTTP Methods**: Standard operations like GET, POST, PUT, DELETE.
  - **Representation**: Data can be represented in formats like JSON, XML, or HTML.

- **Layered System**

  A RESTful system can be designed with layers (e.g., load balancers, caching systems) without affecting the client-server communication.

- **Code on Demand (Optional)**

  Servers can provide executable code to clients, such as JavaScript, to extend functionality dynamically.

---

#### Advantages of REST

- **Simplicity**

  Relies on standard HTTP methods and URIs, making it straightforward to understand and implement.

- **Scalability**

  Stateless communication and caching improve scalability.

- **Flexibility**

  Supports multiple formats (JSON, XML, etc.), making it versatile for various use cases.

- **Interoperability**

  Works across different platforms and languages using standard web protocols.

- **Performance**

  Caching mechanisms and lightweight formats like JSON ensure efficient communication.

---

#### Disadvantages of REST

- **Stateless Nature**

  Statelessness can lead to redundant data being sent in each request, impacting performance.

- **Overhead**

  Using HTTP headers and status codes can add some complexity compared to simpler communication models.

- **Limited for Complex Operations**

  May require workarounds for scenarios like long-running transactions or real-time communication.

---

### REST vs SOAP

| **Aspect**              | **REST**                           | **SOAP**                            |
|-------------------------|------------------------------------|-------------------------------------|
| **Protocol**            | Architectural style (HTTP-based)   | Strict protocol (XML-based)         |
| **Data Format**         | JSON, XML, or others               | Only XML                            |
| **Ease of Use**         | Lightweight and simple             | Verbose and complex                 |
| **State**               | Stateless                          | Stateful or stateless               |
| **Security**            | Relies on HTTPS and OAuth          | Built-in standards like WS-Security |
| **Flexibility**         | Highly flexible                    | Rigid standards                     |

REST is ideal for modern web and mobile applications where simplicity and scalability are key.

---
class: center, middle

### gRPC

---
class: center, middle

#### Why gRPC?

---
class: center, middle

Is it Protobuf?

---
class: center, middle

gRPC uses Protobuf for serialization

---

Protobuf offers:

- smaller payloads

- faster serialization/deserialization

compared to text-based formats like `JSON` or `XML`.

---
class: center, middle

But, you can use protobuf with ReST too!

---
class: center, middle

So, why gRPC!?

---
class: center, middle

Let's look at TCP vs UDP...

---
class: center, middle

![TCP vs UDP](assets/images/tcp-vs-udp.png)

---
class: center, middle

Every TCP new TCP connection requires a new handshake!

---
class: center, middle

gRPC avoids this by using **HTTP 2**!

---
class: center, middle

![HTTP 1 vs HTTP 2](assets/images/http_1_vs_2.png)

---

#### Key Features of gRPC (1/2)

- **Efficient Communication**

  gRPC uses Protobuf for serialization, which is smaller and faster compared to text-based formats like JSON or XML.

- **Cross-Language Support**

  Supports many programming languages (e.g., C++, Java, Python, Go, etc.), allowing seamless communication in heterogeneous environments.

- **Streaming**

  Supports real-time communication through streaming:

  - **Unary RPC**: One request and one response.

  - **Server Streaming**: One request and multiple responses.

  - **Client Streaming**: Multiple requests and one response.

  - **Bidirectional Streaming**: Multiple requests and responses in real-time.

---

#### Key Features of gRPC (2/2)

- **Multiplexing**

  Built on **HTTP/2**, gRPC enables multiplexing of multiple requests over a single connection.

- **Built-in Code Generation**

  Protobuf definitions generate client and server code automatically, reducing boilerplate and ensuring consistency.

- **Strong Typing**

  Protobuf enforces strong typing, catching errors early during development.

---

#### Architecture of gRPC

- **Client**

  - Makes calls to the server using generated client stubs.

  - Stubs are auto-generated from Protobuf definitions.

- **Server**

  Implements the service as defined in the Protobuf file and handles incoming RPCs.

- **Protocol Buffers (Protobuf)**

  Defines the data structures and services in a `.proto` file.

- **Transport Layer**

  Uses HTTP/2 for communication, offering features like compression, multiplexing, and low latency.

---

#### Protobuf Definition (`service.proto`)

```proto
syntax = "proto3";

service UserService {
  rpc GetUser (UserRequest) returns (UserResponse);
}

message UserRequest {
  int32 user_id = 1;
}

message UserResponse {
  int32 user_id = 1;
  string name = 2;
  string email = 3;
}
```

---

#### Advantages of gRPC

- **High Performance**

  Efficient serialization (Protobuf) and HTTP/2 reduce overhead, making gRPC faster than REST for many use cases.

- **Streaming Support**

  Real-time data transmission for use cases like chat applications or live monitoring.

- **Language Interoperability**

  Provides consistent communication across multiple programming languages.

- **Auto-Generated Code**

  Reduces manual effort and potential errors.

- **Scalability**

  Well-suited for microservices and large-scale systems.

---

#### Disadvantages of gRPC

- **Steep Learning Curve**

  Requires familiarity with Protobuf and the gRPC framework.

- **Limited Browser Support**

  Native gRPC does not work in browsers due to HTTP/2 limitations, requiring additional tools like gRPC-Web.

- **Complexity**

  More setup and tooling compared to simpler architectures like REST.

- **Debugging**

  Debugging binary-encoded Protobuf messages can be more challenging than JSON or XML.

---
class: center, middle

### Honorable Mentions: WebSockets, graphQL

---
class: center, middle

### Messaging

---
class: center, middle

Messaging in software architecture refers to the exchange of information between components or systems in the form of messages.

---
class: center, middle

It is a communication mechanism used in distributed systems to enable decoupled interactions between services, applications, or devices.

---

#### Key Concepts in Messaging (1/2)

- **Messages**

  The unit of communication that carries data between components.

  - Messages can be structured (e.g., JSON, XML) or binary (e.g., Protocol Buffers).

- **Message Queue**

  A buffer where messages are temporarily stored until the receiving component processes them.

- **Publish-Subscribe Model**

  A messaging pattern where publishers send messages to a topic, and multiple subscribers receive those messages.

---

#### Key Concepts in Messaging (2/2)

- **Point-to-Point Messaging**

  A pattern where a message is sent from a producer to a single consumer via a queue.

- **Event-Driven Architecture**

  Messaging enables systems to respond to events asynchronously, improving scalability and responsiveness.

---

#### Messaging Patterns (1/2)

- **Request-Response**

  A sender sends a request and waits for a specific response.

  - Example: A client requesting data from a server.

- **Fire-and-Forget**

  A sender sends a message and does not expect a response.

  - Example: Logging messages to a monitoring system.

- **Publish-Subscribe**

  Publishers broadcast messages to topics, and subscribers receive messages from those topics.

  - Example: Notifications or real-time chat systems.

---

#### Messaging Patterns (2/2)

- **Message Routing**

  Messages are directed to specific recipients based on rules or headers.

  - Example: Filtering messages in an email system.

- **Broadcasting**

  Messages are sent to all connected consumers.

  - Example: Live sports updates.

---
class: center, middle

#### Messaging Middleware

Messaging middleware, also known as **Message Brokers**, facilitates message exchange between distributed systems.

---

Examples include:

- **RabbitMQ**:

  - Open-source message broker supporting AMQP protocol.

  - Great for task queues and real-time messaging.

- **Apache Kafka**:

  - Distributed event streaming platform.

  - Ideal for high-throughput, real-time event processing.

- **Amazon SQS**:

  - Fully managed queue service by AWS.

  - Asynchronous processing and task decoupling.

- **ActiveMQ**:

  - Open-source message broker with support for multiple protocols.

---

#### Advantages of Messaging

- **Decoupling**

  Components do not need to know about each other; they interact through messages.

- **Asynchronous Communication**

  Improves system responsiveness and scalability.

- **Scalability**

  Components can process messages at their own pace, handling load spikes effectively.

- **Fault Tolerance**

  Messages can be persisted, ensuring delivery even if components are temporarily unavailable.

- **Real-Time Processing**

  Messaging enables real-time systems, such as notifications or event streaming.

---

#### Disadvantages of Messaging

- **Complexity**

  Introducing messaging middleware adds complexity to the system architecture.

- **Latency**

  Asynchronous communication may introduce delays.

- **Message Loss**

  Improper configuration or failure of brokers can lead to message loss.

- **Dependency**

  Messaging systems can become a single point of failure without proper redundancy.

---
class: center, middle

Messaging is the backbone of modern distributed systems, enabling asynchronous and scalable communication between services.

---
class: center, middle

## [Distributed vs Cloud vs Serverless](https://github.com/AgarwalConsulting/distributed-design-patterns-training/blob/master/examples/01-architecture-overview/distributed-vs-cloud-vs-serverless.md)

---
class: center, middle

## `CAP` vs `BASE` vs `PACLEC`

.content-credits[https://medium.com/@ali.gelenler/distributed-system-trade-offs-cap-vs-base-vs-pacelc-1a3bcac04a7b]

---
class: center, middle

### CAP theorem

defines the limitations and trade-offs in a distributed system

![CAP Theorem](assets/images/cap-theorem.png)

---

It suggests that distributed computer systems can only deliver two out of the following three guarantees:

**Consistency**: Every node sees the same data even when concurrent updates occur

**Availability**: All requests receive responses on whether it was a success or a failure

**Partition tolerance**: The system will keep operating even if there is a network partition in communication between different nodes

---
class: center, middle

In the case of a network partition, the CAP theorem forces a trade-off between *Consistency* and *Availability*.

---

A system must either:

- Maintain consistency, but sacrifice availability (not all requests are responded to).

- Maintain availability, but sacrifice consistency (some responses may be outdated).

---
class: center, middle

### BASE

is an acronym that represents an alternative to strict `ACID` properties

![BASE Model](assets/images/base-model.png)

---
class: center, middle

ACID properties

(**A**tomicity, **C**onsistency, **I**solation, **D**urability)

---

class: center, middle

It focuses on availability and performance over strong consistency.

---

- Basically Available (BA): The system is guaranteed to be available, but not necessarily consistent.

- Soft State (S): The state of the system may change over time, even without input (because of eventual consistency).

- Eventual Consistency (E): The system will become consistent eventually, once all updates have propagated.

---
class: center, middle

BASE is a practical strategy to design systems that are highly available and partition tolerant, at the cost of immediate consistency.

---
class: center, middle

It is often used in distributed, NoSQL systems, which prioritize scalability and availability.

---
class: center, middle

### PACLEC

is an Extension to the CAP theorem

![PACLEC](assets/images/paclec.png)

---
class: center, middle

It acknowledges that network partitions are not the only factor affecting system performance and behavior.

---
class: center, middle

It also considers the trade-offs between latency and consistency during normal operations, which the CAP theorem does not address.

---

- *PAC*: If there is a Partition, you must choose between Availability and Consistency (this is the CAP theorem)

- *ELC*: Else, when the system is running normally (no partition), you must choose between Latency and Consistency

---

In other words:

- If there’s a partition, the system must decide between availability and consistency.

- If there’s no partition, the system must decide between minimizing latency (faster responses) and ensuring strong consistency.

---
class: center, middle

## Application Design for Scalability

---
class: center, middle

### [12 Factor Principles](https://12-factor-apps.slides.algogrit.com/)

---
class: center, middle

Code
https://github.com/AgarwalConsulting/distributed-design-patterns-training

Slides
https://distributed-design-patterns.slides.agarwalconsulting.com
