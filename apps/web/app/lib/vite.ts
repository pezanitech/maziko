import { fileURLToPath } from "url"

/**
 * Creates an alias definition for Vite's resolve.alias configuration.
 *
 * @param alias - The alias pattern to match in import statements (e.g., '@components')
 * @param dir - The directory path to resolve the alias to, relative to this file's location
 * @returns An alias object compatible with Vite's resolve.alias configuration
 *
 * @example
 * ```
 * const aliases = [
 *   defineAlias('@components', './src/components/'),
 *   defineAlias('@utils', './src/utils/')
 * ];
 * ```
 */
export const defineAlias = (alias: string, dir: string) => ({
    find: alias,
    replacement: fileURLToPath(new URL(dir, import.meta.url)),
})
