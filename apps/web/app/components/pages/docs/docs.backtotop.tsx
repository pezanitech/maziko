import clsx from "clsx"
import { Button } from "@/components/ui/button"
import { ChevronUp } from "lucide-react"

const styles = {
    button: clsx`fixed bottom-8 right-8 opacity-0 transition-opacity duration-200`,
    visible: clsx`opacity-100`,
}

type BackToTopProps = {
    show: boolean
}

export const BackToTop = ({ show }: BackToTopProps) => (
    <Button
        variant="secondary"
        size="icon"
        className={clsx(styles.button, show && styles.visible)}
        onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}
    >
        <ChevronUp className="h-4 w-4" />
    </Button>
)