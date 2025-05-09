import * as React from "react"

import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu"
import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`px-2 py-1.5 text-sm font-medium data-[inset]:pl-8`,
}

export const DropdownMenuLabel = ({
    className,
    inset,
    ...props
}: React.ComponentProps<typeof DropdownMenuPrimitive.Label> & {
    inset?: boolean
}) => (
    <DropdownMenuPrimitive.Label
        data-slot="dropdown-menu-label"
        data-inset={inset}
        className={cn(styles.base, className)}
        {...props}
    />
)
