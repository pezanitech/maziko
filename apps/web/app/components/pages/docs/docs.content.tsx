import clsx from "clsx"
import ReactMarkdown from "react-markdown"

import { Separator } from "@/components/ui/separator"

import { type Section } from "./docs"
import { CodeBlockComponent } from "./docs.blocks"

const styles = {
    section: clsx`prose prose-zinc dark:prose-invert max-w-3xl`,
}

const CodeBlock = ({ node, inline, className, children, ...props }: any) => {
    const match = /language-(\w+)/.exec(className || "")
    const code = String(children).replace(/\n$/, "")

    if (!inline && match) {
        return (
            <div className="my-6">
                <CodeBlockComponent
                    code={code}
                    language={match[1]}
                />
            </div>
        )
    }

    return (
        <code
            className={className}
            {...props}
        >
            {children}
        </code>
    )
}

const components = {
    code: CodeBlock,
    pre: ({
        children,
        className,
        style,
    }: React.HTMLAttributes<HTMLPreElement>) => {
        return (
            <div
                className={className}
                style={style}
            >
                {children}
            </div>
        )
    },
}

const processMarkdown = (content: string) => {
    // Remove trailing backticks after code blocks
    return content.replace(/```\n*$/g, "\n```")
}

export const DocsContent = ({ sections }: { sections: Section[] }) => (
    <div className="max-w-3xl">
        {sections.map((section, index) => (
            <div key={section.id}>
                <section
                    id={section.id}
                    className={styles.section}
                >
                    <h2>{section.title}</h2>
                    <ReactMarkdown>
                        {processMarkdown(section.content)}
                    </ReactMarkdown>

                    {section.subsections.map((subsection) => (
                        <div
                            key={subsection.id}
                            id={subsection.id}
                        >
                            <h2>{subsection.title}</h2>
                            <ReactMarkdown components={components}>
                                {processMarkdown(subsection.content)}
                            </ReactMarkdown>
                        </div>
                    ))}
                </section>
                {index < sections.length - 1 && (
                    <Separator className="my-8 opacity-20" />
                )}
            </div>
        ))}
    </div>
)
