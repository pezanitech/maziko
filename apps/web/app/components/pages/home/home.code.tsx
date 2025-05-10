import { useState } from "react"

import clsx from "clsx"
import { Code, Terminal as TerminalIcon } from "lucide-react"

import { Separator } from "@/components/ui/separator"
import { Terminal } from "@/components/ui/terminal"

import { Container } from "@/components/layout/container"

import { CodeExample } from "./home"

const styles = {
    wrapper: clsx`py-8`,
    grid: clsx`grid grid-cols-1 items-start gap-8 lg:grid-cols-12`,
    topics: clsx`hidden space-y-4 lg:col-span-4 lg:block xl:col-span-3`,
    topicTitle: clsx`text-foreground mb-3 text-2xl font-bold`,
    terminal: clsx`lg:col-span-8 xl:col-span-9`,
    terminalWrapper: clsx`not-prose`,
    topic: clsx`relative cursor-pointer overflow-hidden rounded-lg border p-4 transition-all duration-200`,
    topicActive: clsx`bg-accent/10 border-accent shadow-sm`,
    topicInactive: clsx`bg-secondary/20 hover:bg-accent/5 border-transparent`,
    topicHeader: clsx`mb-2 flex items-center justify-between`,
    topicName: clsx`text-foreground text-sm font-semibold`,
    separator: clsx`mx-auto my-12 max-w-5xl`,
    codeSection: clsx`space-y-8`,
    stepTitle: clsx`mb-1 text-2xl font-semibold`,
    stepDescription: clsx`text-muted-foreground mb-4`,
    sidebarSeparator: clsx`my-4`,
    topicDescription: clsx`text-muted-foreground text-sm leading-relaxed`,
    topicIconWrapper: clsx`mb-2 flex items-center gap-2`,
    topicIconContainer: clsx`flex h-6 w-6 items-center justify-center rounded-full`,
    tabBar: clsx`scrollbar-none mb-5 flex gap-2 overflow-x-auto pb-3 lg:hidden`,
    tab: clsx`flex-none cursor-pointer px-2 py-2 text-sm whitespace-nowrap transition-all duration-200`,
    tabActive: clsx`text-foreground border-accent hover:bg-accent/5 border-b-2 font-medium`,
    tabInactive: clsx`text-muted-foreground hover:bg-secondary/20 border-b-2`,
    tabIcon: clsx`mr-2 inline-block h-4 w-4`,
}

type CodeSectionProps = {
    codeExamples?: CodeExample[]
}

