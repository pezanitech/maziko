import clsx from "clsx"
import { Moon, Sun } from "lucide-react"

import { Button } from "@/components/ui/button"

import { useTheme } from "@/components/providers/themeProvider"
import { cn } from "@/lib/utils"

const styles = {
    icon: clsx`h-[1.2rem] w-[1.2rem]`,

    sunIcon: clsx`scale-100 rotate-0 dark:scale-0 dark:-rotate-90`,

    moonIcon: clsx`absolute scale-0 rotate-90 dark:scale-100 dark:rotate-0`,
}

export function ModeToggle() {
    const { theme, setTheme } = useTheme()

    return (
        <Button
            variant="outline"
            size="icon"
            onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
        >
            <Sun className={cn(styles.icon, styles.sunIcon)} />
            <Moon className={cn(styles.icon, styles.moonIcon)} />
            <span className="sr-only">Toggle theme</span>
        </Button>
    )
}
