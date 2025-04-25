import { defineConfig } from "vite"
import laravel from "laravel-vite-plugin"
import react from "@vitejs/plugin-react"
import tailwindcss from "@tailwindcss/vite"

export default defineConfig({
    plugins: [
        tailwindcss(),

        react({
            include: /\.(mdx|js|jsx|ts|tsx)$/,
        }),

        laravel({
            input: "frontend/app.jsx",
            publicDirectory: ".tmp",
            buildDirectory: "build",
            refresh: true,
        }),
    ],
    build: {
        manifest: true, // Generate manifest.json file
        outDir: "build",
        rollupOptions: {
            input: "frontend/app.jsx",
            output: {
                entryFileNames: "assets/[name].js",
                chunkFileNames: "assets/[name].js",
                assetFileNames: "assets/[name].[ext]",
                manualChunks: undefined, // Disable automatic chunk splitting
            },
        },
    },
    server: {
        hmr: {
            host: "localhost",
        },
    },
})
