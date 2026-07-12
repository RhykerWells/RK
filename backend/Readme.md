# Records Keeper Backend
The Records Keeper (RK) backend is a REST API written in Go. It provides the core services that power the Records Keeper platform, including authentication, document management, file storage, permissions, portal management, and integrations.

# Responsibilities
The backend is responsible for:
- Serving the HTTP API
- Authenticating users
- Authorizing requests using roles and permissions
- Managing portals (workspaces)
- Managing documents and their version history
- Managing uploaded files
- Integrating with Discord
- Persisting data in PostgreSQL
- Providing a stable API for the frontend

The backend does not render HTML pages. Presentation is handled entirely by the frontend application.

# Project Structure
```
backend/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── api/             # HTTP routing and handlers
│   ├── auth/            # Authentication
│   ├── config/          # Configuration loading
│   ├── database/        # Database connection
│   ├── discord/         # Discord integration
│   ├── documents/       # Document services
│   ├── files/           # File management
│   ├── permissions/     # Role & permission logic
│   ├── portals/         # Workspace management
│   ├── users/           # User management
│   └── versioning/      # Document revision history
├── migrations/          # SQL migrations
├── pkg/                 # Reusable public packages (if needed)
├── go.mod
└── README.md
```