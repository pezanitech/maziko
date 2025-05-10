import clsx from "clsx"
import {
    Folder,
    HelpCircle,
    Rocket,
    Server,
    Sparkles,
    Wrench,
    Zap,
} from "lucide-react"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`flex items-center justify-center transition-all duration-200`,
    sizes: {
        sm: clsx`h-8 w-8`,
        md: clsx`h-10 w-10`,
        lg: clsx`h-12 w-12`,
    },
    variants: {
        default: clsx`bg-accent/10 text-accent hover:bg-accent/20 rounded-lg`,
        solid: clsx`bg-accent text-accent-foreground hover:bg-accent/90 rounded-lg`,
        outline: clsx`border-accent/20 text-accent hover:border-accent/40 rounded-lg border-2`,
        ghost: clsx`text-accent hover:bg-accent/10 rounded-xl`,
    },
}

const iconMap = {
    rocket: Rocket,
    sparkles: Sparkles,
    folder: Folder,
    server: Server,
    bolt: Zap,
    wrench: Wrench,
} as const

type FeatureIconProps = {
    icon: string
    variant?: keyof typeof styles.variants
    iconSize?: number
    size?: keyof typeof styles.sizes
    className?: string
}

export const FeatureIcon = (props: FeatureIconProps) => {
    const Icon = iconMap[props.icon as keyof typeof iconMap] || HelpCircle

    return (
        <div
            className={cn(
                styles.base,
                styles.sizes[props.size ?? "md"],
                styles.variants[props.variant ?? "default"],
                props.className,
            )}
        >
            <Icon size={props.iconSize} />
        </div>
    )
}
