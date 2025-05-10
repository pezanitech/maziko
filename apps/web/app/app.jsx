import { createInertiaApp } from "@inertiajs/react"
import { resolvePageComponent } from "laravel-vite-plugin/inertia-helpers"

import { Footer } from "@/components/layout/footer"
import { Header } from "@/components/layout/header"

// Component resolver function
const resolveComponent = (name) => {
    return resolvePageComponent(
        `./routes/${name}/page.tsx`,
        import.meta.glob("./**/*.tsx"),
    )
}

// CLIENT-SIDE ENTRY POINT
const createClientApp = async () => {
    const { createRoot } = await import("react-dom/client")
    const { ThemeProvider } = await import(
        "@/components/providers/themeProvider"
    )

    createInertiaApp({
        resolve: resolveComponent,
        setup({ el, App, props }) {
            createRoot(el).render(
                <ThemeProvider>
                    <Header />
                    <App {...props} />
                    <Footer />
                </ThemeProvider>,
            )
        },
    })
}

// SERVER-SIDE ENTRY POINT
const createSSRApp = async () => {
    if (import.meta.env.SSR) {
        const { renderToString } = await import("react-dom/server")
        const { default: createServer } = await import(
            "@inertiajs/react/server"
        )

        createServer((page) =>
            createInertiaApp({
                page,
                render: renderToString,
                resolve: resolveComponent,
                setup({ App, props }) {
                    return (
                        <>
                            <Header />
                            <App {...props} />
                            <Footer />
                        </>
                    )
                },
            }),
        )
    }
}

// Select entry point based on environment
if (import.meta.env.SSR) {
    createSSRApp() // SSR
} else {
    createClientApp() // Browser
}
