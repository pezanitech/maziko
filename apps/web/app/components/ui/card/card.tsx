import * as React from "react"

import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    card: clsx`bg-card text-card-foreground ring-accent flex flex-col gap-3 overflow-hidden rounded-xl border transition-all`,
}

export const Card = ({
    className,
    hoverRing,
    ...props
}: React.ComponentProps<"div"> & { hoverRing?: boolean }) => {
    return (
        <div
            data-slot="card"
            className={cn(
                styles.card,
                className,
                hoverRing ? "hover:ring-2" : "",
            )}
            {...props}
        />
    )
}
