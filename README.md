# Incubyte Salary Management API

A robust HTTP API built as part of the Incubyte engineering hiring assessment.

This application manages employee data, calculates salaries based on country-specific tax rules, and provides salary-related metrics. The project prioritizes software craftsmanship, leveraging a clean layered architecture, Test-Driven Development (TDD), and production-ready code standards.

> Note: This is an API-only application and does not include a web UI.

---

## Table of Contents

- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
- [Salary Rules](#salary-rules)
- [Development Approach](#development-approach)
- [AI Usage Disclosure](#ai-usage-disclosure)
- [How to Run](#how-to-run)

---

## Tech Stack

- Language: Go (Golang)
- HTTP Server: net/http
- Database: SQLite
- Testing: Standard Go testing package (testing)
- Architecture: Layered (Domain → Repository → Service → Handler)

---

## Project Structure

The project follows a standard Go layout with a clean separation of concerns:

```text
├── cmd/
│   └── server/
│       └── main.go            # Application entry point
├── internal/
│   ├── db/
│   │   └── db.go              # SQLite connection & setup helpers
│   ├── employee/
│   │   ├── employee.go        # Domain model & validation logic
│   │   ├── handler.go         # HTTP request handlers
│   │   ├── repository.go      # SQLite persistence logic
│   │   └── service.go         # Business logic orchestration
│   ├── metrics/
│   │   ├── handler.go         # Metrics HTTP handlers
│   │   └── metrics.go         # Salary aggregation logic
│   └── salary/
│       └── calculator.go      # Salary calculation logic
├── go.mod                     # Go module dependencies
└── README.md                  # This file
```

---

## API Endpoints

### Employees

#### 1. Create Employee

- Creates a new employee record in the system.

  ##### Endpoint: POST /employee

```bash
Request Body:

{
"full_name": "Amish Jha",
"job_title": "Engineer",
"country": "India",
"salary": 1000
}
```

```bash
Response: 201 Created

{
"id": 1
}
```

#### 2. Get Employee by ID

- Retrieves details of a specific employee.

#### Endpoint: GET /employees/{id}

```bash
Response: 200 OK

{
"id": 1,
"full_name": "Amish Jha",
"job_title": "Engineer",
"country": "India",
"salary": 1000
}
```

#### 3. Get Salary Calculation

- Calculates the gross, deduction, and net salary for a specific employee based on their country.

#### Endpoint: GET /employees/{id}/salary

```bash
Response: 200 OK

{
"gross": 1000,
"deduction": 100,
"net": 900
}
```

### Metrics

#### 4. Metrics by Country

- Retrieves salary statistics (min, max, average) for a specific country.

#### Endpoint: GET /metrics/country/{country}

```bash
Response: 200 OK

{
"min": 1000,
"max": 3000,
"avg": 2000
}
```

#### 5. Metrics by Job Title

- Retrieves the average salary for a specific job title across the company.

##### Endpoint: GET /metrics/job-title/{title}

```bash
Response: 200 OK

{
"avg": 5000
}
```

---

## Salary Rules

Deduction logic is applied automatically based on the employee's country of residence:

| Country       | Deduction Rule      |
| :------------ | :------------------ |
| India         | 10% of gross salary |
| United States | 12% of gross salary |
| All Others    | No deductions (0%)  |

---

## Development Approach

### Test-Driven Development (TDD)

This project was developed using a strict TDD workflow to ensure code quality and maintainability:

1. Red: Write a failing test case.
2. Green: Implement the minimal code required to pass the test.
3. Refactor: Improve the code structure and clarity without changing behavior.

The commit history reflects this process through small, incremental commits.

### Design Principles

- Thin HTTP Handlers: Handlers strictly parse requests and return responses; they contain no business logic.
- Repository Pattern: Abstracts database access, making it easy to swap storage implementations later.
- Service Layer: Orchestrates business rules and validation.
- Pure Domain Logic: Salary calculations are isolated in pure functions to ensure testability.
- Explicit Routing: Uses the standard net/http library without external router dependencies.
- Precision: Salary values are rounded to avoid common floating-point arithmetic errors.

---

## AI Usage Disclosure

This project utilized AI-assisted tooling (such as GitHub Copilot) strictly as a development accelerator, similar to an advanced IDE autocomplete feature.

### What AI Assisted With:

- Syntax suggestions and code completion for boilerplate Go code.
- Identifying standard library idioms and refactoring opportunities.
- Improving code readability and documentation generation.

### What Was Built Manually:

- Core Architecture: All system design and structural decisions were made independently.
- Business Logic: Salary rules, validation logic, and data flows were authored by the developer.
- Testing Strategy: Test case design, failure analysis, and TDD implementation cycles were performed manually.
- Code Review: All AI suggestions were reviewed, tested, and validated before acceptance.

> Summary: AI was used to speed up typing, not thinking. The engineering judgment, algorithmic correctness, and final implementation are entirely my own.

---

## How to Run

### Prerequisites

- Go (version 1.24 or higher recommended)

### 1. Run Tests

- Execute the full test suite to ensure everything is working:

```bash
go test ./...
```

### 2. Run the Server

- Start the application:

```bash
go run cmd/server/main.go
```

- The server will start on: http://localhost:8080

You can test the health of the API or proceed to create an employee using the endpoints listed above.

### Troubleshooting

If you see an error stating the port is already in use (address already in use), you can kill the process using:

```bash
lsof -ti:8080 | xargs kill -9
```
