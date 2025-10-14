# Backend API Reference

This document provides a detailed reference for the backend Auth Service API.

The API is served by the Go `http-service` application and is accessible through the Kong API Gateway.

## Endpoints

### Health Check

#### `GET /ping`

A simple health check endpoint to verify that the service is running.

-   **Success Response (200 OK):**
    ```json
    {
      "message": "pong"
    }
    ```

---

### Authentication

#### `POST /auth/login`

Authenticates a user and returns a JWT if successful.

-   **Request Body:** `application/json`
    ```json
    {
      "user_identity": "john.doe",
      "password": "password123"
    }
    ```
    -   `user_identity` (string, required): The username or email of the user.
    -   `password` (string, required): The user's password.

-   **Success Response (200 OK):**
    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```

-   **Error Response (400 Bad Request):**
    ```json
    {
      "message": "Invalid email or password"
    }
    ```

#### `POST /auth/register`

Registers a new user in the system.

-   **Request Body:** `application/json`
    ```json
    {
      "username": "John Doe",
      "email": "john.doe@example.com",
      "password": "securepassword123"
    }
    ```
    -   `username` (string, required): The user's full name.
    -   `email` (string, required): The user's email address. Must be unique.
    -   `password` (string, required): The user's password.

-   **Success Response (201 Created):**
    ```json
    {
      "message": "User registered successfully"
    }
    ```

-   **Error Response (400 Bad Request):**
    ```json
    {
      "message": "User with this email already exists"
    }
    ```

## Data Models

The API uses the following data models, which correspond to the `user` model in the Go `internal/model/user.go` file.

### User

| Field      | Type   | Description                               |
| :--------- | :----- | :---------------------------------------- |
| `id`       | int    | The unique identifier for the user.       |
| `username` | string | The name of the user.                     |
| `email`    | string | The user's email address (must be unique).|
| `password` | string | The hashed password of the user.          |
