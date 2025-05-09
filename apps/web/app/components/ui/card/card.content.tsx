import { cn } from "@/lib/utils"

export const CardContent = ({
    className,
    ...props
}: React.ComponentProps<"div">) => (
    <div
        data-slot="card-content"
        className={cn("p-3", className)}
        {...props}
    />
)
