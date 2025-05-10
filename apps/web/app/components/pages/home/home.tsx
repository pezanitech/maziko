import clsx from "clsx"

import { CodeSection } from "./home.code"
import { FeaturesSection } from "./home.features"
import { HeroSection } from "./home.hero"

const styles = {
    root: clsx`min-h-screen`,
}

export type CodeStep = {
    title: string
    description: string
    code: string
    type: "code" | "shell"
    filename?: string
}

export type CodeExample = {
    name: string
    description: string
    code: string
    type: "code" | "shell"
    filename?: string
    steps?: CodeStep[]
}

export type Feature = {
    title: string
    description: string
    icon: string
}

export type HomePageProps = {
    description: string
    features: Feature[]
    featuresTitle: string
    featuresSubtitle?: string
    codeExamples?: CodeExample[]
}

export const HomePage = (props: HomePageProps) => {
    return (
        <div className={styles.root}>
            <HeroSection description={props.description} />
            <CodeSection codeExamples={props.codeExamples} />
            <FeaturesSection
                features={props.features}
                title={props.featuresTitle}
                subtitle={props.featuresSubtitle}
            />
        </div>
    )
}
