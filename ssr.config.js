import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"
import laravel from "laravel-vite-plugin"

export default defineConfig({
    plugins: [
        laravel({
            input: ["frontend/js/app.jsx", "frontend/css/app.css"],
            ssr: "frontend/js/ssr.jsx", // Enable SSR
            publicDirectory: "public",
            buildDirectory: "bootstrap",
            refresh: true,
        }),
        react(),
    ],
    build: {
        ssr: true, // Enable SSR
        outDir: "bootstrap",
        rollupOptions: {
            input: "frontend/js/ssr.jsx",
            output: {
                entryFileNames: "assets/[name].js",
                chunkFileNames: "assets/[name].js",
                assetFileNames: "assets/[name][extname]",
                manualChunks: undefined, // Disable automatic chunk splitting
            },
        },
    },
})
