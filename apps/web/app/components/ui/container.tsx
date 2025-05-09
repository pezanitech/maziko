import React from "react"

import { cn } from "@/lib/utils"

interface ContainerProps extends React.HTMLAttributes<HTMLDivElement> {
    children: React.ReactNode
}

export const Container = React.forwardRef<HTMLDivElement, ContainerProps>(
    ({ className, children, ...props }, ref) => {
        return (
            <div
                ref={ref}
                className={cn("container mx-auto px-4 md:px-6", className)}
                {...props}
            >
                {children}
            </div>
        )
    },
)

Container.displayName = "Container"
