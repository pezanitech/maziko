import clsx from "clsx"
import { Copy } from "lucide-react"
import { CodeBlock, monokai } from "react-code-blocks"

import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"

import { cn } from "@/lib/utils"

const styles = {
    wrapper: clsx`not-prose relative`,
    card: clsx`rounded-lg border bg-[#191a20]`,
    header: clsx`flex h-10 items-center justify-between rounded-t-lg px-4`,
    dots: clsx`flex items-center gap-1.5`,
    dot: clsx`h-2.5 w-2.5 rounded-full`,
    label: clsx`text-muted-foreground ml-2 text-xs font-medium`,
    copyButton: clsx`text-muted-foreground hover:text-foreground`,
    cardContent: clsx`overflow-x-auto p-0`,
    codeWrapper: clsx`px-2 pb-2 [&_*]:!bg-transparent`,
    code: clsx`font-mono text-base`,
}

type CodeBlockComponentProps = {
    code: string
    language: string
    title?: string
}

export const CodeBlockComponent = ({
    code,
    language,
    title,
}: CodeBlockComponentProps) => {
    const handleCopy = () => {
        navigator.clipboard.writeText(code)
    }

    return (
        <div className={styles.wrapper}>
            <Card className={styles.card}>
                <div className={styles.header}>
                    <div className={styles.dots}>
                        <div className={cn(styles.dot, "bg-destructive/70")} />
                        <div className={cn(styles.dot, "bg-warning/70")} />
                        <div className={cn(styles.dot, "bg-success/70")} />
                        {title && <span className={styles.label}>{title}</span>}
                    </div>
                    <Button
                        variant="ghost"
                        size="icon"
                        className={styles.copyButton}
                        onClick={handleCopy}
                    >
                        <Copy className="h-4 w-4" />
                    </Button>
                </div>

                <CardContent className={styles.cardContent}>
                    <div className={styles.codeWrapper}>
                        <div className={styles.code}>
                            <CodeBlock
                                text={code}
                                language={language}
                                showLineNumbers={true}
                                theme={monokai}
                                wrapLongLines={false}
                            />
                        </div>
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}
