import { cn } from "@/lib/utils"

export const CardFooter = ({
    className,
    ...props
}: React.ComponentProps<"div">) => (
    <div
        data-slot="card-footer"
        className={cn(
            "flex items-center px-3 pb-3 [.border-t]:pt-3",
            className,
        )}
        {...props}
    />
)
