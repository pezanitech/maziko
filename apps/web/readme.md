# Maziko Website

The Maziko website is a web application built with the Maziko Framework itself. It leverages Tailwind CSS for styling with a purpose-based token system and follows a functional-first approach to component development. This readme provides a comprehensive guide to the project's structure, patterns, and best practices.

## Table of Contents
- [Features](#features)
- [Core Design Principles](#core-design-principles)
- [Directory Structure](#directory-structure)
- [File Naming Conventions](#file-naming-conventions)
- [Component Patterns](#component-patterns)
- [Styling System](#styling-system)
- [Routing and Pages](#routing-and-pages)
- [Utilities](#utilities)
- [Testing and Skeletons](#testing-and-skeletons)
- [Performance Considerations](#performance-considerations)

## Features
- **Routes directory**: Is only responsible for providing routing, and page metadata.
- **Tailwind CSS**: Used extensively for styling with dynamic class application using clsx and cn utilities.
- **Purpose-Based Color System**: Custom color tokens mapped to semantic purposes rather than raw color values.
- **Radix UI**: Enhances accessibility and reusability of UI components.

## Core Design Principles

### 1. Development Configuration
- **Prettier Configuration**:
  - Uses @trivago/prettier-plugin-sort-imports for automatic import sorting
  - Uses prettier-plugin-tailwindcss for automatic Tailwind class sorting

### 2. Component Styling Pattern
Components with styles must follow this pattern:

```tsx
import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`@container/card-header has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6 grid auto-rows-min grid-rows-[auto_auto] items-start gap-1.5 px-4 sm:px-6`,
}

export const CardHeader = ({
    className,
    ...props
}: React.ComponentProps<"div">) => (
    <div
        data-slot="card-header"
        className={cn(styles.base, className)}
        {...props}
    />
)
```

### 3. Purpose-Based Color System
Colors are defined using semantic tokens instead of raw color values:

- **Base Colors**:
  - `bg-background`/`foreground`: Main application colors
  - `bg-card`/`card-foreground`: Card component colors
  - `bg-popover`/`popover-foreground`: Popover component colors

- **Functional Colors**:
  - `bg-primary`/`primary-foreground`: Primary actions and highlights
  - `bg-secondary`/`secondary-foreground`: Secondary actions
  - `bg-accent`: Brand color (primary brand identity)
  - `bg-success`: Success states
  - `bg-destructive`: Error states and destructive actions

- **UI Colors**:
  - `bg-muted`/`muted-foreground`: Subdued UI elements
  - `border-border`: Border elements
  - `ring-ring`: Focus rings and outlines

- **Chart Colors**:
  - `bg-chart-1` through `bg-chart-5`: Data visualization

- **Sidebar Colors**:
  - `bg-sidebar`/`sidebar-foreground`: Sidebar-specific theming
  - `bg-sidebar-primary`/`sidebar-primary-foreground`: Sidebar primary elements
  - `bg-sidebar-accent`/`sidebar-accent-foreground`: Sidebar accents

### 4. Functional-First Components
Components follow a strict functional-first approach using arrow function expressions to encourage Pure Functional Programming:

```tsx
// Direct return when no additional logic needed
const SimpleComponent = () => <div>Simple Component</div>

// Block syntax only when required
const ComplexComponent = () => {
    const [state, setState] = useState()
    useEffect(() => {
        // Side effects
    }, [])
    return <div>{state}</div>
}
```

### 5. Component Patterns
- **UI Components**: Use React.ComponentProps for type inheritance and focus on styling unstyled components or creating reusable, accessible components.
  
  Example:
  ```tsx
  export const DropdownMenuItem = ({
      className,
      inset,
      variant = "default",
      ...props
  }: React.ComponentProps<typeof DropdownMenuPrimitive.Item>) => (
      <DropdownMenuPrimitive.Item
          data-slot="dropdown-menu-item"
          data-inset={inset}
          data-variant={variant}
          className={cn(styles.base, className)}
          {...props}
      />
  )
  ```

- **Page Components**: Props must always be used from the props scope, not destructured. Example:
  ```tsx
  const PageComponent = (props: PropsType) => <Markup prop={props.prop} />
  ```

- **Styling**: Scoped styles defined within component files.
- **Error Handling**: Error boundaries for graceful degradation.

## Directory Structure

### 1. Core Directories
The project is organized into the following core directories:

```
.
├── app                    # Application source code
│   ├── components         # React components
│   │   ├── providers      # Global context providers
│   │   ├── ui             # Reusable UI components
│   │   ├── layout         # Layout components
│   │   └── pages          # Page components
│   ├── lib                # Utility functions and helpers
│   ├── public             # Static assets
│   ├── routes             # Application routes
│   ├── styles             # Global and shared styles
│   ├── types              # TypeScript type definitions
│   ├── app.jsx            # Application entry point
│   └── global.css         # Global styles
├── gen                    # Generated files
│   └── routesgen.go       # Auto-generated route definitions
├── main.go                # Go application entry point
├── maziko.json            # Project configuration
└── vite.config.ts         # Vite configuration
```

This structure ensures modularity and scalability, with clear separation of concerns between application logic, UI components, and application configuration.

### 2. Component Organization
Components are grouped by feature or purpose:

- **UI Components**:
  - Reusable and styled components are placed under `app/components/ui`.
  - Example: `app/components/ui/dropdownMenu` contains all files related to the dropdown menu component.

- **Page Components**:
  - Components specific to a page are placed under `app/components/pages`.
  - Example: `app/components/pages/home` contains components for the home page.

- **Layout Components**:
  - Shared layout components like header and footer are placed under `app/components/layout`.

### 3. Providers
**Purpose**: The providers directory contains global context providers or higher-order components that manage application-wide state or settings.

**Key Components**:
- `themeProvider.tsx`: Handles theming, such as light/dark mode or custom themes.

**Usage**:
- These providers are wrapped around the application in app.jsx to ensure consistent state management and accessibility across all routes.

Example:
```tsx
// In app.jsx
import { ThemeProvider } from "@/components/providers/themeProvider"

createInertiaApp({
    resolve: resolveComponent,
    setup({ el, App, props }) {
        createRoot(el).render(
            <ThemeProvider>
                <App {...props} />
            </ThemeProvider>
        )
    },
})
```

## File Naming Conventions

The naming and structure of components was chosen in order to make components easy to plug in and out and to make the relationships between components more visible at a glance.

### 1. General Naming Rules
- **Camel Cased Parent Scoped Naming**:
  - Files are named in camelCase, with descriptive scoped names indicating their parent and purpose.
  - Example: `dropdownMenu.trigger.tsx`, `companies.header.tsx`.

- **Index Files**:
  - `index.ts` files are used for re-exporting components or utilities within a directory.
  - Example: `src/components/ui/dropdownMenu/index.ts`.

### 2. Component Naming
- **PascalCase for Components**:
  - React components are named in PascalCase.
  - Example: `DropdownMenu`, `CompaniesHeader`.

- **Descriptive Names**:
  - Components and files are named descriptively to indicate their purpose.
  - Example: `dropdownMenu.shortcut.tsx`, `companies.priceInfo.tsx`.

## Styling System

### 1. CSS Variables
Color tokens are defined in `globals.css`:

```css
:root {
    --background: oklch(0.98 0 0);
    --accent: oklch(0.48 0.0737 203.79);
    /* Other color tokens */
}

.dark {
    --background: oklch(0.145 0 0);
    /* Dark theme overrides */
}
```

### 2. Usage in Components
```tsx
import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`bg-background text-foreground hover:bg-accent`,
}

const Component = ({ className, ...props }: React.ComponentProps<"div">) => (
    <div
        className={cn(styles.base, className)}
        {...props}
    >
        Purpose-based styling
    </div>
)
```

### 3. Dynamic Styling
- `clsx` is used for conditional class application.
- `cn` is used for merging classes.

## Routing and Pages

### 1. App Directory
- Routes are defined in `app/routes`, following the app directory structure.
- Example: `app/routes/companies/_companyid/page.tsx` for dynamic routes.

### 2. Page Components
- Route components import and render components from `app/components/pages`.
```tsx
import { PageComponent } from "@/components/pages/pageComponent"

const Page = () => <PageComponent />

export default Page
```

## Utilities

### 1. Utility Functions
- Shared utility functions are placed in `app/lib` (e.g., `app/lib/vite.ts`, `app/lib/utils.ts`).

### 2. Reusable Types
- Type definitions are centralized in `app/types` (e.g., `app/types/company.ts`).

## Testing and Skeletons

### 1. Skeleton Components
- Placeholder components (e.g., `CardSkeleton`) are used for loading states.

### 2. Error Boundaries
- Components like tables use error boundaries for graceful error handling.

## Performance Considerations

### 1. Component Optimization
- Functional-first approach reduces complexity and minimizes the surface area for bugs to exist.
- Minimal prop drilling.
- Error boundary placement.

### 2. Styling Performance
- Purpose-based tokens reduce CSS size.
- Tailwind compilation and removal of unused styles.
- Scoped styles prevent conflicts.