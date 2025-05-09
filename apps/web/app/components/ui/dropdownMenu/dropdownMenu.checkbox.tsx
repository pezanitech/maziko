import * as React from "react"

import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu"
import clsx from "clsx"
import { CheckIcon } from "lucide-react"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`focus:bg-accent focus:text-accent-foreground relative flex cursor-default items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4`,

    indicator: clsx`pointer-events-none absolute left-2 flex size-3.5 items-center justify-center`,
}

export const DropdownMenuCheckboxItem = ({
    className,
    children,
    checked,
    ...props
}: React.ComponentProps<typeof DropdownMenuPrimitive.CheckboxItem>) => (
    <DropdownMenuPrimitive.CheckboxItem
        data-slot="dropdown-menu-checkbox-item"
        className={cn(styles.base, className)}
        checked={checked}
        {...props}
    >
        <span className={styles.indicator}>
            <DropdownMenuPrimitive.ItemIndicator>
                <CheckIcon className="size-4" />
            </DropdownMenuPrimitive.ItemIndicator>
        </span>
        {children}
    </DropdownMenuPrimitive.CheckboxItem>
)
