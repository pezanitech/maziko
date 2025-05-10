import clsx from "clsx"

import { MazikoLogo } from "@/components/ui/icons"
import { Link } from "@/components/ui/link"

import { brand } from "./resources/data"

const styles = {
    wrapper: clsx`text-foreground hover:text-accent flex items-center gap-2 transition-colors duration-300`,
    logo: clsx`h-8 w-8`,
    name: clsx`text-2xl font-semibold`,
}

export const HeaderContentLogo = () => (
    <Link
        href={brand.href}
        className={styles.wrapper}
    >
        <MazikoLogo className={styles.logo} />
        <span className={styles.name}>Maziko</span>
    </Link>
)
