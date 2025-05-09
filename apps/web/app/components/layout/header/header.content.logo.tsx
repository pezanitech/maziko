import clsx from "clsx"

import { MazikoLogo } from "@/components/ui/icons"
import { Link } from "@/components/ui/link"

import { brand } from "./resources/data"

const styles = {
    wrapper: clsx`text-foreground flex items-center gap-2`,
    logo: clsx`h-8 w-8`,
    name: clsx`hidden text-2xl font-semibold sm:block`,
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
