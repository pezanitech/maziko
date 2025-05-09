import { cn } from "@/lib/utils"

export const CardAction = ({
    className,
    ...props
}: React.ComponentProps<"div">) => {
    return (
        <div
            data-slot="card-action"
            className={cn(
                "col-start-2 row-span-2 row-start-1 self-start justify-self-end",
                className,
            )}
            {...props}
        />
    )
}
