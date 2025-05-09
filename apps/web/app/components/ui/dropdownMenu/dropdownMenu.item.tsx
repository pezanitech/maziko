import * as React from "react"

import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu"
import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`focus:bg-accent focus:text-accent-foreground relative flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4`,

    variants: {
        destructive: clsx`data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 dark:data-[variant=destructive]:focus:bg-destructive/20 data-[variant=destructive]:focus:text-destructive data-[variant=destructive]:*:[svg]:!text-destructive`,
    },
}

export interface DropdownMenuItemProps
    extends React.ComponentProps<typeof DropdownMenuPrimitive.Item> {
    inset?: boolean
    variant?: "default" | "destructive"
}

export const DropdownMenuItem = ({
    className,
    inset,
    variant = "default",
    ...props
}: DropdownMenuItemProps) => (
    <DropdownMenuPrimitive.Item
        data-slot="dropdown-menu-item"
        data-inset={inset}
        data-variant={variant}
        className={cn(
            styles.base,
            variant === "destructive" && styles.variants.destructive,
            className,
        )}
        {...props}
    />
)
