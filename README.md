# Task API (Go + Gin + GORM)

A simple REST API built for learning Go backend fundamentals.

This project uses:
- [Gin](https://github.com/gin-gonic/gin) for HTTP routing
- [GORM](https://gorm.io/) as ORM
- [SQLite](https://www.sqlite.org/index.html) as local database

## Project Goal

The main goal is educational: understand how to organize a Go API with:
- route setup
- controllers
- models
- database initialization and migration
- basic CRUD-style endpoints

## Current Features

- Health check endpoint (`/ping`)
- Create a task
- List all tasks
- Get one task by ID
- Automatic database migration for `Task` model

## Tech Stack

- Go
- Gin
- GORM
- SQLite

## Project Structure

```text
.
├── controllers/
│   └── task_controller.go
├── database/
│   └── db.go
├── models/
│   └── task.go
├── main.go
├── go.mod
└── tasks.db
```

## Getting Started

### Prerequisites

- Go (version declared in `go.mod`)

### Run Locally

```bash
go run main.go
```

By default, the server starts on:

```text
http://localhost:8080
```

A local SQLite database file (`tasks.db`) is created automatically.

## API Documentation (Current)

Base URL:

```text
http://localhost:8080
```

### 1) Health Check

- **Method:** `GET`
- **Path:** `/ping`

#### Response

- **Status:** `200 OK`

```json
{
  "message": "pong"
}
```

---

### 2) Create Task

- **Method:** `POST`
- **Path:** `/tasks`
- **Content-Type:** `application/json`

#### Request Body

```json
{
  "title": "Study Go",
  "description": "Read about Gin and GORM",
  "status": false
}
```

#### Response (success)

- **Status:** `201 Created`

```json
{
  "id": 1,
  "title": "Study Go",
  "description": "Read about Gin and GORM",
  "status": false,
  "created_at": "2026-03-07T03:00:00Z",
  "updated_at": "2026-03-07T03:00:00Z"
}
```

#### Response (invalid JSON)

- **Status:** `400 Bad Request`

```json
{
  "error": "<validation or bind error>"
}
```

#### Response (database error)

- **Status:** `500 Internal Server Error`

```json
{
  "error": "Error creating task"
}
```

---

### 3) List Tasks

- **Method:** `GET`
- **Path:** `/tasks`

#### Response (success)

- **Status:** `200 OK`

```json
[
  {
    "id": 1,
    "title": "Study Go",
    "description": "Read about Gin and GORM",
    "status": false,
    "created_at": "2026-03-07T03:00:00Z",
    "updated_at": "2026-03-07T03:00:00Z"
  }
]
```

#### Response (database error)

- **Status:** `500 Internal Server Error`

```json
{
  "error": "Error fetching tasks"
}
```

---

### 4) Get Task by ID

- **Method:** `GET`
- **Path:** `/tasks/:id`

#### Example

```text
GET /tasks/1
```

#### Response (success)

- **Status:** `200 OK`

```json
{
  "id": 1,
  "title": "Study Go",
  "description": "Read about Gin and GORM",
  "status": false,
  "created_at": "2026-03-07T03:00:00Z",
  "updated_at": "2026-03-07T03:00:00Z"
}
```

#### Response (not found)

- **Status:** `404 Not Found`

```json
{
  "error": "Task not found"
}
```

## Example cURL Commands

```bash
# Health check
curl -X GET http://localhost:8080/ping

# Create task
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Study Go",
    "description": "Read about Gin and GORM",
    "status": false
  }'

# List tasks
curl -X GET http://localhost:8080/tasks

# Get task by ID
curl -X GET http://localhost:8080/tasks/1
```

## Notes

- This is an early-stage learning project.
- There are no automated tests yet.
- Input validation rules are still minimal and can be expanded.
