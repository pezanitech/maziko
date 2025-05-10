import { cn } from "@/lib/utils"

export const Container = ({
    className,
    ...props
}: React.ComponentProps<"div">) => (
    <div
        className={cn("container mx-auto px-4 py-8 md:px-6", className)}
        {...props}
    />
)
