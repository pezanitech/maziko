export const brand = {
    name: "Maziko Logo",
    copyright: `© ${new Date().getFullYear()} All rights reserved`,
}

export interface FooterLink {
    name: string
    href: string
    external?: boolean
}

export const links: FooterLink[] = [
    { name: "Privacy Policy", href: "/privacy" },
    { name: "Terms of Service", href: "/terms" },
    { name: "Feedback", href: "/feedback" },
]
