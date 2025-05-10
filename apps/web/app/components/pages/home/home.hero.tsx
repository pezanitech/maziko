import clsx from "clsx"
import { BookOpen, Github } from "lucide-react"

import { Button } from "@/components/ui/button"
import { MazikoLogo } from "@/components/ui/icons"
import { Pattern } from "@/components/ui/pattern"

import { Container } from "@/components/layout/container"

const styles = {
    container: clsx`relative isolate overflow-hidden text-center`,
    gradient: clsx`from-accent/5 via-background to-background absolute inset-0 -z-10 bg-gradient-to-b`,
    header: clsx`mb-4 flex flex-col items-center justify-center gap-4`,
    logo: clsx`animate-float bg-accent/30 text-foreground h-16 w-16 rounded-full p-3 md:h-20 md:w-20`,
    title: clsx`text-accent text-5xl font-black md:text-7xl`,
    headline: clsx`text-muted-foreground mx-auto max-w-3xl text-2xl font-medium md:text-3xl`,
    description: clsx`text-muted-foreground/80 mx-auto mt-4 max-w-2xl text-lg`,
    buttons: clsx`mt-8 flex flex-wrap justify-center gap-6`,
}

type HeroSectionProps = {
    headline: string
    description: string
}

export const HeroSection = (props: HeroSectionProps) => (
    <section className={styles.container}>
        <div className={styles.gradient} />
        <Pattern />
        <Container className="pt-16">
            <div className={styles.header}>
                <MazikoLogo className={styles.logo} />
                <h1 className={styles.title}>Maziko</h1>
            </div>
            <p className={styles.headline}>{props.headline}</p>
            <p className={styles.description}>{props.description}</p>
            <div className={styles.buttons}>
                <Button
                    asChild
                    variant="default"
                    size="lg"
                    className="bg-accent rounded-full md:p-6"
                >
                    <a href="/docs">
                        <BookOpen />
                        Get Started
                    </a>
                </Button>
                <Button
                    asChild
                    variant="outline"
                    size="lg"
                    className="rounded-full md:p-6"
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
