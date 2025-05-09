import * as React from "react"

import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu"

export const DropdownMenuTrigger = ({
    ...props
}: React.ComponentProps<typeof DropdownMenuPrimitive.Trigger>) => (
    <DropdownMenuPrimitive.Trigger
        data-slot="dropdown-menu-trigger"
        {...props}
    />
)
