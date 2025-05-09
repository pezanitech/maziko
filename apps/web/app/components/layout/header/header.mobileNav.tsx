import clsx from "clsx"

import { cn } from "@/lib/utils"

import { Link } from "../../ui/link"
import { navigation } from "./resources/data"

const styles = {
    mobileNav: clsx`border-border/40 overflow-hidden border-b bg-transparent transition-all md:hidden`,

    mobileLinks: clsx`flex flex-col space-y-4 py-6`,

    mobileLink: clsx`text-muted-foreground hover:text-accent text-base font-medium transition-colors duration-300`,
}

interface HeaderMobileNavProps {
    isOpen: boolean
    onOpenChange: (open: boolean) => void
}

export const HeaderMobileNav = ({
    isOpen,
    onOpenChange,
}: HeaderMobileNavProps) => (
    <div className={cn(styles.mobileNav, isOpen ? "h-40" : "h-0")}>
        <div className={styles.mobileLinks}>
            {navigation.map((item) => (
                <Link
                    key={item.name}
                    href={item.href}
                    className={styles.mobileLink}
                    onClick={() => onOpenChange(false)}
                >
                    {item.name}
                </Link>
            ))}
        </div>
    </div>
)
