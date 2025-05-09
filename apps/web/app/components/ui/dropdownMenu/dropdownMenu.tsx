import * as React from "react"

import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu"

export const DropdownMenu = ({
    ...props
}: React.ComponentProps<typeof DropdownMenuPrimitive.Root>) => (
    <DropdownMenuPrimitive.Root
        data-slot="dropdown-menu"
        {...props}
    />
)
