import { usePage } from "@inertiajs/react"

export default function Page() {
    const { props } = usePage()

    const styles = {
        main: "h-[100dvh] bg-amber-500 text-white flex items-center justify-center px-4",
        container:
            "max-w-sm bg-[#080808] border border-[#121212] px-4 py-16 rounded-3xl w-full text-center space-y-6",
        heading: "text-4xl font-extrabold tracking-tight sm:text-5xl",
        paragraph: "text-gray-400 text-lg",
        buttonContainer: "flex justify-center gap-4",
        buttonPrimary:
            "bg-white text-black px-5 py-2 rounded-xl font-medium hover:bg-gray-200 transition",
        buttonSecondary:
            "border border-white px-5 py-2 rounded-xl font-medium hover:bg-white hover:text-black transition",
    }

    return (
        <main className={styles.main}>
            <div className={styles.container}>
                <h1 className={styles.heading}>Maziko</h1>
                <p className={styles.paragraph}>
                    <span className="block">{props.line1}</span>
                    <span className="block">{props.line2}</span>
                </p>

                <div className={styles.buttonContainer}>
                    <a href="/docs" className={styles.buttonPrimary}>
                        Docs
                    </a>
                    <a
                        href="https://github.com/pezanitech/maziko"
                        target="_blank"
                        rel="noopener noreferrer"
                        className={styles.buttonSecondary}
                    >
                        GitHub
                    </a>
                </div>
            </div>
        </main>
    )
}
