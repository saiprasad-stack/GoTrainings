# Go Project Layout

```
myproject/
├── go.mod
├── go.sum
├── Makefile
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── auth.go
│   │   └── auth_test.go
│   ├── db/
│   │   ├── db.go
│   │   └── db_test.go
│   ├── models/
│   │   └── models.go
│   └── services/
│       └── service.go
├── pkg/
│   └── utils/
│       └── stringutils.go
├── api/
│   ├── proto/
│   └── openapi/
├── configs/
│   └── config.yaml
├── scripts/
│   ├── build.sh
│   └── deploy.sh
└── test/
└── integration_test.go
```
---

## Description of Each Directory

## Root Level

### `go.mod`
Defines the module name and manages dependencies.

### `go.sum`
Stores checksums for dependencies to ensure integrity.

### `Makefile`
Defines common commands such as:
- `make build`
- `make run`
- `make test`

---

## `cmd/` (Entry Points)

Contains executable applications.

### Purpose
Each subdirectory represents a **separate binary**.

### Example
- `cmd/api` → API server
- `cmd/worker` → background job processor

### Key Idea
Keep `main.go` minimal:
- Initialize dependencies
- Start application

---

## `internal/` (Private Application Code)

Contains code that **cannot be imported by other modules**.

### Purpose
- Enforces encapsulation
- Safe refactoring without breaking external users

---

### `internal/auth/`
Authentication and authorization logic.

**Responsibilities**
- JWT handling
- OAuth
- Middleware

---

### `internal/db/`
Database layer.

**Responsibilities**
- DB connection setup
- Query execution
- ORM usage (e.g., GORM)

---

### `internal/models/`
Domain models (data structures).

**Responsibilities**
- Struct definitions
- DB mappings
- Validation tags

---

### `internal/services/`
Business logic layer.

**Responsibilities**
- Core application rules
- Coordination between DB and handlers

---

## `pkg/` (Public Packages - Optional)

Reusable code that **can be imported by other projects**.

### `pkg/utils/`
Generic helper functions.

**Examples**
- String utilities
- Common helpers
- Formatting logic

### Important Note
Do NOT put business logic here unless it is truly reusable.

---

## `api/` (API Definitions)

Defines external interfaces.

### `api/proto/`
- gRPC `.proto` files
- Used to generate Go code

### `api/openapi/`
- Swagger / OpenAPI specifications
- Used for REST API documentation

---

## `configs/`

Stores configuration files.

### Example
- `config.yaml`

**Usage**
- Environment configs
- DB settings
- Service endpoints

---

## `scripts/`

Automation scripts.

### Examples
- `build.sh`
- `deploy.sh`

**Purpose**
- CI/CD tasks
- Local automation

---

## `test/` (Integration Tests)

Contains external tests.

