import clsx from "clsx"

import { Link } from "@/components/ui/link"

import { navigation } from "./resources/data"

const styles = {
    nav: clsx`hidden flex-1 items-center space-x-8 md:flex`,

    link: clsx`text-foreground hover:text-accent text-xl font-bold transition-colors duration-300`,
}

export const HeaderContentNav = () => (
    <nav className={styles.nav}>
        {navigation.map((item) => (
            <Link
                key={item.name}
                href={item.href}
                className={styles.link}
            >
                {item.name}
            </Link>
        ))}
    </nav>
)
