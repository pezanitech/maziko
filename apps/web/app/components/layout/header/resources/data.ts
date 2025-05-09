export interface NavigationItem {
    name: string
    href: string
}

export const navigation: NavigationItem[] = [{ name: "News", href: "/news" }]

export const help = {
    href: "/help",
    ariaLabel: "Help",
}

export const brand = {
    name: "MSE Today",
    href: "/",
}
