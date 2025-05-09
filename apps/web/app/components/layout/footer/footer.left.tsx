import clsx from "clsx"

import { MazikoLogo } from "@/components/ui/icons"

import { brand } from "./resources/data"

const styles = {
    wrapper: clsx`flex flex-col items-center gap-4 md:flex-row md:items-center`,

    logo: clsx`h-14 w-14`,

    description: clsx`text-muted-foreground text-center text-sm md:text-left`,
}

export const FooterLeft = () => (
    <div className={styles.wrapper}>
        <MazikoLogo className={styles.logo} />
        <p className={styles.description}>{brand.copyright}</p>
    </div>
)
