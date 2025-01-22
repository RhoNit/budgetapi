# BudgetAPI

## Description
BudgetAPI is a simple API that helps manage and track budgets for personal finance. It allows users to create, view, update, and delete budget entries and track expenses. Built using Echo/Golang, Docker and Postgres.

## Features
- **Create Budget**: Allows users to create a new budget entry.
- **View Budget**: View all budget entries for a user.
- **Update Budget**: Update details of an existing budget entry.
- **Delete Budget**: Remove a budget entry.
- **Expense Tracking**: Track expenses within each budget entry.
- **Authentication**: Secured API with authentication using JWT-Auth.

## Project Structure
Here’s an overview of the project structure:
```plaintext
budgetapi/
  ├── cmd/api
  │   ├── handlers
  │   │   ├── app_handler.go
  │   │   ├── auth_handler.go
  │   │   ├── category_handler.go
  │   │   ├── handler.go
  │   │   └── validate_request_handler.go
  │   ├── middlewares
  │   │   └── auth_middleware.go
  │   ├── requests
  │   │   ├── category_request.go
  │   │   ├── param_request.go
  │   │   └── user_request.go
  │   ├── routes
  │   │   └── endpoints.go
  │   ├── services
  │   │   ├── category_service.go
  │   │   └── user_service.go
  │   ├── validation
  │   │   └── validation.go
  │   └── main.go
  |
  ├── common
  │   ├── custom_errors
  │   │   └── not_found_error.go
  │   ├── api_response.go
  │   ├── db_connection.go
  │   ├── jwt.go
  │   ├── pagination.go
  │   └── password.go
  |
  ├── internal
  │   ├── mailer
  │   │   ├── templates
  │   │   │   ├── hello.html
  │   │   │   └── welcome.html
  │   │   └── mailer.go
  │   ├── migration
  │   │   ├── seeders
  │   │   │   └── category_seeders.go
  │   │   └── migrate_up.go
  │   └── models
  │       ├── base_model.go
  │       ├── category.go
  │       └── user.go
  |
  ├── .env
  ├── .gitignore
  ├── README.md
  ├── docker-compose.yml
  ├── go.mod
  └── go.sum
