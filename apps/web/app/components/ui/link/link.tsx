import * as React from "react"

import { Link as InertiaLink } from "@inertiajs/react"
import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`focus-visible:ring-ring/50 inline-flex items-center justify-center rounded-3xl px-1 outline-none focus-visible:ring-[3px]`,
}

export type LinkProps = React.ComponentPropsWithoutRef<typeof InertiaLink>

export const Link = React.forwardRef<HTMLAnchorElement, LinkProps>(
    ({ className, ...props }, ref) => (
        <InertiaLink
            className={cn(styles.base, className)}
            ref={ref}
            {...props}
        />
    ),
)
Link.displayName = "Link"
