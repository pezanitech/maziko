import { defineConfig, UserConfig } from "vite"
import laravel from "laravel-vite-plugin"
import react from "@vitejs/plugin-react"
import tailwindcss from "@tailwindcss/vite"
import fs from "fs"
import path from "path"

// create .tmp directory
const tmpDir = path.resolve(__dirname, ".tmp")
if (!fs.existsSync(tmpDir)) {
    fs.mkdirSync(tmpDir, { recursive: true })
    console.log("created .tmp directory")
}

const common = {
    plugins: [
        tailwindcss(),
        react({
            include: /\.(mdx|js|jsx|ts|tsx)$/,
        }),
    ],

    inputFiles: ["app/global.css"],

    rollupOutput: {
        entryFileNames: "assets/[name].js",
        chunkFileNames: "assets/[name].js",
        assetFileNames: "assets/[name].[ext]",
        manualChunks: undefined, // Disable automatic chunk splitting
    },
}

const defaultConfig: UserConfig = {
    plugins: [
        ...common.plugins,
        laravel({
            input: [...common.inputFiles, "app/app.jsx"],
            publicDirectory: ".tmp",
            buildDirectory: "build",
            refresh: true,
        }),
    ],
    build: {
        manifest: true, // Generate manifest.json file
        outDir: "build",
        rollupOptions: {
            input: [...common.inputFiles, "app/app.jsx"],
            output: common.rollupOutput,
        },
    },
    server: {
        hmr: {
            host: "localhost",
        },
    },
}

const ssrConfig = {
    plugins: [
        ...common.plugins,
        laravel({
            input: [...common.inputFiles, "app/app.ssr.jsx"],
            ssr: "app/app.ssr.jsx", // SSR Entry point
            publicDirectory: ".tmp",
            buildDirectory: "ssrBuild",
            refresh: true,
        }),
    ],
    build: {
        ssr: true,
        outDir: "ssrBuild",
        rollupOptions: {
            input: [...common.inputFiles, "app/app.ssr.jsx"],
            output: common.rollupOutput,
        },
    },
}

export default defineConfig(({ isSsrBuild }) => {
    if (isSsrBuild) {
        return ssrConfig
    }

    return defaultConfig
})
