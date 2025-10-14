# Architecture Overview

This document provides a high-level overview of the system architecture for the Digital Queue System.

## System Components

The system is composed of the following main services, all running in Docker containers:

![System Architecture Diagram](https://i.imgur.com/A8LIqS3.png)
*(This is a simplified diagram showing the main data flow)*

1.  **Kong API Gateway (`kong-gw`)**
    -   **Purpose:** Acts as the single entry point for all incoming API requests. It is responsible for routing, authentication, and logging.
    -   **Database:** Uses its own dedicated PostgreSQL database (`kong-db`) to store its configuration (routes, services, consumers, etc.).
    -   **Ports:**
        -   `8000`: Proxy for API traffic.
        -   `8001`: Admin API for configuration.
        -   `8002`: Admin GUI.

2.  **Auth Service (`http-service`)**
    -   **Purpose:** A Go-based microservice that handles user registration and login. (Currently, this service is not defined in the Docker Compose files but is intended to be containerized).
    -   **Database:** Connects to the main application database (`postgres_db`).
    -   **API:** Exposes endpoints like `/auth/register` and `/auth/login`.

3.  **Application Database (`postgres_db`)**
    -   **Purpose:** A PostgreSQL database that stores the primary application data, such as user information.
    -   **Container Name:** `rbac_postgres_db`
    -   **Port:** `5432`

4.  **Kong Database (`kong-db`)**
    -   **Purpose:** A dedicated PostgreSQL database for the Kong API Gateway.
    -   **Container Name:** `kong-db`

## Network

All services are connected via a Docker bridge network named `app-net`. This allows them to communicate with each other using their service names as hostnames (e.g., `kong-gw` can connect to `postgres_db`).

## Request Flow

1.  A client sends a request (e.g., a login request to `/auth/login`) to the Kong API Gateway on port `8000`.
2.  Kong, based on its configuration, forwards the request to the appropriate upstream service (the `auth-service`).
3.  The `auth-service` processes the request, interacting with the `postgres_db` to validate credentials or create a new user.
4.  The `auth-service` returns a response to Kong.
5.  Kong forwards the response back to the client.
