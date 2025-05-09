import { cn } from "@/lib/utils"

export const CardTitle = ({
    className,
    ...props
}: React.ComponentProps<"div">) => {
    return (
        <div
            data-slot="card-title"
            className={cn("leading-none font-semibold", className)}
            {...props}
        />
    )
}
