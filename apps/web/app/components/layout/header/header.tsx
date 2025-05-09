"use client"

import { useState } from "react"

import clsx from "clsx"

import { sharedStyles } from "@/styles/sharedStyles"

import { HeaderContent } from "./header.content"
import { HeaderMobileNav } from "./header.mobileNav"

const styles = {
    header: clsx`border-border/40 bg-card sticky top-0 z-50 min-h-16 w-full border-b`,

    container: sharedStyles.container,
}

export function Header() {
    const [isOpen, setIsOpen] = useState(false)

    return (
        <header className={styles.header}>
            <div className={styles.container}>
                <HeaderContent
                    isOpen={isOpen}
                    onOpenChange={setIsOpen}
                />
                <HeaderMobileNav
                    isOpen={isOpen}
                    onOpenChange={setIsOpen}
                />
            </div>
        </header>
    )
}
