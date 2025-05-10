import clsx from "clsx"
import { BookOpen } from "lucide-react"

import { CircuitPattern } from "@/components/ui/circuitPattern"

const styles = {
    header: clsx`relative pb-6`,
    gradient: clsx`from-accent/5 via-background to-background absolute inset-0 -z-10 bg-gradient-to-b`,
    content: clsx`relative z-10`,
    titleWrapper: clsx`mb-4 flex items-center gap-3`,
    icon: clsx`bg-accent/10 text-accent inline-flex items-center justify-center rounded-lg p-2`,
    title: clsx`text-4xl font-bold`,
    description: clsx`text-muted-foreground text-lg`,
}

type DocsHeaderProps = {
    title: string
    description: string
}

export const DocsHeader = ({ title, description }: DocsHeaderProps) => (
    <header className={styles.header}>
        <div className={styles.gradient} />
        <CircuitPattern />
        <div className={styles.content}>
            <div className={styles.titleWrapper}>
                <div className={styles.icon}>
                    <BookOpen className="h-6 w-6" />
                </div>
                <h1 className={styles.title}>{title}</h1>
            </div>
            <p className={styles.description}>{description}</p>
        </div>
    </header>
)
