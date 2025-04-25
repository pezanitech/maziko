import { usePage } from "@inertiajs/react"

export default function Home() {
    const { text } = usePage().props

    return <div className="bg-red-500">Home: {text}</div>
}
