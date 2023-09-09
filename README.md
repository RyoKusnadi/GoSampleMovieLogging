# Scalable Go Movie [WIP]

## Requirements:

- Golang 1.21

## Description:

A robust and scalable movie data processing system that leverages concurrent programming, GraphQL for CRUD operations, Kafka for asynchronous communication, and the Onion architecture pattern for modularity. Incorporate Docker for containerization and Kubernetes for orchestration, ensuring seamless deployment and scalability.

## Technical Components:

1. **Concurrency:** Use Goroutines and Channels for concurrent movie data fetching, processing, and CRUD operations.
2. **TMDb API:** Fetch movie data from the TMDb API, including details, cast info, and ratings.
3. **GraphQL:** Implement a GraphQL server for performing CRUD operations on movie data.
4. **Kafka:** Utilize Kafka for asynchronous communication between components and services.
5. **Onion Architecture:** Organize your application into core domain logic, GraphQL resolvers, services, and Kafka consumers.
6. **Docker:** Containerize application components, Kafka consumers, and the database using Docker containers.
7. **Kubernetes:** Set up a Kubernetes cluster to manage and orchestrate containerized components.

## Project Features:

1. **Containerization:** Dockerize Golang app, GraphQL server, Kafka consumers, and the database into separate containers.
2. **Kubernetes Deployment:** Deploy Docker containers using Kubernetes, ensuring scaling, load balancing, and updates.
3. **Service Discovery:** Utilize Kubernetes Services for seamless communication between components.
4. **Horizontal Pod Autoscaling:** Implement Kubernetes' Horizontal Pod Autoscaling to optimize resource utilization.
5. **Concurrent Data Processing:** Fetch movie data concurrently from TMDb API and store it in the database.
6. **GraphQL API:** Create GraphQL types, queries, mutations, and resolvers for CRUD operations.
7. **Kafka Messaging:** Publish Kafka messages for new movie data and consume them for processing.
8. **Onion Architecture:** Separate application layers for maintainable and modular code.
