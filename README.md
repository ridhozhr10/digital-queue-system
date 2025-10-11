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

To run the Swagger UI for the authentication service, run the following command:

```bash
go run cmd/auth-service/main.go swagger
```

This will start a server on port 8081 (or the port specified in the `SWAGGER_PORT` environment variable) and serve the Swagger UI at [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html).

The swagger documentation is generated using `swag init` and the generated files are located in the `pkg/swagger` directory.

### Frontend

To run the frontend, navigate to the `web` directory and run the following commands:

```bash
npm install
npm start
```
