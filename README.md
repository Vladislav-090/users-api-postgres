# Users API (In-Memory)

A simple REST API written in Go for managing users in memory.

This project was created as part of my Golang learning journey to practice HTTP handlers, JSON processing, validation, routing, and project structure.

## Features

* Add user
* Get all users
* Get user by name
* Delete user
* Clear all users
* Count users
* JSON success responses
* JSON error responses
* HTTP status codes
* Duplicate user validation

## Technologies

* Go
* net/http
* encoding/json

## Project Structure

```text
users-api-memory/
│
├── go.mod
├── main.go
├── README.md
│
└── internal
    ├── handlers
    │   └── users_handler.go
    │
    ├── models
    │   └── user.go
    │
    └── response
        └── response.go
```

## API Endpoints

### Add User

```http
POST /addUser
```

Request:

```json
{
  "name": "Max",
  "age": 10
}
```

Success Response

```json
{
  "message": "User added successfully!",
  "user": {
    "name": "Max",
    "age": 10
  }
}
```

---

### Get All Users

```http
GET /getUsers
```

---

### Get User

```http
GET /getUser?name=Max
```

---

### Delete User

```http
DELETE /deleteUser?name=Max
```

---

### Clear Users

```http
DELETE /clearUsers
```

---

### Count Users

```http
GET /getCount
```

## Validation

The API validates:

* HTTP method
* JSON body
* Empty name
* Positive age
* Duplicate users

## Error Response

```json
{
  "error": "User already exist!"
}
```

## Run

```bash
go run .
```

## Future Improvements

* PostgreSQL
* Docker
* Environment variables
* Logging
* Middleware
* Unit tests
* Authentication
