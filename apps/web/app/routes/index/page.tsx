// HomePage
import { usePage } from "@inertiajs/react"
import {
    Folder,
    HelpCircle,
    Rocket,
    Server,
    Sparkles,
    Wrench,
    Zap,
} from "lucide-react"
import { CodeBlock, nord } from "react-code-blocks"

import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Container } from "@/components/ui/container"

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
        <div className="min-h-screen space-y-16 bg-gradient-to-b from-gray-900 to-black text-white">
            {/* Hero Section */}
            <Container className="pt-32 text-center">
                <h1 className="text-accent mb-4 text-5xl font-black md:text-6xl">
                    Maziko
                </h1>
                <p className="mb-3 bg-gradient-to-r from-emerald-400 to-emerald-600 bg-clip-text text-2xl font-semibold text-transparent md:text-3xl">
                    {props.headline}
                </p>
                <p className="mx-auto mb-8 max-w-3xl text-lg text-gray-400 md:text-xl">
                    {props.description}
                </p>
                <div className="flex flex-wrap justify-center gap-4">
                    <Button
                        asChild
                        variant="default"
                        className="bg-emerald-500 px-8 py-3 text-black hover:bg-emerald-600"
                    >
                        <a href="/docs">Get Started</a>
                    </Button>
                    <Button
                        asChild
                        variant="outline"
                        className="hover:border-emerald-500 hover:text-emerald-500"
                    >
                        <a
                            href="https://github.com/pezanitech/maziko"
                            target="_blank"
                            rel="noopener noreferrer"
                        >
                            GitHub
                        </a>
                    </Button>
                </div>
            </Container>

            {/* Code Example Section */}
            <Container>
                <Card className="mx-auto max-w-3xl border-gray-700 bg-gray-800">
                    <div className="flex items-center gap-2 px-4 py-2">
                        <div className="h-3 w-3 rounded-full bg-red-500"></div>
                        <div className="h-3 w-3 rounded-full bg-yellow-500"></div>
                        <div className="h-3 w-3 rounded-full bg-green-500"></div>
                        <span className="ml-2 text-sm text-gray-400">
                            Terminal
                        </span>
                    </div>
                    <CardContent className="p-0">
                        <CodeBlock
                            text={props.codeExample}
                            language="bash"
                            showLineNumbers={false}
                            theme={nord}
                        />
                    </CardContent>
                </Card>
            </Container>

            {/* Features Section */}
            <Container>
                <h2 className="mb-12 text-center text-3xl font-bold">
                    Key Features
                </h2>
                <div className="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
                    {props.features.map((feature, i) => (
                        <Card
                            key={i}
                            className="bg-opacity-50 border-gray-900 bg-gray-950 transition-colors hover:border-emerald-500/40 hover:bg-emerald-500/10"
                        >
                            <CardHeader>
                                <div className="bg-opacity-20 mb-4 flex h-12 w-12 items-center justify-center rounded-2xl bg-emerald-600 text-black">
                                    <FeatureIcon icon={feature.icon} />
                                </div>
                                <CardTitle>{feature.title}</CardTitle>
                                <CardDescription className="text-gray-400">
                                    {feature.description}
                                </CardDescription>
                            </CardHeader>
                        </Card>
                    ))}
                </div>
            </Container>

            {/* Footer */}
            <footer className="mt-12 border-t border-gray-800 py-8 text-center text-gray-500">
                <Container>
                    <p>© {new Date().getFullYear()} Maziko. MIT License.</p>
                </Container>
            </footer>
        </div>
    )
}

// Feature Icon Component
function FeatureIcon({ icon }: { icon: string }) {
    switch (icon) {
        case "rocket":
            return <Rocket className="h-6 w-6" />
        case "sparkles":
            return <Sparkles className="h-6 w-6" />
        case "folder":
            return <Folder className="h-6 w-6" />
        case "server":
            return <Server className="h-6 w-6" />
        case "bolt":
            return <Zap className="h-6 w-6" />
        case "wrench":
            return <Wrench className="h-6 w-6" />
        default:
            return <HelpCircle className="h-6 w-6" />
    }
}
