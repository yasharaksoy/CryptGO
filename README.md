# CryptGO

&nbsp;&nbsp;&nbsp;&nbsp; A cryptocurrency exchange built in microservice architecture.

&nbsp;&nbsp;&nbsp;&nbsp; It follows from the ViaBTC exchange server, by design.

<img src="https://github.com/ReshiAdavan/CryptGO/blob/master/imgs/CryptGO_Arch.PNG" />

## Inspiration

&nbsp;&nbsp;&nbsp;&nbsp; Before I was invested in web3.0, blockchain, decentralized systems, etc., as mentioned in [Catena](https://github.com/ReshiAdavan/Catena), I was interested in Crypto because of a Senior Engineer I worked with over Winter 2023. Looking into Crypto and exchanges, I wanted to make one (at this point whenever I want to learn more of something I make a tool that focuses on it/revolves around it). Thats exactly what I did.

&nbsp;&nbsp;&nbsp;&nbsp; The choice of cryptocurrency was ethereum because of [Vitalik Buterin](https://www.linkedin.com/in/vitalik-buterin-267a7450/?originalSubdomain=ca), the infamous Waterloo student who founded [Ethereum](https://ethereum.org/en/). At the very least I am inspired by him and motivated to be like him.

## Architecture

**Relevant Folder Structure**:

```
.
├── gateway
│   ├── ...
│   ├── Dockerfile
│   ├── handlers.go
│   ├── main.go
│   └── utils.go
├── match
│   ├── ...
│   ├── fast-skiplist/
│   ├── asset.go
│   ├── Dockerfile
│   ├── history.go
│   ├── main.go
│   ├── market.go
│   └── message.go
├── memorystore
│   ├── ...
│   ├── Dockerfile
│   └── main.go
├── utils
│   ├── ...
│   ├── glog/
│   ├── recover.go
│   └── time.go
├── wallet
│   ├── ...
│   ├── bitcoin/
│   └── ethereum/
├── ws
│   ├── ...
│   ├── Dockerfile
│   ├── main.go
│   ├── message.go
│   └── time.go
└── REAMDE.md
```

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

#### Gateway Service

`Dockerfile`:

- Includes set up the Docker environment for the gateway service.
  Includes instructions for building the Docker image, specifying the base image, setting environment variables, copying files into the container, and defining the commands to run.

`Handler`:

- Defines various HTTP handlers for the gateway service.
- Contains functions to handle different types of HTTP requests (GET, POST, PUT, DELETE) and their corresponding business logic.
- Key aspects include importing packages for HTTP handling, context management, and possibly parsing and converting data.

`Main`:

- The main entry point of the gateway.
- Includes initialization of the service, sets up routing and middleware, and starts the HTTP server.
- Involves importing packages for context handling, JSON encoding, error handling, and other utilities.

`Utils`:

- Contains utility functions and helpers used across the gateway service.
  - Logging helpers, error handling functions, data processing utilities, etc.
- Includes imports for byte handling, HTTP client functionality, and JSON or protobuf handling for data serialization.

#### Match Service

&nbsp;&nbsp;&nbsp;&nbsp; A trading/transactional system, involved in handling and processing orders in a financial or trading manner. It includes mechanisms for managing orders, executing transactions, and maintaining historical data.

Functionality:

- Order Management: Handling buy/sell orders, maintaining order books.
- Trade Execution: Matching buy and sell orders, calculating trade values and fees.
- Asset Management: Tracking and updating asset balances for users.
- Market Handling: Managing market data, possibly involving different trading pairs or assets.
- Historical Data Management: Recording and storing transaction history for auditing or analytical purposes.
- Messaging and Notifications: Sending updates or notifications, likely to external systems or users, regarding order status, market changes, etc.
- Utilities and Helpers: Functions for common tasks like error handling, logging, and data conversions.

`Asset`:

- Manages asset-related functionalities in the trading/matching service.
  - Handling asset data structures.
  - Performing operations related to assets, such as balance checks and updates.

`Dockerfile`:

- Builds a containerized environment for the service.

`History`:

- Manages historical data for trades and transactions.
  - Storing and retrieving trade history.
  - Handling data structures related to historical records.

`Main`:

- Serves as the entry point for the service.
  - Initializing the service components.
  - Setting up configurations, connections, and starting the server.

`Market`:

- Handles market-related data and operations within the service.
  - Managing market information, such as trading pairs and market states.
  - Interfacing with market data for order processing.

`Message`:

- Manages messaging and notification functionalities.
  - Sending updates or alerts regarding trades, market changes, or order statuses.
  - Handling data structures for messaging.

#### Memory Store Service

&nbsp;&nbsp;&nbsp;&nbsp; Back-end service focused on managing in-memory data storage. Responsible for temporarily storing data, caching, or providing fast access to frequently used information. The service is containerized using Docker for ease of deployment and scalability.

&nbsp;&nbsp;&nbsp;&nbsp; The service is optimized for speed and efficiency, providing quick read and write capabilities.

`Dockerfile`:

- Defines a container for the `memorystore` service.
- Includes instructions for setting up the environment, installing dependencies, and specifying how the service should be run within the container.

`Main`:

- Entry point of the `memorystore` service.
- Contains initialization procedure to start the service.
- Sets up necessary configurations, initializes connections, and starts the server to listen for incoming requests.
- Handles routing for various endpoints and integrate middleware for various functionalities like logging, error handling, or request parsing.

#### Wallet Service

- Manages cryptocurrency transactions, specifically Ethereum.
- Includes functionality for interacting with the Ethereum blockchain, handling HD wallets, managing transaction events, and performing timed transfers.

`Main`:

- **Purpose**: Serves as the entry point for the Wallet service.
- **Functionality**: Contains the `main()` function which appears to be a placeholder (prints "vim-go"). This indicates the actual service logic might be defined in other files or this file is a template for further development.

`Dockerfile`:

- Defines container for the Wallet service.
  - Uses `gcr.io/cloud-builders/go:latest` as the base image.
  - Installs dependencies like `librdkafka-dev`.
  - Sets the working directory and copies the source code.
  - Builds the Go applications located in different subdirectories (`gateway`, `match`, `memorystore`, `sql`, `wallet`, `ethereum`, `ws`).
  - Uses `alpine:3.8` as the final image for a smaller footprint.
  - Copies the built application (`ethereum`) into the final image and sets it as the entry point.

`Events`:

- Manages events related to Ethereum transactions.
  - Connects to an Ethereum client, subscribes to head events (new blocks), and processes them.
  - Manages a map (`pendingTxs`) for tracking pending transactions.
  - Interacts with Apache Kafka for messaging and event handling.
  - Calculates transaction confirmations and updates the status of transactions accordingly.

`Hdwallet`:

- Handles operations related to Hierarchical Deterministic (HD) wallets for Ethereum.
  - Manage creation and management of HD wallets.
  - Involve key generation, address derivation, and transaction signing.

`Messaging`:

- Manages messaging aspects of the Wallet service.
  - Handles internal and external communication for the service.
  - Integration with messaging system aka Kafka.

`Timer`:

- Manages timed operations or scheduled tasks.
  - Periodically executes a task to transfer funds from multiple addresses to a specified address.
  - Checks balances of addresses and initiates transactions if the balance is above a threshold.
  - Interacts with Ethereum client for transaction creation and submission.

#### WS Service

&nbsp;&nbsp;&nbsp;&nbsp; Service designed for handling and processing messages, with a particular emphasis on time-related functionalities.

`Dockerfile`:

- Sets up the Docker environment for the WS service.
  - Details the base image, environment configurations, dependencies, and execution instructions for the service.

`Main`:

- Acts as the entry point of the WS service.
  - Initializes the service, configures routing and middleware, starts the server, and may include configuration loading and logging.

`Message`:

- Manages the structure and processing of messages.
  - Includes definitions of message formats (structs or interfaces), along with functions for processing, validating, and possibly serializing/deserializing messages.

`Time`:

- Provides utilities related to time manipulation or scheduling.
- Contains functions for time formatting, parsing, and managing time-based tasks or events.
