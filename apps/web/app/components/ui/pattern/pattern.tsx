import clsx from "clsx"

import { cn } from "@/lib/utils"

const styles = {
    base: clsx`absolute inset-0 -z-10`,
    pattern: clsx`absolute inset-0 opacity-5 [mask-image:radial-gradient(100%_100%_at_top_center,black,transparent)]`,
}

type PatternProps = {
    className?: string
    strokeWidth?: number
    size?: number
    maskPosition?: 'top' | 'center' | 'bottom'
}

export const Pattern = ({
    className,
    strokeWidth = 1,
    size = 32,
    maskPosition = 'top',
}: PatternProps) => (
    <div className={cn(styles.base, className)}>
        <div className={cn(
            styles.pattern,
            maskPosition === 'center' && '[mask-image:radial-gradient(100%_100%_at_center,black,transparent)]',
            maskPosition === 'bottom' && '[mask-image:radial-gradient(100%_100%_at_bottom_center,black,transparent)]'
        )}>
            <svg className="absolute inset-0 h-full w-full" xmlns="http://www.w3.org/2000/svg">
                <defs>
                    <pattern
                        id="grid"
                        width={size}
                        height={size}
                        patternUnits="userSpaceOnUse"
                    >
                        <path
                            d={`M${size} ${size}V.5H.5`}
                            fill="none"
                            stroke="currentColor"
                            strokeWidth={strokeWidth}
                        />
                    </pattern>
                </defs>
                <rect width="100%" height="100%" fill="url(#grid)" />
            </svg>
        </div>
    </div>
)