# CryptGO

A cryptocurrency exchange built in microservice architecture.

# Architecture

## Solution Architecture

### DevOps

- Docker for containerization.
- GCP for running instances of each M.S and dependencies on the cloud.
  - [Deprecated]: GCP free trial expired.

### Messaging

- Kafka :- Publisher/Consumer Model (Message Broker).
<!-- - gRPC :- . -->

### Database & Storage

- PostgreSQL for saving operation log, user balance history, order history and trade history.
- Redis for saving market data.
- Skiplist (in-memory matching engine) for business logic and use cases.
  - Records user balance and executes user order.
  - Saves operation logs in MySQL and redoes the operation logs upon start.
  - Writes user history into MySQL, push balance, orders and deals messages to kafka.

### Authentication & Security

- Auth0 (golang JWTs).

<!-- ### Services

- Gateway Service
- Match Service
- Memory Store Service
- Wallet Service
- WS Service

- Rest is TBD -->
