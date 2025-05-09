import clsx from "clsx"
import { BookOpen, Github } from "lucide-react"

import { Button } from "@/components/ui/button"
import { CircuitPattern } from "@/components/ui/circuitPattern"
import { MazikoLogo } from "@/components/ui/icons"

import { Container } from "@/components/layout/container"

const styles = {
    container: clsx`relative isolate overflow-hidden text-center`,
    gradient: clsx`from-accent/5 via-background to-background absolute inset-0 -z-10 bg-gradient-to-b`,
    header: clsx`mb-4 flex flex-col items-center justify-center gap-4`,
    logo: clsx`animate-float bg-accent/50 text-foreground h-16 w-16 rounded-full p-3 md:h-20 md:w-20`,
    title: clsx`text-accent text-6xl font-black tracking-tight md:text-8xl`,
    description: clsx`text-muted-foreground mx-auto max-w-2xl text-lg leading-relaxed font-medium md:text-2xl`,
    buttons: clsx`mt-10 flex flex-wrap justify-center gap-6`,
    button: clsx`w-32 justify-center md:w-40`,
}

type HeroSectionProps = {
    description: string
}

export const HeroSection = (props: HeroSectionProps) => (
    <section className={styles.container}>
        <div className={styles.gradient} />
        <CircuitPattern
            glowEffect
            opacity={0.3}
        />
        <Container className="pt-16 pb-8">
            <div className={styles.header}>
                <MazikoLogo className={styles.logo} />
                <h1 className={styles.title}>Maziko</h1>
            </div>
            <p className={styles.description}>{props.description}</p>
            <div className={styles.buttons}>
                <Button
                    asChild
                    variant="accent"
                    size="lg"
                    className={`rounded-full md:p-6 ${styles.button}`}
                >
                    <a href="/docs">
                        <BookOpen />
                        Get Started
                    </a>
                </Button>
                <Button
                    asChild
                    variant="default"
                    size="lg"
                    className={`rounded-full md:p-6 ${styles.button}`}
                >
                    <a
                        href="https://github.com/pezanitech/maziko"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        <Github />
                        GitHub
                    </a>
                </Button>
            </div>
        </Container>
    </section>
)
