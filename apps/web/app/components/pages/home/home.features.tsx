import clsx from "clsx"

import {
    Card,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { FeatureIcon } from "@/components/ui/featureIcon"

import { Container } from "@/components/layout/container"

const styles = {
    title: clsx`mb-8 text-center text-3xl font-bold`,
    grid: clsx`grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3`,
    card: clsx`bg-opacity-50 border-border bg-card/50 hover:border-accent/40 hover:bg-accent/10 group transition-colors`,
    iconWrapper: clsx`mb-2`,
}

type Feature = {
    title: string
    description: string
    icon: string
}

type FeaturesSectionProps = {
    features: Feature[]
}

export const FeaturesSection = (props: FeaturesSectionProps) => (
    <Container>
        <h2 className={styles.title}>Key Features</h2>
        <div className={styles.grid}>
            {props.features.map((feature, i) => (
                <Card
                    key={i}
                    className={styles.card}
                >
                    <CardHeader>
                        <div className={styles.iconWrapper}>
                            <FeatureIcon
                                icon={feature.icon}
                                size="lg"
                            />
                        </div>
                        <CardTitle>{feature.title}</CardTitle>
                        <CardDescription className="text-muted-foreground">
                            {feature.description}
                        </CardDescription>
                    </CardHeader>
                </Card>
            ))}
        </div>
    </Container>
)
