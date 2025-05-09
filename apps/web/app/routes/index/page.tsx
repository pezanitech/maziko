// HomePage

import { usePage } from "@inertiajs/react"
import { CodeBlock, nord } from "react-code-blocks"

// Styles object containing all styling
const styles = {
    // Layout styles
    container: "container mx-auto px-4",
    minHeightScreen:
        "min-h-screen bg-gradient-to-b space-y-16 from-gray-900 to-black text-white",

    // Header styles
    header: "pt-32 text-center",
    title: "text-5xl md:text-6xl font-black mb-4",
    headline:
        "text-2xl md:text-3xl font-semibold mb-3 bg-gradient-to-r from-emerald-400 to-emerald-600 bg-clip-text text-transparent",
    description: "text-gray-400 text-lg md:text-xl max-w-3xl mx-auto mb-8",

    // Button styles
    buttonContainer: "flex flex-wrap justify-center gap-4",
    primaryButton:
        "bg-emerald-500 hover:bg-emerald-600 text-black px-8 py-3 rounded-xl font-medium transition-colors",
    secondaryButton:
        "border border-gray-600 hover:border-emerald-500 hover:text-emerald-500 px-8 py-3 rounded-xl font-medium transition-colors",

    // Code example styles
    codeSection: "",
    codeBlock: "bg-gray-800 rounded-xl overflow-hidden max-w-3xl mx-auto",
    codeHeader: "flex items-center gap-2 px-4 py-2",
    codeDot: "w-3 h-3 rounded-full",
    codeRedDot: "w-3 h-3 rounded-full bg-red-500",
    codeYellowDot: "w-3 h-3 rounded-full bg-yellow-500",
    codeGreenDot: "w-3 h-3 rounded-full bg-green-500",
    codeLabel: "text-gray-400 text-sm ml-2",
    codeFont: "font-mono text-sm",

    // Features section styles
    featuresSection: "",
    featuresTitle: "text-3xl font-bold text-center mb-12",
    featuresGrid: "grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8",
    featureCard:
        "bg-gray-950 bg-opacity-50 p-6 rounded-xl hover:bg-emerald-500/10 transition-colors border border-gray-900 hover:border-emerald-500/40",
    featureIcon:
        "w-12 h-12 bg-emerald-600 text-black bg-opacity-20 rounded-2xl flex items-center justify-center mb-4",
    featureTitle: "text-xl font-bold mb-2",
    featureDescription: "text-gray-400",

    // Footer styles
    footer: "border-t border-gray-800 mt-12 py-8 text-center text-gray-500",

    // SVG icon
    iconSvg: "w-6 h-6",
}

export default function Page() {
    const { props } = usePage<{
        headline: string
        description: string
        features: {
            title: string
            description: string
            icon: string
        }[]
        codeExample: string
    }>()

    return (
        <div className={styles.minHeightScreen}>
            {/* Hero Section */}
            <header className={`${styles.container} ${styles.header}`}>
                <h1 className={styles.title}>Maziko</h1>
                <p className={styles.headline}>{props.headline}</p>
                <p className={styles.description}>{props.description}</p>
                <div className={styles.buttonContainer}>
                    <a href="/docs" className={styles.primaryButton}>
                        Get Started
                    </a>
                    <a
                        href="https://github.com/pezanitech/maziko"
                        target="_blank"
                        rel="noopener noreferrer"
                        className={styles.secondaryButton}
                    >
                        GitHub
                    </a>
                </div>
            </header>

            {/* Code Example Section */}
            <section className={`${styles.container} ${styles.codeSection}`}>
                <div className={styles.codeBlock}>
                    <div className={styles.codeHeader}>
                        <div className={styles.codeRedDot}></div>
                        <div className={styles.codeYellowDot}></div>
                        <div className={styles.codeGreenDot}></div>
                        <span className={styles.codeLabel}>Terminal</span>
                    </div>
                    <div className={styles.codeFont}>
                        <CodeBlock
                            text={props.codeExample}
                            language="bash"
                            showLineNumbers={false}
                            theme={nord}
                        />
                    </div>
                </div>
            </section>

            {/* Features Section */}
            <section
                className={`${styles.container} ${styles.featuresSection}`}
            >
                <h2 className={styles.featuresTitle}>Key Features</h2>
                <div className={styles.featuresGrid}>
                    {props.features.map((feature, i) => (
                        <div key={i} className={styles.featureCard}>
                            <div className={styles.featureIcon}>
                                <FeatureIcon icon={feature.icon} />
                            </div>
                            <h3 className={styles.featureTitle}>
                                {feature.title}
                            </h3>
                            <p className={styles.featureDescription}>
                                {feature.description}
                            </p>
                        </div>
                    ))}
                </div>
            </section>

            {/* Footer */}
            <footer className={styles.footer}>
                <div className={styles.container}>
                    <p>© {new Date().getFullYear()} Maziko. MIT License.</p>
                </div>
            </footer>
        </div>
    )
}

