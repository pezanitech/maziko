export type HomePageProps = {
    title?: string
    description?: string
    features: {
        title: string
        description: string
        icon?: string
    }[]
}

const styles = {
    container: "min-h-screen",
    section: {
        hero: "relative overflow-hidden py-16 text-center",
        features: "py-16",
    },
    background: "bg-background absolute inset-0 -z-10",
    contentContainer: "container mx-auto px-4",
    logoWrapper: "mb-4 flex flex-col items-center justify-center gap-4",
    logo: "bg-primary text-primary-foreground flex h-16 w-16 items-center justify-center rounded-full p-3",
    title: "text-primary text-5xl font-black tracking-tight md:text-7xl",
    description:
        "text-foreground mx-auto max-w-2xl text-lg leading-relaxed font-normal md:text-2xl",
    buttonsWrapper: "mt-10 flex flex-wrap justify-center gap-6",
    primaryButton:
        "bg-primary text-primary-foreground hover:bg-primary/90 rounded-full px-6 py-3 font-medium transition-colors",
    secondaryButton:
        "bg-muted text-muted-foreground hover:bg-muted/90 rounded-full px-6 py-3 font-medium transition-colors",
    featuresHeader: {
        wrapper: "mb-12 text-center",
        title: "mb-4 text-4xl font-extrabold tracking-tight",
        description: "text-muted-foreground mx-auto max-w-2xl",
    },
    featuresGrid: "grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3",
    featureCard: {
        wrapper:
            "border-border bg-card rounded-lg border p-6 transition-shadow hover:shadow-md",
        title: "text-card-foreground mb-2 text-xl font-semibold",
        description: "text-muted-foreground",
        footer: "border-border mt-4 flex items-center justify-between border-t pt-4",
        badge: "bg-muted text-muted-foreground rounded-full px-2 py-1 text-xs",
        link: "text-primary text-sm hover:underline",
    },
}

export const HomePage = (props: HomePageProps) => {
    return (
        <div className={styles.container}>
            {/* Hero Section */}
            <section className={styles.section.hero}>
                <div className={styles.background}></div>

                <div className={styles.contentContainer}>
                    <div className={styles.logoWrapper}>
                        <h1 className={styles.title}>Maziko</h1>
                    </div>
                    <p className={styles.description}>{props.description}</p>

                    <div className={styles.buttonsWrapper}>
                        <a
                            href="https://maziko.pezani.com/docs"
                            className={styles.primaryButton}
                        >
                            Get Started
                        </a>
                        <a
                            href="https://github.com/pezanitech/maziko"
                            className={styles.secondaryButton}
                            target="_blank"
                            rel="noopener noreferrer"
                        >
                            GitHub
                        </a>
                    </div>
                </div>
            </section>

            {/* Features Section */}
            <section className={styles.section.features}>
                <div className={styles.contentContainer}>
                    <div className={styles.featuresHeader.wrapper}>
                        <h2 className={styles.featuresHeader.title}>
                            {props.title || "Features"}
                        </h2>
                        <p className={styles.featuresHeader.description}>
                            Powerful features to build your next project
                        </p>
                    </div>

                    <div className={styles.featuresGrid}>
                        {props.features.map((feature, i) => (
                            <div
                                key={i}
                                className={styles.featureCard.wrapper}
                            >
                                <h3 className={styles.featureCard.title}>
                                    {feature.title}
                                </h3>
                                <p className={styles.featureCard.description}>
                                    {feature.description}
                                </p>
                                <div className={styles.featureCard.footer}>
                                    <span className={styles.featureCard.badge}>
                                        Core
                                    </span>
                                    <a
                                        href="/docs"
                                        className={styles.featureCard.link}
                                    >
                                        Learn more →
                                    </a>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </section>
        </div>
    )
}
