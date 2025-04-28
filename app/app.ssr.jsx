import { renderToString } from "react-dom/server"
import { createInertiaApp } from "@inertiajs/react"
import createServer from "@inertiajs/react/server"
import { resolvePageComponent } from "laravel-vite-plugin/inertia-helpers"

createServer(
    (page) =>
        createInertiaApp({
            page,
            render: renderToString,
            resolve: (name) => {
                return resolvePageComponent(
                    `./routes/${name}/page.tsx`,
                    import.meta.glob("./**/*.tsx"),
                )
            },
            setup({ App, props }) {
                return <App {...props} />
            },
        }),
    { cluster: true },
)
