# Maziko

Maziko is a full-stack web framework combining the power of Go for the backend with modern frontend technologies (React, Vue, Svelte), connected through Inertia.js for a seamless single-page application experience.

## 🚀 Features

-   **Full-Stack Go & Frontend**: Develop your entire application with Go on the backend and your choice of frontend technology
-   **Multiple Frontend Templates**: Choose from React, Vue, or Svelte templates to start your project
-   **Single-Page Application Feel**: Uses Inertia.js to provide SPA-like navigation without API boilerplate
-   **File-Based Routing**: Automatic route generation based on directory structure
-   **Server-Side Rendering**: Built-in SSR support for improved SEO and initial load performance
-   **Hot Module Replacement**: Fast development feedback with HMR for both Go and frontend code
-   **CLI Tool**: Easy project scaffolding and management with the `maziko` command-line tool
-   **Tailwind CSS**: Modern utility-first CSS framework included out of the box
-   **Vite Building**: Fast builds and development server with Vite
-   **Nx Monorepo**: Efficient workspace management for multiple projects and templates

## 📋 Getting Started

### Prerequisites

-   Go 1.24.2
-   Node.js >= 18
-   pnpm >= 10 

### CLI Installation

Install the Maziko CLI globally:

```bash
go install github.com/pezanitech/maziko/cli/maziko@latest
```

### Creating a New Project

Use the Maziko CLI to create a new project with your preferred frontend:

```bash
# Create a React project
maziko create react -n my-app

# Create a Vue project (coming soon)
maziko create vue -n my-app

# Create a Svelte project (coming soon)
maziko create svelte -n my-app
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
  completion  Generate the autocompletion script for the specified shell
  create      Create a new Maziko project
  help        Help about any command

Flags:
  -h, --help      help for maziko
  -v, --verbose   Enable verbose output
      --version   version for maziko
```

## 🏗️ Project Structure

```
your-project/
├── app/                     # Application code
│   ├── global.css           # Global CSS imports
│   ├── app.jsx              # Main frontend entry point
│   ├── public/              # Static public assets
│   └── routes/              # Routes (both backend and frontend)
│       ├── index/           # Home page route
│       │   ├── handler.go   # Go route handler
│       │   └── page.tsx     # Frontend component for route
│       └── docs/            # Another route example
├── gen/                     # Auto-generated files
├── build/                   # Production build output
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
        <div className="p-8">
            <h1 className="text-3xl font-bold">{props.title}</h1>
            <p className="mt-2">{props.description}</p>
        </div>
    )
}
```

4. Run `pnpm genroutes` to generate the route or restart the development server

### Building for Production

```bash
pnpm build
```

This will:

1. Build the Vite assets (both client and SSR versions)
2. Generate route definitions
3. Build the Go binary

### Starting the Production Server

```bash
pnpm start
```

## 🛠️ Available Scripts

-   `pnpm dev` - Start development server with hot reloading
-   `pnpm build` - Build for production
-   `pnpm start` - Start production server
-   `pnpm genroutes` - Generate route definitions
-   `pnpm clean` - Clean temporary files
-   `pnpm reset` - Clean everything and reinstall dependencies

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
-   `dev:live` - Start development server with live reloading
-   `build` - Build for production (generates routes, builds frontend and backend)
-   `build:vite` - Build only the frontend assets with Vite
-   `build:go` - Build only the Go backend
-   `build:go:tmp` - Build Go to temporary directory for development
-   `genroutes` - Generate route definitions
-   `start` - Start production server
-   `start:ssr` - Start production server with SSR
-   `clean` - Clean temporary files
-   `reset` - Clean everything and reinstall dependencies
-   `test` - Run tests (specific to each project)
-   `lint` - Run linters
-   `tidy` - Run Go mod tidy

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

### Monorepo Development

For development on the Maziko framework itself:

1. Clone the repository:

    ```bash
    git clone https://github.com/pezanitech/maziko.git
    cd maziko
    ```

2. Install dependencies:

    ```bash
    pnpm install
    ```

3. Build the framework:

    ```bash
    pnpm build
    ```

4. Run tests:
    ```bash
    pnpm test
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

## 👥 Contributing

We welcome contributions to Maziko! Please see [contributing.md](contributing.md) for guidelines on how to report bugs, suggest features, and submit pull requests.

## 📄 License

[MIT](license)

## 🙏 Acknowledgements

Maziko is built on top of several amazing open-source technologies:

-   [Go](https://golang.org/)
-   [React](https://react.dev/)
-   [Vue](https://vuejs.org/)
-   [Svelte](https://svelte.dev/)
-   [Inertia.js](https://inertiajs.com/)
-   [Vite](https://vite.dev/)
-   [Gonertia](https://github.com/romsar/gonertia/)
-   [Tailwind CSS](https://tailwindcss.com/)
-   [Nx](https://nx.dev/)
