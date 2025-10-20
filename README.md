# Employee CRUD API

NOTE: Project for job test

A CRUD api for managing employee built using Go, implementing standard Go layout and clean architecture


## 📋 Prerequisites

Before you begin, ensure you have installed:

- Go 1.21 or higher
- PostgreSQL 15

## 🎯 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/MaulanaAhmadSulami/juke_test.git
cd juke_test
```

### 2. Environment Configuration

Create a `.env` file per .env.example

### 3. Run with Docker

```bash
# Start all services (PostgreSQL + API)
docker-compose up --build

# Stop services
docker-compose down
```

## 📚 API Documentation

Once the application is running, access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

### 4. Seed Sample Data (Optional)

After starting the application, run these:

```bash
# Create Employee 1
curl -X POST http://localhost:8080/api/v1/employees \
  -H "Content-Type: application/json" 
  -d '{
    "name": "John Doe",
    "email": "john.doe@company.com",
    "position": "Software Engineer",
    "salary": 75000
  }'

# Create Employee 2
curl -X POST http://localhost:8080/api/v1/employees \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane.smith@company.com",
    "position": "Product Manager",
    "salary": 85000
  }'

# Create Employee 3
curl -X POST http://localhost:8080/api/v1/employees \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bob Johnson",
    "email": "bob.johnson@company.com",
    "position": "DevOps Engineer",
    "salary": 80000
  }'
```

### API Endpoints

| Method | Endpoint                        | Description          |
|--------|---------------------------------|----------------------|
| GET    | `/api/v1/employees`             | Get all employees    |
| GET    | `/api/v1/employees/{id}`        | Get employee by ID   |
| POST   | `/api/v1/employees`             | Create new employee  |
| PUT    | `/api/v1/employees/{id}`        | Update employee      |
| DELETE | `/api/v1/employees/{id}`        | Delete employee      |
| GET    | `/health`                       | Health check         |


## 🏗️ Project Structure

```
project_hometest/
├── cmd/
│   ├── app/
│   │   └── main.go                 # Application entry point
│   └── migrate/
│       └── migrations/             # Database migrations
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── db/
│   │   └── db.go                  # Database connection
│   ├── entities/
│   │   └── employees/
│   │       └── employee.go        # Employee model
│   ├── repository/
│   │   └── postgres/
│   │       ├── employee/
│   │       │   └── employee.go    # Data access layer
│   │       └── repository.go      # Repository interfaces
│   ├── service/
│   │   ├── employee/
│   │   │   └── employee.go        # Business logic
│   │   └── service.go             # Service interfaces
│   └── server/
│       └── http/
│           ├── handler/
│           │   └── employee/
│           │       ├── handler.go # HTTP handlers
│           │       └── route.go   # Route definitions
│           └── protocol/
│               └── status.go      # Response utilities
├── docs/                          # Swagger documentation (auto-generated)
├── docker-compose.yml             # Docker orchestration
├── Dockerfile                     # Docker image definition
├── go.mod                         # Go dependencies
└── README.md                      # This file
```

## 🗄️ Database Schema

### Employee Table

```sql
CREATE TABLE employees (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    position VARCHAR(255) NOT NULL,
    salary DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
```

## 🔧 Development

### Run Tests

```bash
go test ./...
```

## 🐛 Troubleshooting

### Port Already in Use

```bash
# Check what's using the port
netstat -ano | findstr :8080

# Change port in docker-compose.yml or .env
```

### Database Connection Issues

```bash
# Check if PostgreSQL container is running
docker ps

# View PostgreSQL logs
docker logs employee_db

# Restart containers
docker-compose restart
```

### Migration Not Running

```bash
# Manually run migration
docker exec -i employee_db psql -U postgres -d employee_db < cmd/migrate/migrations/000001_create_employee_table_up.sql
```

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.
