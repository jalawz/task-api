# Task API (Go + Gin + GORM)

A robust REST API built with Go, featuring a clean architecture with separate layers for Controllers and Services.

## Tech Stack

- **Go:** Programming language
- **Gin:** HTTP web framework
- **GORM:** Object Relational Mapper for Go
- **SQLite:** Lightweight, disk-based database

## Architecture

The project follows a modular structure to ensure maintainability and testability:

- **Models:** Data structures and database schema definitions.
- **Services:** Business logic layer. Handles database operations and provides an interface for the controllers.
- **Controllers:** HTTP layer. Handles request binding, calls the service layer, and returns standardized JSON responses.
- **Routes:** Centralized route management and dependency injection.
- **Database:** Connection setup and automatic migrations.

---

## Getting Started

### Prerequisites
- Go 1.25+

### Run Locally
```bash
go run main.go
```
The server starts on `http://localhost:8000`. A local SQLite file `tasks.db` will be created automatically.

---

## API Documentation

### 1) List All Tasks
- **Method:** `GET`
- **Path:** `/tasks`

#### Response (200 OK)
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

---

### 2) Create Task
- **Method:** `POST`
- **Path:** `/tasks`
- **Content-Type:** `application/json`

#### Request Body
```json
{
  "title": "Master Go Interfaces",
  "description": "Study how to use interfaces for DI",
  "status": false
}
```

#### Response (201 Created)
```json
{
  "id": 2,
  "title": "Master Go Interfaces",
  "description": "Study how to use interfaces for DI",
  "status": false,
  "created_at": "2026-03-07T18:00:00Z",
  "updated_at": "2026-03-07T18:00:00Z"
}
```

---

### 3) Get Task by ID
- **Method:** `GET`
- **Path:** `/tasks/:id`

#### Response (200 OK)
```json
{
  "id": 2,
  "title": "Master Go Interfaces",
  "description": "Study how to use interfaces for DI",
  "status": false,
  "created_at": "2026-03-07T18:00:00Z",
  "updated_at": "2026-03-07T18:00:00Z"
}
```

#### Response (404 Not Found)
```json
{
  "error": "Task with id 99 not found"
}
```

---

### 4) Update Task
- **Method:** `PUT`
- **Path:** `/tasks/:id`
- **Content-Type:** `application/json`

#### Request Body
> Note: Only fields provided will be updated. Since `status` is a pointer, `false` is correctly updated.
```json
{
  "title": "Advanced Go Patterns",
  "status": true
}
```

#### Response (200 OK)
```json
{
  "id": 2,
  "title": "Advanced Go Patterns",
  "description": "Study how to use interfaces for DI",
  "status": true,
  "created_at": "2026-03-07T18:00:00Z",
  "updated_at": "2026-03-07T18:10:00Z"
}
```

---

### 5) Delete Task
- **Method:** `DELETE`
- **Path:** `/tasks/:id`

#### Response (204 No Content)
*No body returned.*

#### Response (404 Not Found)
```json
{
  "error": "Task not found"
}
```

---

## Summary Table

| Method | Endpoint      | Description           | Status Codes |
|--------|---------------|-----------------------|--------------|
| GET    | `/tasks`      | List all tasks        | 200, 500     |
| POST   | `/tasks`      | Create a new task     | 201, 400, 500|
| GET    | `/tasks/:id`  | Get task details      | 200, 404, 500|
| PUT    | `/tasks/:id`  | Update an existing task| 200, 400, 404, 500|
| DELETE | `/tasks/:id`  | Delete a task         | 204, 404, 500|

---

## Postman Collection
A pre-configured Postman collection is available in the root directory: `Task_API.postman_collection.json`. 
Import it into Postman to start testing immediately.

## License
MIT
