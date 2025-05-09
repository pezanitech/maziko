// Documentation Page
import { useEffect, useState } from "react"

import { usePage } from "@inertiajs/react"
import { CodeBlock, nord } from "react-code-blocks"

type Subsection = {
    title: string
    id: string
    content: string
}

type Section = {
    title: string
    id: string
    content: string
    subsections: Subsection[]
}

type DocProps = {
    title: string
    description: string
    sections: Section[]
}

// Styles for the documentation page
const styles = {
    // Layout
    container: "container mx-auto px-4 py-8",
    flexLayout: "flex flex-col lg:flex-row gap-8",
    contentArea: "lg:flex-1 min-w-0",

    // Navigation
    sidebar: "lg:w-64 flex-shrink-0 pb-8",
    sidebarSticky: "lg:sticky lg:top-8",
    navSection: "mb-4",
    navTitle: "text-lg font-semibold mb-2 text-emerald-500",
    navList: "space-y-1",
    navItem: "block py-1 px-2 rounded-lg hover:bg-gray-800 transition-colors",
    navItemActive: "bg-emerald-500 bg-opacity-20 text-emerald-400",

    // Main content
    docHeader: "mb-8",
    docTitle: "text-4xl font-bold mb-3 pb-3 border-b border-gray-800",
    docDescription: "text-gray-400 text-xl",

    // Sections
    section: "mb-16 scroll-mt-6",
    sectionHeader: "mb-6",
    sectionTitle: "text-3xl font-bold mb-3",
    sectionDescription: "text-gray-300 text-lg",

    // Subsections
    subsection: "mb-12 scroll-mt-6",
    subsectionTitle: "text-2xl font-semibold mb-4",
    prose: "prose prose-invert prose-emerald max-w-none prose-pre:bg-gray-800 prose-pre:border prose-pre:border-gray-700",

    // Code blocks
    codeBlock: "my-6 rounded-lg overflow-hidden font-mono",

    // Back to top
    backToTop:
        "fixed bottom-4 right-4 bg-emerald-500 hover:bg-emerald-600 text-black p-3 rounded-full shadow-lg transition-colors",

    // Mobile menu
    mobileMenuButton:
        "lg:hidden mb-4 flex items-center gap-2 text-emerald-500 font-medium",
    mobileMenuIcon: "w-5 h-5",
}

