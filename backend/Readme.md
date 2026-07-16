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