import clsx from "clsx"
import { ArrowRight } from "lucide-react"

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import {
    Card,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { FeatureIcon } from "@/components/ui/featureIcon"
import { Separator } from "@/components/ui/separator"

import { Container } from "@/components/layout/container"

const styles = {
    titleContainer: clsx`mb-6 flex flex-col-reverse gap-6 lg:flex-row lg:items-center lg:justify-between`,
    titleContent: clsx`flex-1`,
    title: clsx`text-4xl font-extrabold tracking-tight md:text-5xl`,
    subtitle: clsx`text-muted-foreground mt-3 max-w-none text-lg md:text-xl lg:pr-4`,
    separator: clsx`mb-12`,
    grid: clsx`grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3`,
    cardTitle: clsx`flex items-center gap-2 text-xl font-semibold tracking-tight md:text-xl`,
    cardDescription: clsx`text-sm font-light`,
    rightSide: clsx`flex w-full flex-row items-start justify-between gap-3 lg:w-auto lg:flex-col lg:items-end`,
    badge: clsx`p-1 px-2`,
    link: clsx`flex flex-nowrap items-center gap-2 font-medium whitespace-nowrap`,
}

type Feature = {
    title: string
    description: string
    icon: string
}

type FeaturesSectionProps = {
    features: Feature[]
    title: string
    subtitle?: string
}

export const FeaturesSection = (props: FeaturesSectionProps) => (
    <Container className="py-12">
        <div className={styles.titleContainer}>
            <div className={styles.titleContent}>
                <h2 className={styles.title}>{props.title}</h2>
                {props.subtitle && (
                    <p className={styles.subtitle}>{props.subtitle}</p>
                )}
            </div>
            <div className={styles.rightSide}>
                <Badge
                    className={styles.badge}
                    variant="outline"
                >
                    Latest Release
                </Badge>
                <Button
                    asChild
                    variant="accent"
                    size="sm"
                >
                    <a
                        href="/docs"
                        className={styles.link}
                    >
                        View all features{" "}
                        <ArrowRight className="ml-2 h-4 w-4" />
                    </a>
                </Button>
            </div>
        </div>
        <Separator className={styles.separator} />
        <div className={styles.grid}>
            {props.features.map((feature, i) => (
                <Card key={i}>
                    <CardHeader>
                        <CardTitle className={styles.cardTitle}>
                            <FeatureIcon
                                icon={feature.icon}
                                iconSize={20}
                                size="md"
                                variant="default"
                            />
                            {feature.title}
                        </CardTitle>
                        <CardDescription className={styles.cardDescription}>
                            {feature.description}
                        </CardDescription>
                    </CardHeader>
                </Card>
            ))}
        </div>
    </Container>
)
