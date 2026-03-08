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
- **Services:** Business logic layer. Handles database operations using GORM and provides an interface for the controllers.
- **Controllers:** HTTP layer. Handles request binding, calls the service layer, and returns standardized JSON responses.
- **Routes:** Centralized route management and dependency injection.
- **Database:** Connection setup and automatic migrations.

## Current Features

- **Create Task:** Validates input and persists new tasks.
- **List Tasks:** Retrieves all tasks from the database.
- **Get Task by ID:** Finds a specific task or returns a 404 error if not found.
- **Update Task:** Supports partial updates (including boolean status) and returns the updated record.
- **Delete Task:** Removes a task by ID with proper error handling for non-existent records.
- **Auto-Migration:** Automatically updates the database schema based on the `Task` model.

## Project Structure

```text
.
├── controllers/    # Request handling & HTTP responses
├── services/       # Business logic & Database interaction
├── database/       # DB connection & initialization
├── models/         # GORM models
├── routes/         # API route definitions & DI
├── main.go         # Application entry point
├── go.mod          # Dependencies
└── tasks.db        # SQLite database file
```

## Getting Started

### Prerequisites

- Go 1.25+

### Run Locally

```bash
go run main.go
```

The server starts on `http://localhost:8000`.

## API Documentation

### Base URL: `http://localhost:8000`

| Method | Endpoint      | Description           |
|--------|---------------|-----------------------|
| GET    | `/tasks`      | List all tasks        |
| POST   | `/tasks`      | Create a new task     |
| GET    | `/tasks/:id`  | Get task details      |
| PUT    | `/tasks/:id`  | Update an existing task|
| DELETE | `/tasks/:id`  | Delete a task         |

---

### Sample Request (Create/Update)

```json
{
  "title": "Master Go Interfaces",
  "description": "Study how to use interfaces for DI",
  "status": false
}
```

## License
MIT
