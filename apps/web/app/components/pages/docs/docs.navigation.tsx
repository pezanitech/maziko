import clsx from "clsx"
import { ChevronRight } from "lucide-react"
import { Button } from "@/components/ui/button"
import { ScrollArea } from "@/components/ui/scroll-area"
import { type Section } from "./docs"

const styles = {
    nav: clsx`w-full lg:w-64 sticky top-4 h-fit prose-headings:mt-0`,
    mobileHidden: clsx`hidden lg:block`,
    title: clsx`font-semibold text-lg leading-7`,
    list: clsx`mt-4 space-y-2`,
    subsectionList: clsx`pl-4 mt-1 space-y-1`,
    link: clsx`w-full justify-start text-left text-sm h-8 px-2`,
    linkActive: clsx`bg-accent/10 text-accent font-medium hover:bg-accent/20`,
    chevron: clsx`mr-2 h-3 w-3 transition-transform group-data-[state=open]:rotate-90`,
}

type DocsNavigationProps = {
    sections: Section[]
    activeSection: string
    mobileMenuOpen: boolean
}

export const DocsNavigation = ({ sections, activeSection, mobileMenuOpen }: DocsNavigationProps) => (
    <nav className={clsx(styles.nav, !mobileMenuOpen && styles.mobileHidden)}>
        <ScrollArea className="h-[calc(100vh-8rem)]">
            <h2 className={styles.title}>Documentation</h2>
            <div className={styles.list}>
                {sections.map((section) => (
                    <div key={section.id}>
                        <Button
                            variant="ghost"
                            size="sm"
                            className={clsx(
                                styles.link,
                                activeSection === section.id && styles.linkActive
                            )}
                            asChild
                        >
                            <a href={`#${section.id}`}>
                                <ChevronRight className={styles.chevron} />
                                {section.title}
                            </a>
                        </Button>
                        
                        {section.subsections.length > 0 && (
                            <div className={styles.subsectionList}>
                                {section.subsections.map((subsection) => (
                                    <Button
                                        key={subsection.id}
                                        variant="ghost"
                                        size="sm"
                                        className={clsx(
                                            styles.link,
                                            activeSection === subsection.id && styles.linkActive
                                        )}
                                        asChild
                                    >
                                        <a href={`#${subsection.id}`}>
                                            {subsection.title}
                                        </a>
                                    </Button>
                                ))}
                            </div>
                        )}
                    </div>
                ))}
            </div>
        </ScrollArea>
    </nav>
)