export default function Page() {
    const { props } = usePage<DocProps>()
    const [activeSection, setActiveSection] = useState("")
    const [showBackToTop, setShowBackToTop] = useState(false)
    const [mobileMenuOpen, setMobileMenuOpen] = useState(false)

    // Handle scroll event to update active section and show/hide back to top button
    useEffect(() => {
        const handleScroll = () => {
            // Show back to top button when scrolled down
            setShowBackToTop(window.scrollY > 500)

            // Update active section based on scroll position
            const sectionElements =
                document.querySelectorAll('[id^="section-"]')
            let currentSection = ""

            sectionElements.forEach((section) => {
                const rect = section.getBoundingClientRect()
                if (rect.top <= 100) {
                    currentSection = section.id
                }
            })

            if (currentSection) {
                setActiveSection(currentSection)
            }
        }

        window.addEventListener("scroll", handleScroll)
        return () => window.removeEventListener("scroll", handleScroll)
    }, [])

    const scrollToTop = () => {
        window.scrollTo({ top: 0, behavior: "smooth" })
    }

    const renderContent = (content: string) => {
        // Split content by code blocks
        const parts = content.split("```")

        return parts.map((part, index) => {
            if (index % 2 === 0) {
                // Regular text content
                return (
                    <div
                        key={index}
                        dangerouslySetInnerHTML={{ __html: part }}
                    />
                )
            } else {
                // Code block
                const [language, ...codeLines] = part.split("\n")
                const code = codeLines.join("\n")

                return (
                    <div
                        key={index}
                        className={styles.codeBlock}
                    >
                        <CodeBlock
                            text={code}
                            language={language}
                            showLineNumbers={true}
                            theme={nord}
                        />
                    </div>
                )
            }
        })
    }

    return (
        <div className="min-h-screen bg-gray-900 text-white">
            <div className={styles.container}>
                {/* Mobile menu button */}
                <button
                    className={styles.mobileMenuButton}
                    onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth={1.5}
                        stroke="currentColor"
                        className={styles.mobileMenuIcon}
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
                        />
                    </svg>
                    {mobileMenuOpen
                        ? "Hide Documentation Menu"
                        : "Show Documentation Menu"}
                </button>

                <div className={styles.flexLayout}>
                    {/* Sidebar Navigation */}
                    <aside
                        className={`${styles.sidebar} ${
                            mobileMenuOpen ? "block" : "hidden lg:block"
                        }`}
                    >
                        <div className={styles.sidebarSticky}>
                            <nav className={styles.navSection}>
                                <h3 className={styles.navTitle}>
                                    Documentation
                                </h3>
                                <ul className={styles.navList}>
                                    {props.sections.map((section) => (
                                        <li key={section.id}>
                                            <a
                                                href={`#${section.id}`}
                                                className={`${styles.navItem} ${
                                                    activeSection === section.id
                                                        ? styles.navItemActive
                                                        : ""
                                                }`}
                                            >
                                                {section.title}
                                            </a>
                                            {section.subsections.length > 0 && (
                                                <ul className="mt-1 ml-4 space-y-1">
                                                    {section.subsections.map(
                                                        (subsection) => (
                                                            <li
                                                                key={
                                                                    subsection.id
                                                                }
                                                            >
                                                                <a
                                                                    href={`#${subsection.id}`}
                                                                    className={`${
                                                                        styles.navItem
                                                                    } ${
                                                                        activeSection ===
                                                                        subsection.id
                                                                            ? styles.navItemActive
                                                                            : ""
                                                                    }`}
                                                                >
                                                                    {
                                                                        subsection.title
                                                                    }
                                                                </a>
                                                            </li>
                                                        ),
                                                    )}
                                                </ul>
                                            )}
                                        </li>
                                    ))}
                                </ul>
                            </nav>
                        </div>
                    </aside>

                    {/* Main Content */}
                    <main className={styles.contentArea}>
                        <header className={styles.docHeader}>
                            <h1 className={styles.docTitle}>{props.title}</h1>
                            <p className={styles.docDescription}>
                                {props.description}
                            </p>
                        </header>

                        {/* Documentation Sections */}
                        <div>
                            {props.sections.map((section) => (
                                <section
                                    key={section.id}
                                    id={section.id}
                                    className={styles.section}
                                >
                                    <header className={styles.sectionHeader}>
                                        <h2 className={styles.sectionTitle}>
                                            {section.title}
                                        </h2>
                                    </header>

                                    <div className={styles.prose}>
                                        {renderContent(section.content)}
                                    </div>

                                    {/* Subsections */}
                                    {section.subsections.map((subsection) => (
                                        <div
                                            key={subsection.id}
                                            id={subsection.id}
                                            className={styles.subsection}
                                        >
                                            <h3
                                                className={
                                                    styles.subsectionTitle
                                                }
                                            >
                                                {subsection.title}
                                            </h3>
                                            <div className={styles.prose}>
                                                {renderContent(
                                                    subsection.content,
                                                )}
                                            </div>
                                        </div>
                                    ))}
                                </section>
                            ))}
                        </div>
                    </main>
                </div>
            </div>

            {/* Back to top button */}
            {showBackToTop && (
                <button
                    onClick={scrollToTop}
                    className={styles.backToTop}
                    aria-label="Back to top"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth={2}
                        stroke="currentColor"
                        className="h-6 w-6"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            d="M4.5 15.75l7.5-7.5 7.5 7.5"
                        />
                    </svg>
                </button>
            )}
        </div>
    )
}
