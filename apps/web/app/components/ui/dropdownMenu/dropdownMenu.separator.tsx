import * as React from "react"

import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu"
import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`bg-border -mx-1 my-1 h-px`,
}

export const DropdownMenuSeparator = ({
    className,
    ...props
}: React.ComponentProps<typeof DropdownMenuPrimitive.Separator>) => (
    <DropdownMenuPrimitive.Separator
        data-slot="dropdown-menu-separator"
        className={cn(styles.base, className)}
        {...props}
    />
)
