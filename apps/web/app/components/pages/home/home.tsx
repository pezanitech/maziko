import clsx from "clsx"

import { CodeSection } from "./home.code"
import { FeaturesSection } from "./home.features"
import { HeroSection } from "./home.hero"

const styles = {
    root: clsx`min-h-screen`,
}

export type HomePageProps = {
    description: string
    features: {
        title: string
        description: string
        icon: string
    }[]
    codeExample: string
}

export const HomePage = (props: HomePageProps) => {
    return (
        <div className={styles.root}>
            <HeroSection description={props.description} />
            <CodeSection codeExample={props.codeExample} />
            <FeaturesSection features={props.features} />
        </div>
    )
}