// Feature Icon Component
function FeatureIcon({ icon }: { icon: string }) {
    switch (icon) {
        case "rocket":
            return (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className={styles.iconSvg}
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M15.59 14.37a6 6 0 0 1-5.84 7.38v-4.8m5.84-2.58a14.98 14.98 0 0 0 6.16-12.12A14.98 14.98 0 0 0 9.631 8.41m5.96 5.96a14.926 14.926 0 0 1-5.841 2.58m-.119-8.54a6 6 0 0 0-7.381 5.84h4.8m2.581-5.84a14.927 14.927 0 0 0-2.58 5.84m2.699 2.7c-.103.021-.207.041-.311.06a15.09 15.09 0 0 1-2.448-2.448 14.9 14.9 0 0 1 .06-.312m-2.24 2.39a4.493 4.493 0 0 0-1.757 4.306 4.493 4.493 0 0 0 4.306-1.758M16.5 9a1.5 1.5 0 1 1-3 0 1.5 1.5 0 0 1 3 0Z"
                    />
                </svg>
            )
        case "sparkles":
            return (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className={styles.iconSvg}
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M9.813 15.904 9 18.75l-.813-2.846a4.5 4.5 0 0 0-3.09-3.09L2.25 12l2.846-.813a4.5 4.5 0 0 0 3.09-3.09L9 5.25l.813 2.846a4.5 4.5 0 0 0 3.09 3.09L15.75 12l-2.846.813a4.5 4.5 0 0 0-3.09 3.09ZM18.259 8.715 18 9.75l-.259-1.035a3.375 3.375 0 0 0-2.455-2.456L14.25 6l1.036-.259a3.375 3.375 0 0 0 2.455-2.456L18 2.25l.259 1.035a3.375 3.375 0 0 0 2.456 2.456L21.75 6l-1.035.259a3.375 3.375 0 0 0-2.456 2.456Z"
                    />
                </svg>
            )
        case "folder":
            return (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className={styles.iconSvg}
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M2.25 12.75V12A2.25 2.25 0 0 1 4.5 9.75h15A2.25 2.25 0 0 1 21.75 12v.75m-8.69-6.44-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44Z"
                    />
                </svg>
            )
        case "server":
            return (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className={styles.iconSvg}
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M21.75 17.25v-.228a4.5 4.5 0 0 0-.12-1.03l-2.268-9.64a3.375 3.375 0 0 0-3.285-2.602H7.923a3.375 3.375 0 0 0-3.285 2.602l-2.268 9.64a4.5 4.5 0 0 0-.12 1.03v.228m19.5 0a3 3 0 0 1-3 3H5.25a3 3 0 0 1-3-3m19.5 0a3 3 0 0 0-3-3H5.25a3 3 0 0 0-3 3m16.5 0h.008v.008h-.008v-.008Zm-3 0h.008v.008h-.008v-.008Z"
                    />
                </svg>
            )
        case "bolt":
            return (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className={styles.iconSvg}
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="m3.75 13.5 10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75Z"
                    />
                </svg>
            )
        case "wrench":
            return (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className={styles.iconSvg}
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M21.75 6.75a4.5 4.5 0 0 1-4.884 4.484c-1.076-.091-2.264.071-2.95.904l-7.152 8.684a2.548 2.548 0 1 1-3.586-3.586l8.684-7.152c.833-.686.995-1.874.904-2.95a4.5 4.5 0 0 1 6.336-4.486l-3.276 3.276a3.004 3.004 0 0 0 2.25 2.25l3.276-3.276c.256.565.398 1.192.398 1.852Z"
                    />
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M4.867 19.125h.008v.008h-.008v-.008Z"
                    />
                </svg>
            )
        default:
            return (
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className={styles.iconSvg}
                >
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z"
                    />
                </svg>
            )
    }
}
