import * as React from "react"

import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu"
import clsx from "clsx"
import { ChevronRightIcon } from "lucide-react"

import { cn } from "@/lib/utils"

const styles = {
    trigger: clsx`focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground flex cursor-default items-center rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[inset]:pl-8`,

    content: clsx`bg-popover text-popover-foreground data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 z-50 min-w-[8rem] origin-(--radix-dropdown-menu-content-transform-origin) overflow-hidden rounded-md border p-1 shadow-lg`,
}

export const DropdownMenuSub = ({
    ...props
}: React.ComponentProps<typeof DropdownMenuPrimitive.Sub>) => (
    <DropdownMenuPrimitive.Sub
        data-slot="dropdown-menu-sub"
        {...props}
    />
)

export const DropdownMenuSubTrigger = ({
    className,
    inset,
    children,
    ...props
}: React.ComponentProps<typeof DropdownMenuPrimitive.SubTrigger> & {
    inset?: boolean
}) => (
    <DropdownMenuPrimitive.SubTrigger
        data-slot="dropdown-menu-sub-trigger"
        data-inset={inset}
        className={cn(styles.trigger, className)}
        {...props}
    >
        {children}
        <ChevronRightIcon className="ml-auto size-4" />
    </DropdownMenuPrimitive.SubTrigger>
)

export const DropdownMenuSubContent = ({
    className,
    ...props
}: React.ComponentProps<typeof DropdownMenuPrimitive.SubContent>) => (
    <DropdownMenuPrimitive.SubContent
        data-slot="dropdown-menu-sub-content"
        className={cn(styles.content, className)}
        {...props}
    />
)
