import tailwindcss from "@tailwindcss/vite"
import react from "@vitejs/plugin-react"
import fs from "fs"
import laravel from "laravel-vite-plugin"
import path from "path"
import { fileURLToPath } from "url"
import { defineConfig, UserConfig } from "vite"

// create tmp directory
const tmpDir = path.resolve(__dirname, "tmp")
if (!fs.existsSync(tmpDir)) {
    fs.mkdirSync(tmpDir, { recursive: true })
    console.log("created tmp directory")
}

export const defineAlias = (alias: string, dir: string) => ({
    find: alias,
    replacement: fileURLToPath(new URL(dir, import.meta.url)),
})

const common = {
    publicDir: "./app/public",
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
    publicDir: common.publicDir,
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
    publicDir: common.publicDir,
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
