import { defineConfig, UserConfig } from "vite"
import laravel from "laravel-vite-plugin"
import react from "@vitejs/plugin-react"
import tailwindcss from "@tailwindcss/vite"
import fs from "fs"
import path from "path"
import { fileURLToPath } from "url"
import { defineAlias } from "@/lib/vite"

// create tmp directory
const tmpDir = path.resolve(__dirname, "tmp")
if (!fs.existsSync(tmpDir)) {
    fs.mkdirSync(tmpDir, { recursive: true })
    console.log("created tmp directory")
}

const common = {
    resolve: {
        alias: [defineAlias("@", "./app")],
    },
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

const clientConfig: UserConfig = {
    resolve: common.resolve,
    plugins: [
        ...common.plugins,
        laravel({
            input: [...common.inputFiles, "app/app.jsx"],
            publicDirectory: "tmp",
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
    resolve: common.resolve,
    plugins: [
        ...common.plugins,
        laravel({
            input: [...common.inputFiles, "app/app.jsx"],
            ssr: "app/app.jsx", // SSR Entry point
            publicDirectory: "tmp",
            buildDirectory: "ssrBuild",
            refresh: true,
        }),
    ],
    build: {
        ssr: true,
        outDir: "ssrBuild",
        rollupOptions: {
            input: [...common.inputFiles, "app/app.jsx"],
            output: common.rollupOutput,
        },
    },
}

export default defineConfig(({ isSsrBuild }) => {
    if (isSsrBuild) {
        return ssrConfig
    }

    return clientConfig
})
