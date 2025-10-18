# Development Setup

This guide explains how to set up and run the Digital Queue System for local development.

## Prerequisites

Before you begin, ensure you have the following installed:

- Go (version 1.25.1 or higher)
- Node.js (for the frontend)
- **Docker & Docker Compose:** Required to run the services in containers.
- **Make:** Used to simplify running development commands.
  - **On Linux/macOS:** `make` is typically pre-installed.
  - **On Windows:** Install via [Chocolatey](https://chocolatey.org/) with `choco install make` or use the version included with Git for Windows / WSL.

## Running with Docker (Recommended)

The easiest way to get all services (backend, database, API gateway) running is to use the provided `Makefile`.

Run the following commands from the project root:

- `make compose-up`: To start all services in detached mode.
- `make compose-down`: To stop all services.
- `make compose-logs`: To view the logs from all running services.
- `make compose-ps`: To list the running Docker containers.
- `make compose-rebuild`: To rebuild and start all services.
- `make kong-sync`: To sync the declarative Kong configuration (`configs/kong/declarative.yml`) with the database.

After running `make compose-up`, the following services will be available:

- **Kong API Gateway:** `http://localhost:8000`
- **Kong Admin GUI:** `http://localhost:8002`
- **Application Database (PostgreSQL):** Accessible on port `5432`.

## Database Migrations

This project uses [dbmate](https://github.com/amacneil/dbmate) to manage database schema changes. The migration files are located in the `/db/migrations` directory.

Commands are run using `npx` from the project root, which will automatically use the `DATABASE_URL` defined in your `.env` file.

### Create a new migration

To create a new SQL migration file, run the following command, replacing `<migration_name>` with a descriptive name for your migration (e.g., `add_users_table`).

```bash
npx dbmate new <migration_name>
```

### Apply migrations

To apply all pending migrations to your database, run:

```bash
npx dbmate up
```

### Rollback migrations

To roll back the most recent migration, run:

```bash
npx dbmate down
```

## Running the Backend Manually

If you prefer to run the Go backend service outside of Docker, you can use the following commands.

**Note:** You will still need to run the databases via Docker (`make compose-up`) for the backend to connect to.

### Start the API Server

To run the authentication service:

```bash
go run cmd/http-service/main.go auth
```

The server will start on port `8080` by default.

### Run Swagger UI

If you run the API Server in development mode, you can view the API documentation locally via Swagger UI on `http://localhost:8080/swagger-ui`.

## Running Tests

To run all unit tests for the backend:

```bash
go test ./cmd/auth-service/app
go test ./internal/...
```
