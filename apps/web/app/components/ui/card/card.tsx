import * as React from "react"

import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    card: clsx`text-card-foreground bg-opacity-50 border-border bg-card/50 hover:border-accent/40 hover:bg-accent/10 group flex flex-col gap-3 overflow-hidden rounded-xl border shadow transition-colors`,
}

export const Card = ({
    className,
    ...props
}: React.ComponentProps<"div"> & { hoverRing?: boolean }) => {
    return (
        <div
            data-slot="card"
            className={cn(styles.card, className)}
            {...props}
        />
    )
}
