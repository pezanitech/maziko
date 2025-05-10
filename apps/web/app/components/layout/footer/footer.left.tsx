import clsx from "clsx"

import { MazikoLogo } from "@/components/ui/icons"

import { cn } from "@/lib/utils"

import { brand } from "./resources/data"

const styles = {
    wrapper: clsx`flex flex-col items-center gap-4 md:flex-row md:items-center`,
    logo: clsx`h-14 w-14`,
    logoWrapper: clsx`flex flex-col items-center gap-1 md:hidden`,
    brandName: clsx`text-2xl font-bold`,
    description: clsx`text-muted-foreground text-center text-sm md:text-left`,
}

export const FooterLeft = () => (
    <div className={styles.wrapper}>
        <div className={styles.logoWrapper}>
            <MazikoLogo className={styles.logo} />
            <span className={styles.brandName}>Maziko</span>
        </div>
        <MazikoLogo className={cn(styles.logo, "hidden md:block")} />
        <p className={styles.description}>{brand.copyright}</p>
    </div>
)
