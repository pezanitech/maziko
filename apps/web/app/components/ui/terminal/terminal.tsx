import * as React from "react"

import clsx from "clsx"
import { Copy } from "lucide-react"
import { CodeBlock, vs2015 } from "react-code-blocks"

import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"

import { cn } from "@/lib/utils"

const styles = {
    wrapper: clsx`not-prose relative w-full`,
    card: clsx`rounded-lg border bg-[#191a20] backdrop-blur`,
    header: clsx`flex h-10 items-center justify-between rounded-t-lg bg-[#191a20] px-4`,
    dots: clsx`flex items-center gap-1.5`,
    dot: clsx`h-2.5 w-2.5 rounded-full`,
    label: clsx`text-muted-foreground ml-2 text-xs font-medium tracking-wide`,
    copyButton: clsx`text-muted-foreground hover:text-foreground`,
    cardContent: clsx`overflow-x-auto p-0 pb-3`,
    codeWrapper: clsx`px-4 [&_*]:!bg-transparent`,
    code: clsx`font-mono text-sm md:text-base`,
    successMessage: clsx`bg-success/20 text-success absolute top-2 right-2 rounded px-2 py-1 text-xs font-medium transition-opacity`,
}

export interface TerminalProps
    extends React.ComponentPropsWithoutRef<typeof Card> {
    content: string
    language?: string
    title?: string
}

export const Terminal = React.forwardRef<
    React.ComponentRef<typeof Card>,
    TerminalProps
>(
    (
        { content, language = "bash", title = "terminal", className, ...props },
        ref,
    ) => {
        const [copied, setCopied] = React.useState(false)

        const handleCopy = () => {
            navigator.clipboard.writeText(content)
            setCopied(true)
            setTimeout(() => setCopied(false), 2000)
        }

        return (
            <div className={styles.wrapper}>
                <Card
                    ref={ref}
                    className={cn(styles.card, className)}
                    {...props}
                >
                    <div className="bg-black">
                        <div className={styles.header}>
                            <div className={styles.dots}>
                                <div
                                    className={cn(
                                        styles.dot,
                                        "bg-destructive/70",
                                    )}
                                />
                                <div
                                    className={cn(styles.dot, "bg-warning/70")}
                                />
                                <div
                                    className={cn(styles.dot, "bg-success/70")}
                                />
                                <span className={styles.label}>{title}</span>
                            </div>
                            <Button
                                variant="ghost"
                                size="icon"
                                className={styles.copyButton}
                                onClick={handleCopy}
                                aria-label="Copy code"
                            >
                                <Copy className="h-4 w-4" />
                            </Button>
                        </div>

                        <CardContent className={styles.cardContent}>
                            <div className={styles.codeWrapper}>
                                <div className={styles.code}>
                                    <CodeBlock
                                        text={content}
                                        language={language}
                                        showLineNumbers={false}
                                        theme={vs2015}
                                        wrapLongLines={false}
                                    />
                                </div>
                            </div>
                        </CardContent>
                    </div>
                </Card>
                {copied && <div className={styles.successMessage}>Copied!</div>}
            </div>
        )
    },
)

Terminal.displayName = "Terminal"
