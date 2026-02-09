# Markito - Agent Guidelines

## Project Overview

Markito is a Go web application for editing and sharing Markdown documents. It uses:
- **Leapkit Core** for server/routing and database
- **Gomponents** for type-safe HTML templating
- **HTMX** for frontend interactivity
- **SQLite** with WAL mode for data persistence
- **Tailwind CSS** for styling
- **Kamal** for deployment

## Commands

| Command | Description |
|---------|-------------|
| `go run ./cmd/app` | Run the application locally |
| `go run ./cmd/migrate` | Run database migrations |
| `go test ./...` | Run all tests |
| `go test -run TestName ./package` | Run a single test |
| `go tool tailo --i internal/system/assets/tailwind.css -o internal/system/assets/application.css` | Build Tailwind CSS |
| `go tool tailo download -v v4.0.6` | Download Tailwind binary |
| `go build -o bin/app ./cmd/app` | Build application binary |
| `go build -o bin/migrate ./cmd/migrate` | Build migrator binary |
| `kamal deploy` | Deploy to production |
| `kamal redeploy` | Redeploy without building new image |

## Code Style

### Imports
Organize in 3 groups with blank lines between:
1. Standard library
2. Third-party packages
3. Local/internal packages

```go
import (
    "net/http"
    "os"

    "github.com/example/lib"
    "maragu.dev/gomponents"

    "markito/internal/documents"
)
```

### Formatting
- Use `gofumpt` for stricter formatting
- Run `go fmt ./...` before committing
- Line length: ~100 characters soft limit

### Naming Conventions
- **Packages**: lowercase, single word (e.g., `documents`, `markdown`)
- **Exported**: PascalCase (e.g., `Save`, `ToHTML`)
- **Unexported**: camelCase (e.g., `dbFn`, `defaultMarkdown`)
- **Interfaces**: noun describing capability (e.g., `Server`)
- **Structs**: noun describing entity (e.g., `Document`, `service`)
- **HTTP handlers**: Verb describing action (e.g., `Save`, `Open`, `New`)
- **Private service methods**: receiver is `s`, methods are verbs (e.g., `s.Save()`, `s.Find()`)

### Error Handling
- Return errors, don't log and swallow
- Wrap with context: `fmt.Errorf("saving document: %w", err)`
- HTTP handlers respond with appropriate status codes
- Services return `nil, err` or `zero, err` on failure

### Types & Structs
- Struct tags use double quotes: `` `db:"id"` ``
- Keep structs simple (ID, Content for Document)
- Service pattern: unexported struct with exported constructor
- Prefer concrete returns, accept interfaces

### SQL
- Use `$1`, `$2` placeholders (PostgreSQL/SQLite style)
- Multi-line queries with proper indentation
- Always check for nil DB connections

### HTTP Handlers
- Use standard `http.HandlerFunc` signature
- Extract services from context with type assertion
- Use `r.PathValue("id")` for URL parameters (Go 1.22+)
- Use `r.FormValue("key")` for form data
- Set HTMX headers where needed (e.g., `Hx-Push`)

### HTML/Gomponents
- Use dot import for gomponents: `. "maragu.dev/gomponents"`
- Alias HTMX import: `hx "maragu.dev/gomponents-htmx"`
- Element functions return `Node` or specific types
- Page function wraps content in layout
- Components are exported functions returning `Node`

## Architecture

```
cmd/
  app/         # Main application entrypoint
  migrate/     # Database migration entrypoint
internal/
  app.go       # Server setup and routing
  documents/   # Document domain (handlers, service, views)
  markdown/    # Markdown parsing utilities
  migrations/  # SQL migration files
  system/
    assets/    # Static assets (CSS, JS)
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DATABASE_URL` | `markito.db` | SQLite database path |
| `HOST` | `0.0.0.0` | Server host |
| `PORT` | `3000` | Server port |
| `SESSION_SECRET` | hardcoded | Session cookie secret |
| `SESSION_NAME` | `markito_session` | Session cookie name |
| `BASE_URL` | `http://localhost:3000` | URL for document links |

## Testing

- Place tests in `_test.go` files alongside source
- Use table-driven tests with `t.Run()`
- Run single test: `go test -run TestFunction ./package`
- Integration tests use SQLite in-memory or temp files

## Migrations

- Files: `internal/migrations/*.sql`
- Naming: `YYYYMMDDhhmmss_description.sql`
- SQL: Use PRAGMAs for SQLite tuning (WAL, busy_timeout)
- Run with: `go run ./cmd/migrate`

## Deployment

- Uses Kamal 2.8.1 with Docker
- Ruby 3.3.0 required for Kamal
- Docker multi-stage build with Alpine
- Assets compiled during build phase
- Migrations run before app startup
