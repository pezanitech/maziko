import * as React from "react"

import { Slot } from "@radix-ui/react-slot"
import type { VariantProps } from "class-variance-authority"
import { cva } from "class-variance-authority"
import clsx from "clsx"

import { cn } from "@/lib/utils"

export interface ButtonProps
    extends React.ComponentProps<"button">,
        VariantProps<typeof buttonVariants> {
    asChild?: boolean
}

export const buttonVariants = cva(
    clsx`focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive inline-flex shrink-0 cursor-pointer items-center justify-center gap-2 rounded-xl text-sm font-medium whitespace-nowrap transition-all duration-300 outline-none focus-visible:ring-[3px] disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4`,
    {
        defaultVariants: {
            variant: "default",
            size: "default",
        },

        variants: {
            variant: {
                default: clsx`bg-foreground text-background hover:bg-accent/90 hover:text-foreground shadow-xs`,

                accent: clsx`bg-accent text-accent-foreground hover:bg-accent/90 shadow-xs`,

                destructive: clsx`bg-destructive hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60 text-white shadow-xs`,

                success: clsx`bg-success hover:bg-success/90 focus-visible:ring-success/20 dark:focus-visible:ring-success/40 dark:bg-success/60 text-white shadow-xs`,

                outline: clsx`dark:border-input hover:dark:border-accent hover:border-accent hover:text-accent-foreground dark:hover:text-accent-foreground dark:hover:bg-accent hover:bg-accent border bg-transparent shadow-xs`,

                secondary: clsx`bg-secondary text-secondary-foreground hover:bg-secondary/80 shadow-xs`,

                ghost: clsx`hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50`,

                link: clsx`text-primary underline-offset-4 hover:underline`,
            },
            size: {
                default: clsx`h-9 px-4 py-2 has-[>svg]:px-3`,
                sm: clsx`h-8 gap-1.5 rounded-md px-3 has-[>svg]:px-2.5`,
                lg: clsx`h-10 rounded-md px-6 has-[>svg]:px-4`,
                icon: clsx`size-9`,
            },
        },
    },
)

export function Button({
    className,
    variant,
    size,
    asChild = false,
    ...props
}: ButtonProps) {
    const Comp = asChild ? Slot : "button"

    return (
        <Comp
            data-slot="button"
            className={cn(buttonVariants({ variant, size, className }))}
            {...props}
        />
    )
}
