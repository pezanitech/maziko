import { usePage } from "@inertiajs/react"

import { DocsPage, type DocsPageProps } from "@/components/pages/docs"

export default function Page() {
    const { props } = usePage<DocsPageProps>()

    return <DocsPage {...props} />
}
