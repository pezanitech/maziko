import React from "react"

import clsx from "clsx"

const styles = {
    viewport: clsx`container mx-auto py-16`,
}

export const Viewport = (props: { children: React.ReactNode }) => (
    <div className={styles.viewport}>{props.children}</div>
)
