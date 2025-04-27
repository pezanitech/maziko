import { createRoot } from "react-dom/client"
import { renderToString } from "react-dom/server"
import { createInertiaApp } from "@inertiajs/react"
import createServer from "@inertiajs/react/server"
import { resolvePageComponent } from "laravel-vite-plugin/inertia-helpers"

import "./global.css"

createServer(
    (page) =>
        createInertiaApp({
            page,
            render: renderToString,
            resolve: (name) => {
                return resolvePageComponent(
                    `./app/${name}.tsx`,
                    import.meta.glob("./**/*.tsx"),
                )
            },
            setup({ App, props }) {
                return <App {...props} />
            },
        }),
    { cluster: true },
)
