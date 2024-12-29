# User Service Microservice

This is a microservice for managing user data, built using Go with clean architecture principles. The service provides basic CRUD operations for user management with PostgreSQL as the database.

## Architecture

This project follows clean architecture principles with the following layers:

- Domain: Contains the business entities and repository interfaces
- Repository: Implements data persistence logic using PostgreSQL
- Usecase: Contains business logic
- Delivery: Handles HTTP requests and responses using Gin framework

## Technologies Used

- Go 1.21
- PostgreSQL
- GORM (Go Object Relational Mapper)
- Gin Web Framework
- UUID for unique identifiers
- Godotenv for environment configuration

## Project Structure

```
user-service/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── domain/                 # Business entities and interfaces
│   │   └── user.go
│   ├── repository/             # Data persistence implementations
│   │   └── user_repository.go
│   ├── usecase/               # Business logic implementations
│   │   └── user_usecase.go
│   └── delivery/              # HTTP handlers and routes
│       └── http/
│           ├── handler/
│           │   └── user_handler.go
│           └── routes/
│               └── routes.go
├── pkg/
│   ├── config/               # Configuration helpers
│   │   └── config.go
│   └── database/            # Database connection setup
│       └── postgres.go
├── .env                     # Environment variables
├── go.mod
└── go.sum
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Git

## Getting Started

1. Clone the repository:

```bash
git clone <your-repository-url>
cd user-service
```

2. Install dependencies:

```bash
go mod download
```

3. Create `.env` file in the root directory:

```env
DB_HOST=localhost
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=user_service
DB_PORT=5432
APP_PORT=8080
```

4. Create database:

```sql
CREATE DATABASE user_service;
```

5. Run the application:

```bash
go run cmd/main.go
```

The service will start on `http://localhost:8080`

## API Endpoints

### Create User

- Method: POST
- URL: `/api/users`
- Body:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "secret123"
}
```

### Get User by ID

- Method: GET
- URL: `/api/users/id/:id`

### Get User by Email

- Method: GET
- URL: `/api/users/email/:email`

### Update User

- Method: PUT
- URL: `/api/users/:id`
- Body:

```json
{
  "name": "John Updated",
  "email": "john@example.com",
  "password": "newpassword123"
}
```

### Delete User

- Method: DELETE
- URL: `/api/users/:id`

## Error Handling

The service returns appropriate HTTP status codes:

- 200: Success
- 201: Created
- 400: Bad Request
- 404: Not Found
- 500: Internal Server Error

## Testing

You can test the API endpoints using the provided Postman collection in the repository.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details
