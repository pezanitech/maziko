import { useEffect, useState } from "react"
import { Menu } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet"

import { DocsContent } from "./docs.content"
import { DocsHeader } from "./docs.header"
import { DocsNavigation } from "./docs.navigation"
import { BackToTop } from "./docs.backtotop"

export type Subsection = {
    title: string
    id: string
    content: string
}

export type Section = {
    title: string
    id: string
    content: string
    subsections: Subsection[]
}

export type DocsPageProps = {
    title: string
    description: string
    sections: Section[]
}

export function DocsPage(props: DocsPageProps) {
    const [activeSection, setActiveSection] = useState("")
    const [showBackToTop, setShowBackToTop] = useState(false)

    useEffect(() => {
        const handleScroll = () => {
            setShowBackToTop(window.scrollY > 500)

            const sectionElements = document.querySelectorAll('[id^="section-"]')
            let currentSection = ""

            sectionElements.forEach((section) => {
                const rect = section.getBoundingClientRect()
                if (rect.top <= 100) {
                    currentSection = section.id
                }
            })

            if (currentSection) {
                setActiveSection(currentSection)
            }
        }

        window.addEventListener("scroll", handleScroll)
        return () => window.removeEventListener("scroll", handleScroll)
    }, [])

    return (
        <div className="min-h-screen bg-background">
            <div className="container relative mx-auto px-4 py-8">
                <div className="lg:hidden mb-4">
                    <Sheet>
                        <SheetTrigger asChild>
                            <Button variant="outline" size="icon">
                                <Menu className="h-6 w-6" />
                            </Button>
                        </SheetTrigger>
                        <SheetContent side="left" className="w-80">
                            <DocsNavigation 
                                sections={props.sections}
                                activeSection={activeSection}
                                mobileMenuOpen={true}
                            />
                        </SheetContent>
                    </Sheet>
                </div>

                <div className="flex flex-col lg:flex-row gap-8">
                    <div className="hidden lg:block">
                        <DocsNavigation 
                            sections={props.sections}
                            activeSection={activeSection}
                            mobileMenuOpen={true}
                        />
                    </div>

                    <main className="flex-1 min-w-0">
                        <DocsHeader 
                            title={props.title}
                            description={props.description}
                        />
                        <DocsContent sections={props.sections} />
                    </main>
                </div>
            </div>

            <BackToTop show={showBackToTop} />
        </div>
    )
}