export const CodeSection = ({ codeExamples }: CodeSectionProps) => {
    // Return null if no code examples are provided
    if (!codeExamples || codeExamples.length === 0) {
        return null
    }

    const [activeIndex, setActiveIndex] = useState(0)

    // Check if current example has steps (like Getting Started)
    const hasSteps =
        codeExamples[activeIndex].steps &&
        codeExamples[activeIndex].steps!.length > 0

    // Get the title for the terminal based on type
    const getTerminalTitle = (
        example: CodeExample | NonNullable<CodeExample["steps"]>[number],
    ) => {
        if (example.type === "code" && example.filename) {
            return example.filename
        }
        return "terminal"
    }

    // Get the language for syntax highlighting based on type
    const getLanguage = (
        example: CodeExample | NonNullable<CodeExample["steps"]>[number],
    ) => {
        if (example.type === "code") {
            if (example.filename) {
                const extension = example.filename
                    .split(".")
                    .pop()
                    ?.toLowerCase()

                // Map file extensions to languages
                switch (extension) {
                    case "go":
                        return "go"
                    case "js":
                        return "javascript"
                    case "ts":
                        return "typescript"
                    case "tsx":
                        return "tsx"
                    case "jsx":
                        return "jsx"
                    case "html":
                        return "html"
                    case "css":
                        return "css"
                    default:
                        return "text"
                }
            }
            return "text"
        }
        return "bash" // For shell commands
    }

    return (
        <>
            <Container className={styles.wrapper}>
                {/* Scrollable Tab Bar (Mobile) */}
                <div className={styles.tabBar}>
                    {codeExamples.map((example, index) => (
                        <button
                            key={index}
                            className={`${styles.tab} ${
                                index === activeIndex
                                    ? styles.tabActive
                                    : styles.tabInactive
                            }`}
                            onClick={() => setActiveIndex(index)}
                        >
                            {example.type === "code" ? (
                                <Code
                                    size={14}
                                    className={styles.tabIcon}
                                />
                            ) : (
                                <TerminalIcon
                                    size={14}
                                    className={styles.tabIcon}
                                />
                            )}
                            {example.name}
                        </button>
                    ))}
                </div>

                <div className={styles.grid}>
                    <div className={styles.topics}>
                        <h3 className={styles.topicTitle}>Code Examples</h3>
                        <p className={styles.topicDescription}>
                            Explore some examples to get started with Maziko.
                        </p>

                        <Separator className={styles.sidebarSeparator} />

                        {codeExamples.map((example, index) => (
                            <div
                                key={index}
                                className={`${styles.topic} ${
                                    index === activeIndex
                                        ? styles.topicActive
                                        : styles.topicInactive
                                }`}
                                onMouseEnter={() => setActiveIndex(index)}
                                onClick={() => setActiveIndex(index)}
                            >
                                <div className={styles.topicHeader}>
                                    <div className={styles.topicIconWrapper}>
                                        <div
                                            className={
                                                styles.topicIconContainer
                                            }
                                            style={{
                                                backgroundColor:
                                                    index === activeIndex
                                                        ? "rgba(var(--accent), 0.2)"
                                                        : "rgba(var(--secondary), 0.2)",
                                            }}
                                        >
                                            {example.type === "code" ? (
                                                <Code
                                                    size={14}
                                                    className="text-accent"
                                                />
                                            ) : (
                                                <TerminalIcon
                                                    size={14}
                                                    className="text-accent"
                                                />
                                            )}
                                        </div>
                                        <span className={styles.topicName}>
                                            {example.name}
                                        </span>
                                    </div>
                                </div>
                                <span className={styles.topicDescription}>
                                    {example.description}
                                </span>
                            </div>
                        ))}
                    </div>

                    <div className={styles.terminal}>
                        {hasSteps ? (
                            <div className={styles.codeSection}>
                                {codeExamples[activeIndex].steps!.map(
                                    (step, index) => (
                                        <div key={index}>
                                            <h4 className={styles.stepTitle}>
                                                {step.title}
                                            </h4>
                                            <p
                                                className={
                                                    styles.stepDescription
                                                }
                                            >
                                                {step.description}
                                            </p>
                                            <div
                                                className={
                                                    styles.terminalWrapper
                                                }
                                            >
                                                <Terminal
                                                    content={step.code}
                                                    className="w-full"
                                                    title={getTerminalTitle(
                                                        step,
                                                    )}
                                                    language={getLanguage(step)}
                                                />
                                            </div>
                                        </div>
                                    ),
                                )}
                            </div>
                        ) : (
                            <div className={styles.codeSection}>
                                <div>
                                    <h4 className={styles.stepTitle}>
                                        {codeExamples[activeIndex].name}
                                    </h4>
                                    <p className={styles.stepDescription}>
                                        {codeExamples[activeIndex].description}
                                    </p>
                                    <div className={styles.terminalWrapper}>
                                        <Terminal
                                            content={
                                                codeExamples[activeIndex].code
                                            }
                                            className="w-full"
                                            title={getTerminalTitle(
                                                codeExamples[activeIndex],
                                            )}
                                            language={getLanguage(
                                                codeExamples[activeIndex],
                                            )}
                                        />
                                    </div>
                                </div>
                            </div>
                        )}
                    </div>
                </div>
            </Container>
        </>
    )
}
