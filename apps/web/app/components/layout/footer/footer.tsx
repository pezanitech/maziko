import clsx from "clsx"

import { cn } from "@/lib/utils"
import { sharedStyles } from "@/styles/sharedStyles"

import { FooterLeft } from "./footer.left"
import { FooterRight } from "./footer.right"

const styles = {
    footer: clsx`border-border/40 border-t`,

    container: cn(
        clsx`flex flex-col items-center justify-between gap-8 py-12 md:flex-row`,
        sharedStyles.container,
    ),
}

export const Footer = () => (
    <footer className={styles.footer}>
        <div className={styles.container}>
            <FooterLeft />
            <FooterRight />
        </div>
    </footer>
)
