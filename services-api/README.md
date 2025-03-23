# Services API

A read-only API for a service catalog dashboard widget, built in Go.

## Features

- **GET /services**: List services with filtering (by name/description), sorting (by any field), and pagination (limit/offset).
- **GET /services/:id**: Fetch a specific service by its ID.
- **GET /services/:id/versions**: Retrieve all versions of a specific service.

## Design Considerations

- **Persistence**: Uses SQLite with GORM for lightweight, file-based storage suitable for a small, read-only API.
- **Framework**: Employs Gin for routing and request handling due to its performance and simplicity.
- **IDs**: Utilizes UUIDs for unique identifiers, ensuring flexibility.
- **Error Handling**: Returns appropriate HTTP status codes (e.g., 404 for not found, 500 for server errors).

## Assumptions

- Filtering applies to `name` and `description` fields via a `search` query parameter.
- Sorting is flexible but assumes valid field names (e.g., `name`, `description`).
- Pagination uses `limit` and `offset` parameters; defaults are 10 and 0, respectively.
- The API is read-only, so no CRUD operations beyond reading are implemented.

## Trade-offs

- **SQLite**: Chosen over more robust databases (e.g., PostgreSQL) for simplicity, though it may not scale as well for concurrent writes (not a concern here since the API is read-only).
- **Minimal Validation**: Basic checks for pagination parameters; additional validation could be added for robustness.
- **No Auth**: Authentication/authorization omitted due to time constraints but could be added with API keys or JWT.

## Setup and Running

1. Ensure Go is installed.
2. Clone the repository:
   ```sh
   git clone <repository-url>
   cd services-api