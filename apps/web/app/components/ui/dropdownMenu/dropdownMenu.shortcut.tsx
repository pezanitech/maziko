import * as React from "react"

import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`text-muted-foreground ml-auto text-xs tracking-widest`,
}

export const DropdownMenuShortcut = ({
    className,
    ...props
}: React.ComponentProps<"span">) => (
    <span
        data-slot="dropdown-menu-shortcut"
        className={cn(styles.base, className)}
        {...props}
    />
)
