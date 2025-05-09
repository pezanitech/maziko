import clsx from "clsx"

import { HeaderContentActions } from "./header.content.actions"
import { HeaderContentLogo } from "./header.content.logo"
import { HeaderContentNav } from "./header.content.nav"

const styles = {
    wrapper: clsx`flex h-16`,

    mainWrapper: clsx`flex flex-1 items-center gap-6`,
}

interface HeaderContentProps {
    isOpen: boolean
    onOpenChange: (open: boolean) => void
}

export const HeaderContent = ({ isOpen, onOpenChange }: HeaderContentProps) => (
    <div className={styles.wrapper}>
        <div className={styles.mainWrapper}>
            <HeaderContentLogo />
            <HeaderContentNav />
            <HeaderContentActions
                isOpen={isOpen}
                onOpenChange={onOpenChange}
            />
        </div>
    </div>
)
