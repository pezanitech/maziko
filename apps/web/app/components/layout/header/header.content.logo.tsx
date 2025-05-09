import clsx from "clsx"
import Image from "next/image"

import { Link } from "@/components/ui/link"

import { brand } from "./resources/data"

const styles = {
    logo: clsx`relative h-10 w-10 overflow-hidden rounded-md`,
}

export const HeaderContentLogo = () => (
    <Link
        href={brand.href}
        className={styles.logo}
    >
        <Image
            fill
            alt="MSE Today"
            src="/images/mse-today-logo.png"
        />
    </Link>
)
