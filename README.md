# Keep It Secret API

## Overview
This project is a Go-based API backend following best practices for maintainability, scalability, and clarity. The initial implementation provides a `/health` endpoint for health checks.

## Architecture
- **Entry Point:** `main.go` contains the application entry point and route setup.
- **Routing:** Uses Go's standard `net/http` package with `http.ServeMux` for route management.
- **Handlers:** Each route is handled by a dedicated function for separation of concerns.
- **Response Format:** JSON responses for API consistency.
- **Logging:** Basic logging for server startup and errors.
- **Extensibility:** Structure is ready for modular expansion (e.g., move handlers to separate files/packages as the API grows).

## Endpoints
### `/health`
- **Method:** GET
- **Response:**
	```json
	{
		"message": "The API is running.",
		"number": 1234
	}
	```
	- `message`: Status message
	- `number`: Random number (changes on each request)

## Running Locally
1. Ensure you have Go installed (>=1.18).
2. Run the server:
	 ```sh
	 go run main.go
	 ```
3. Access the health endpoint:
	 [http://localhost:8080/health](http://localhost:8080/health)

## Next Steps
- Add more routes and business logic in separate files/packages.
- Implement configuration management, middleware, and error handling as needed.
- Add unit and integration tests.

## Best Practices Followed
- Clear separation of concerns
- Simple, idiomatic Go code
- Ready for modularization
- JSON API responses
- Logging for observability
# keep-it-secret
This is a repository to store secrets in a local way.
