# digital-queue-system

This is a monorepo for a digital queue system with a Go backend and a React frontend.

## Development

### Prerequisites

- Go
- Node.js

### Backend

To run the authentication service, navigate to the root of the project and run the following command:

```bash
go run cmd/auth-service/main.go serve
```

The server will start on port `8080` by default. You can specify a different port using the `-p` flag or the `PORT` environment variable:

```bash
go run cmd/auth-service/main.go serve -p 9000
# or
PORT=9000 go run cmd/auth-service/main.go serve
```

To run the Swagger UI for the authentication service, run the following command:

```bash
go run cmd/auth-service/main.go swagger
```

This will start a server on port `8081` by default and serve the Swagger UI at `http://localhost:8081/swagger-ui`. You can specify a different port using the `-p` flag or the `SWAGGER_PORT` environment variable:

```bash
go run cmd/auth-service/main.go swagger -p 9001
# or
SWAGGER_PORT=9001 go run cmd/auth-service/main.go swagger
```

The OpenAPI specification for the auth service is located at `api/auth-service.yaml` and is served by the `swagger` command.

### CORS

CORS (Cross-Origin Resource Sharing) is enabled for the authentication service, allowing requests from any origin. This can be configured more restrictively in a production environment.

### Testing

To run all unit tests for the `auth-service` application:

```bash
go test ./cmd/auth-service/app
```

To run all unit tests for the `internal` packages:

```bash
go test ./internal/...
```

### Frontend

To run the frontend, navigate to the `web` directory and run the following commands:

```bash
npm install
npm start
```