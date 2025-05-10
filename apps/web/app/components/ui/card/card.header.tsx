import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`@container/card-header grid auto-rows-min grid-rows-[auto_auto] items-start gap-1.5 p-4 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-3`,
}

export const CardHeader = ({
    className,
    ...props
}: React.ComponentProps<"div">) => (
    <div
        data-slot="card-header"
        className={cn(styles.base, className)}
        {...props}
    />
)
