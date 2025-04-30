# Maziko

Maziko is a full-stack web framework combining the power of Go for the backend and React for the frontend, connected through Inertia.js for a seamless single-page application experience.

## 🚀 Features

- **Full-Stack Go & React**: Develop your entire application with Go on the backend and React on the frontend
- **Single-Page Application Feel**: Uses Inertia.js to provide SPA-like navigation without API boilerplate
- **File-Based Routing**: Automatic route generation based on directory structure
- **Server-Side Rendering**: Built-in SSR support for improved SEO and initial load performance
- **Hot Module Replacement**: Fast development feedback with HMR for both Go and React code
- **Tailwind CSS**: Modern utility-first CSS framework included out of the box
- **Vite Building**: Fast builds and development server with Vite

## 🏗️ Project Structure

```
maziko/
├── app/                     # Application code
│   ├── global.css           # Global CSS imports
│   ├── app.jsx              # Main React entry point
│   ├── public/              # Static public assets
│   └── routes/              # Routes (both backend and frontend)
│       ├── index/           # Home page route
│       │   ├── handler.go   # Go route handler
│       │   └── page.tsx     # React component for route
│       └── docs/            # Another route example
├── core/                    # Framework core
│   ├── cmd/                 # CLI commands
│   ├── config/              # Framework configuration
│   ├── router/              # Routing and Inertia.js integration
│   └── utils/               # Utility functions
├── gen/                     # Auto-generated files
├── build/                   # Production build output
└── tmp/                     # Development build output
```

## 📋 Getting Started

### Prerequisites

- Go 1.18 or higher
- Node.js 16 or higher
- pnpm (recommended) or npm

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/pezanitech/maziko.git
   cd maziko
   ```

2. Install JavaScript dependencies:
   ```bash
   pnpm install
   ```

3. Start the development server:
   ```bash
   pnpm dev
   ```

This will start both the Go server and the Vite development server with hot reloading.

## 🧪 Development Workflow

### Creating a New Route

1. Create a new directory in `app/routes/` with your route name (e.g., `app/routes/about/`)
2. Add a `handler.go` file with the HTTP methods you need:
   ```go
   package about

   import (
       "net/http"
       
       "github.com/pezanitech/maziko/core/utils"
       inertia "github.com/romsar/gonertia"
   )

   func GET(i *inertia.Inertia, w http.ResponseWriter, r *http.Request) {
       err := i.Render(w, r, "about", inertia.Props{
           "title": "About Us",
           "description": "Learn more about our team",
       })

       if err != nil {
           utils.HandleServerErr(w, err)
       }
   }
   ```

3. Create a `page.tsx` file with your React component:
   ```tsx
   import { usePage } from "@inertiajs/react"

   export default function Page() {
       const { title, description } = usePage().props

       return (
           <div className="p-8">
               <h1 className="text-3xl font-bold">{title}</h1>
               <p className="mt-2">{description}</p>
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

- `pnpm dev` - Start development server with hot reloading
- `pnpm build` - Build for production
- `pnpm start` - Start production server
- `pnpm genroutes` - Generate route definitions
- `pnpm clean` - Clean temporary files
- `pnpm reset` - Clean everything and reinstall dependencies

## 🧩 How Maziko Works

### Routing

Maziko uses a file-based routing system:
- Each subdirectory in `app/routes/` becomes a route
- URL paths match directory names (`app/routes/about/` → `/about`)
- `app/routes/index/` maps to the root path (`/`)
- Each route needs:
  - `handler.go` - Backend logic and data loading
  - `page.tsx` - Frontend React component

### Inertia.js Integration

Maziko uses Inertia.js to enable:
- SPA-like navigation without API boilerplate
- Server-side data loading with automatic prop passing
- Server-side rendering for initial page loads

### Development Tools

The framework includes several development tools:
- Hot reloading for Go code
- HMR for React components
- Automatic route generation

## 📄 License

[MIT](LICENSE)

## 🙏 Acknowledgements

Maziko is built on top of several amazing open-source technologies:
- [Go](https://golang.org/)
- [React](https://react.dev/)
- [Inertia.js](https://inertiajs.com/)
- [Vite](https://vite.dev/)
- [Gonertia](https://github.com/romsar/gonertia/)