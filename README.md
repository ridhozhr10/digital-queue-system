# digital-queue-system

This is a monorepo for a digital queue system with a Go backend and a React frontend.

## Documentation

For detailed information on architecture, API reference, and development setup, please see the [**full documentation in the /docs directory**](./docs/index.md).

## Quick Start

This project uses Docker and Make to simplify the development setup.

### Prerequisites

- Go
- Node.js
- Docker & Docker Compose
- Make

(See the [development setup guide](./docs/development-setup.md) for installation instructions.)

### Running the System

Use the following commands from the project root to manage the development environment:

- `make compose-up`: Start all services.
- `make compose-down`: Stop all services.
- `make compose-logs`: View service logs.
- `make compose-ps`: List running services.
