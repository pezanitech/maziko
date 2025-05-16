# Maziko

Maziko is a full-stack web framework combining the power of Go for backend functionality with modern frontend technologies (React, Vue (coming soon), Svelte (coming soon)), connected through Inertia.js for a seamless single-page application experience.

## 🚀 Features

-   **Full-Stack Go & Frontend**: Develop your entire application with Go for backend code and your choice of frontend technology
-   **Multiple Frontend Templates**: Choose from React, Vue (coming soon), or Svelte (coming soon) templates to start your project
-   **Single-Page Applications**: Uses Inertia.js to provide SPA navigation without API boilerplate
-   **File-Based Routing**: Automatic route generation based on directory structure
-   **Server-Side Rendering**: Built-in SSR support for improved SEO and initial load performance
-   **Live Reloading**: Fast development feedback with live reloading for both Go and frontend code
-   **CLI Tool**: Easy project scaffolding and management with the `maziko` command-line tool
-   **Vite Building**: Fast builds and development server with Vite

## 📋 Getting Started

### Prerequisites

-   Go >= 1.24.2
-   Node.js >= 18
-   pnpm >= 10 

### CLI Installation

Install the Maziko CLI globally:

```bash
go install github.com/pezanitech/maziko/cli/maziko@latest
```

### Creating a New Project

Use the Maziko CLI to create a new project

```bash
# Create a React project
maziko create react -n my-app
```

Then follow the next steps:

```bash
cd my-app
pnpm install
pnpm dev
```

### CLI Commands

The `maziko` CLI tool provides several commands:

```
Usage:
  maziko [flags]
  maziko [command]

Available Commands:
  build       Build the project
  clean       Clean build artifacts
  completion  Generate the autocompletion script for the specified shell
  create      Create a new Maziko project
  dev         Start development server
  genroutes   Generate routes
  help        Help about any command
  install     Install dependencies
  start       Start the production server

Flags:
  -h, --help      help for maziko
  -v, --verbose   Enable verbose output (not implemented)
      --version   version for maziko
```

## 🏗️ Project Structure

```
project/
├── app/                     # Application code
│   ├── global.css           # Global CSS imports
│   ├── app.jsx              # Main frontend entry point
│   ├── public/              # Static public assets
│   └── routes/              # Routes (both backend and frontend)
│       ├── index/           # Root route
│       │   ├── handler.go   # Go route handler
│       │   └── page.tsx     # Frontend page for route
│       └── docs/            # Another route example
├── gen/                     # Auto-generated files
├── build/                   # Production build output
├── ssrBuild/                # SSR build output
└── tmp/                     # Development build output
```

## 🧪 Development Workflow

### Creating a New Route

1. Create a new directory in `app/routes/` with your route name (e.g., `app/routes/about/`)
2. Add a `handler.go` file with the HTTP methods you need:

```go
package about

import (
    "net/http"

    "github.com/pezanitech/maziko/libs/core/router"
    inertia "github.com/romsar/gonertia"
)

func Route() {
    router.GET(func(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
        router.RenderPage(i, w, r, inertia.Props{
            "title": "About Us",
            "description": "Learn more about our team",
        })
    })
}
```

3. Create a `page.tsx` file with your frontend component (React example):

```tsx
import { usePage } from "@inertiajs/react"

export default function Page() {
    const { props } = usePage()

    return (
        <div>
            <h1>{props.title}</h1>
            <p>{props.description}</p>
        </div>
    )
}
```

4. Run `maziko genroutes` to generate the route or restart the development server

### Building for Production

```bash
maziko build
```

This will:

1. Build the Vite assets (both client and SSR versions)
2. Generate route definitions
3. Build the Go binary

### Starting the Production Server

```bash
maziko start
```

## 📜 Configuration

Maziko uses a configuration file (`maziko.json`) located in your project root:

```json
{
    "app": {
        "name": "Maziko App",
        "url": "http://localhost:3080",
        "port": 3080
    },
    "build": {
        "prefix": "/build/",
        "dir": "./build",
        "tempDir": "./tmp",
        "ssrDir": "./ssrBuild"
    },
    "vite": {
        "manifestFile": "./build/.vite/manifest.json",
        "hotFile": "tmp/hot",
        "detectionAttempts": 5,
        "detectionInterval": 2000
    },
    "paths": {
        "routes": "./app/routes",
        "public": "./app/public",
        "gen": "./gen"
    },
    "logger": {
        "type": "concise",
        "level": "debug"
    }
}
```

Environment variables can override configuration values:

-   `APP_URL` - URL for the application
-   `LOGGER_TYPE` - Logger type (text, json, concise)
-   `LOG_LEVEL` - Logging level (debug, info, warn, error)

## 🧩 Monorepo Structure

Maziko is organized as an Nx monorepo with the following structure:

```
maziko/
├── libs/                 # Core libraries
│   ├── core/             # Core framework functionality
│   └── maziko/           # CLI tool implementation
├── templates/            # Project templates
│   ├── react/            # React template
│   ├── vue/              # Vue template (coming soon)
│   └── svelte/           # Svelte template (coming soon)
└── package.json          # Root package.json
```

## 📦 Nx Commands

Maziko uses Nx for managing the monorepo and optimizing the build process. Here are some useful Nx commands:

-   `nx run [project]:build` - Build a specific project
-   `nx run-many --target=build` - Build all projects
-   `nx affected --target=build` - Build only affected projects
-   `nx graph` - Visualize the project dependencies
-   `nx dev --project=[project]` - Start development server for a specific project
-   `nx build --project=[project]` - Build a specific project

### Available Nx Targets

-   `dev` - Start development server
-   `build` - Build for production (generates routes, builds frontend and backend)
-   `start` - Start production server
-   `clean` - Clean temporary files
-   `reset` - Clean everything and reinstall dependencies
-   `test` - Run tests (specific to each project)
-   `lint` - Run linters
-   `tidy` - Run Go mod tidy

## 📄 License

[MIT](license)

## 🙏 Acknowledgements

Maziko is built on top of several amazing open-source technologies:

-   [Go](https://golang.org/)
-   [React](https://react.dev/)
-   [Inertia.js](https://inertiajs.com/)
-   [Vite](https://vite.dev/)
-   [Nx](https://nx.dev/)
