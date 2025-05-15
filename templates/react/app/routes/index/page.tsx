import { usePage } from "@inertiajs/react"

import { HomePage } from "@/components/pages/home"

const Page = () => {
    const { props } = usePage()

    return <HomePage {...(props as any)} />
}

export default Page
