# Records Keeper Frontend
The Records Keeper (RK) frontend is a web application responsible for providing the user interface for the Records Keeper platform. It communicates with the backend via the REST API to display, create, edit, and manage documents, files, portals, and user data.

# Responsibilities
The frontend is responsible for:
- Rendering the user interface
- Displaying documentation and personal records
- Providing the TipTap document editor
- Managing client-side routing
- Handling user interactions
- Managing application state
- Uploading files through the backend API
- Displaying notifications and activity
- Providing portal administration interfaces
- Communicating with the backend REST API

The frontend does not contain business logic or directly access the database. All persistent data is managed by the backend.

# Project Structure
```
frontend/
├── public/              # Static assets
├── src/
│   ├── api/             # API client and requests
│   ├── assets/          # Images, icons, fonts
│   ├── components/      # Reusable UI components
│   ├── features/        # Feature-specific components
│   ├── hooks/           # Custom React hooks
│   ├── layouts/         # Application layouts
│   ├── pages/           # Route pages
│   ├── router/          # Client-side routing
│   ├── stores/          # Global state management
│   ├── styles/          # Global styles
│   ├── types/           # Shared TypeScript types
│   ├── utils/           # Utility functions
│   ├── App.tsx
│   └── main.tsx
├── package.json
├── tsconfig.json
└── README.md
```