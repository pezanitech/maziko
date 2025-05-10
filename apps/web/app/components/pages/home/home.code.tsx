import clsx from "clsx"
import { Copy } from "lucide-react"
import { CodeBlock, vs2015 } from "react-code-blocks"

import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"

import { Container } from "@/components/layout/container"
import { cn } from "@/lib/utils"

const styles = {
    wrapper: clsx`not-prose relative`,
    card: clsx`mx-auto max-w-3xl rounded-lg border bg-[#191a20] backdrop-blur`,
    header: clsx`flex h-10 items-center justify-between rounded-t-lg bg-[#191a20] px-4`,
    dots: clsx`flex items-center gap-1.5`,
    dot: clsx`h-2.5 w-2.5 rounded-full`,
    label: clsx`text-muted-foreground ml-2 text-xs font-medium`,
    copyButton: clsx`text-muted-foreground hover:text-foreground`,
    cardContent: clsx`overflow-x-auto`,
    codeWrapper: clsx`[&_*]:!bg-transparent`,
    code: clsx`font-mono text-base`,
}

type CodeSectionProps = {
    codeExample: string
}

export const CodeSection = (props: CodeSectionProps) => {
    const handleCopy = () => {
        navigator.clipboard.writeText(props.codeExample)
    }

    return (
        <Container>
            <div className={styles.wrapper}>
                <Card className={styles.card}>
                    <div className={styles.header}>
                        <div className={styles.dots}>
                            <div
                                className={cn(styles.dot, "bg-destructive/70")}
                            />
                            <div className={cn(styles.dot, "bg-warning/70")} />
                            <div className={cn(styles.dot, "bg-success/70")} />
                            <span className={styles.label}>Terminal</span>
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
                                    text={props.codeExample}
                                    language="bash"
                                    showLineNumbers={false}
                                    theme={vs2015}
                                    wrapLongLines={false}
                                />
                            </div>
                        </div>
                    </CardContent>
                </Card>
            </div>
        </Container>
    )
}
