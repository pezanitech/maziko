import { usePage } from "@inertiajs/react"

import { HomePage, type HomePageProps } from "@/components/pages/home/home"

const Page = () => {
    const { props } = usePage<HomePageProps>()

    return <HomePage {...props} />
}

export default Page
