# CryptGO

A cryptocurrency exchange built in microservice architecture.

It follows from the ViaBTC exchange server, by design.

<img src="https://github.com/ReshiAdavan/CryptGO/blob/master/imgs/CryptGO_Arch.PNG" />

## Inspiration

Before I was invested in web3.0, blockchain, decentralized systems, etc., as mentioned in [Catena](https://github.com/ReshiAdavan/Catena), I was interested in Crypto because of a Senior Engineer I worked with over Winter 2023. Looking into Crypto and exchanges, I wanted to make one (at this point whenever I want to learn more of something I make a tool that focuses on it/revolves around it). Thats exactly what I did.

The choice of cryptocurrency was ethereum because of [Vitalik Buterin](https://www.linkedin.com/in/vitalik-buterin-267a7450/?originalSubdomain=ca), the infamous Waterloo student who founded [Ethereum](https://ethereum.org/en/). At the very least I am inspired by him and motivated to be like him.

## Architecture

### DevOps

- Docker for containerization.
- GCP for running instances of each M.S and dependencies on the cloud.
  - [Will Deprecate]: GCP free trial expiry imminent.

### Messaging

- Kafka :- Publisher/Consumer Model (Message Broker).
- gRPC :- Microservice Cross-Communication.

### Database & Storage

- PostgreSQL for saving operation log, user balance history, order history and trade history.
- Redis for saving market data.
- Skiplist (in-memory matching engine) for business logic and use cases.
  - Records user balance and executes user order.
  - Saves operation logs in MySQL and redoes the operation logs upon start.
  - Writes user history into MySQL, push balance, orders and deals messages to kafka.

### Authentication & Security

- Auth0 (Golang JWTs).

### Services

- Gateway Service
- Match Service
- Memory Store Service
- Wallet Service
- WS Service
