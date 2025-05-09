import clsx from "clsx"

import { Link } from "@/components/ui/link"

import { links } from "./resources/data"

const styles = {
    wrapper: clsx`flex flex-col items-center gap-4 md:items-end`,

    links: clsx`text-muted-foreground flex items-center gap-4 text-sm`,

    link: clsx`hover:text-accent transition-colors`,
}

export const FooterRight = () => (
    <div className={styles.wrapper}>
        <div className={styles.links}>
            {links.map((link) => (
                <Link
                    key={link.name}
                    href={link.href}
                    className={styles.link}
                >
                    {link.name}
                </Link>
            ))}
        </div>
    </div>
)
