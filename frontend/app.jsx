import { createRoot } from "react-dom/client"
import { createInertiaApp } from "@inertiajs/react"
import { resolvePageComponent } from "laravel-vite-plugin/inertia-helpers"

createInertiaApp({
    resolve: (name) => {
        return resolvePageComponent(
            `./app/${name}.tsx`,
            import.meta.glob("./**/*.tsx"),
        )
    },
    setup({ el, App, props }) {
        const root = createRoot(el)

        root.render(<App {...props} />)
    },
})
