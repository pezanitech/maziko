import React from "react"

import clsx from "clsx"
import { HelpCircle, Menu, X } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Link } from "@/components/ui/link"
import { ModeToggle } from "@/components/ui/modeToggle"

import { help } from "./resources/data"

const styles = {
    base: clsx`ml-auto flex`,
    left: clsx`flex items-center gap-x-4`,
    right: clsx`ml-4 md:ml-0`,
    mobileButton: clsx`md:hidden`,
    githubIcon: clsx`h-[1.2rem] w-[1.2rem]`,
    menuIcon: clsx`h-[1.2rem] w-[1.2rem]`,
    notificationBadge: clsx`absolute -top-1 -right-1 flex h-4 w-4 items-center justify-center rounded-full bg-red-500 text-xs text-white`,
}

/* ----TODO: ADD NOTIFICATIONS -----
// Type definitions
type NotificationType = "market" | "news" | "event"
type NotificationIconType = "trend" | "news" | "calendar"

interface Notification {
    id: number
    type: NotificationType
    title: string
    time: string
    icon: NotificationIconType
    read?: boolean
}

// Mock API function to fetch notifications
const fetchNotifications = async (): Promise<Notification[]> => {
    // Simulate API delay
    await new Promise((resolve) => setTimeout(resolve, 500))

    // Return mock notifications data
    return [
        {
            id: 1,
            type: "market",
            title: "AAPL up 3.5% in pre-market trading",
            time: "10 minutes ago",
            icon: "trend",
            read: false,
        },
        {
            id: 2,
            type: "news",
            title: "Fed signals rate cut in next meeting",
            time: "1 hour ago",
            icon: "news",
            read: false,
        },
        {
            id: 3,
            type: "event",
            title: "Tesla earnings call today at 5PM EST",
            time: "2 hours ago",
            icon: "calendar",
            read: false,
        },
    ]
}

const HeaderContentActionsNotifications: React.FC = () => {
    const [notifications, setNotifications] = useState<Notification[]>([])
    const [loading, setLoading] = useState<boolean>(true)

    useEffect(() => {
        const getNotifications = async (): Promise<void> => {
            setLoading(true)
            try {
                const data = await fetchNotifications()
                setNotifications(data)
            } catch (error) {
                console.error("Failed to fetch notifications:", error)
            } finally {
                setLoading(false)
            }
        }

        getNotifications()
    }, [])

    const markAllAsRead = (): void => {
        // In a real app, you would call an API here
        setNotifications([])
    }

    const getIconComponent = (
        iconType: NotificationIconType,
    ): React.ReactNode => {
        switch (iconType) {
            case "trend":
                return (
                    <TrendingUp
                        size={16}
                        className="text-primary"
                    />
                )
            case "news":
                return (
                    <FileText
                        size={16}
                        className="text-primary"
                    />
                )
            case "calendar":
                return (
                    <Calendar
                        size={16}
                        className="text-primary"
                    />
                )
            default:
                return (
                    <Bell
                        size={16}
                        className="text-primary"
                    />
                )
        }
    }

    return (
        <Dialog>
            <DialogTrigger asChild>
                <Button
                    variant="outline"
                    size="icon"
                    className="hover:bg-muted relative"
                >
                    <Bell size={18} />
                    {notifications.length > 0 && (
                        <span className={styles.notificationBadge}>
                            {notifications.length}
                        </span>
                    )}
                </Button>
            </DialogTrigger>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Notifications</DialogTitle>
                    <DialogDescription>
                        Your latest market alerts and news updates
                    </DialogDescription>
                </DialogHeader>
                <div className="mt-4 space-y-4">
                    {loading ? (
                        <p className="text-muted-foreground text-center text-sm">
                            Loading notifications...
                        </p>
                    ) : notifications.length > 0 ? (
                        notifications.map((notification) => (
                            <div
                                key={notification.id}
                                className="flex items-start gap-2 border-b pb-3"
                            >
                                <div className="bg-primary/10 rounded-full p-2">
                                    {getIconComponent(notification.icon)}
                                </div>
                                <div>
                                    <p className="text-sm font-medium">
                                        {notification.title}
                                    </p>
                                    <p className="text-muted-foreground text-xs">
                                        {notification.time}
                                    </p>
                                </div>
                            </div>
                        ))
                    ) : (
                        <p className="text-muted-foreground text-center text-sm">
                            No new notifications
                        </p>
                    )}
                </div>
                {notifications.length > 0 && (
                    <div className="mt-2 flex justify-end">
                        <Button
                            variant="outline"
                            size="sm"
                            onClick={markAllAsRead}
                        >
                            Mark all as read
                        </Button>
                    </div>
                )}
            </DialogContent>
        </Dialog>
    )
}----*/

interface HeaderContentActionsProps {
    isOpen: boolean
    onOpenChange: (open: boolean) => void
}

export const HeaderContentActions: React.FC<HeaderContentActionsProps> = ({
    isOpen,
    onOpenChange,
}: HeaderContentActionsProps) => (
    <div className={styles.base}>
        <div className={styles.left}>
            {/* <HeaderContentActionsNotifications /> */}
            <Button
                variant="outline"
                size="icon"
                asChild
            >
                <Link href={help.href}>
                    <HelpCircle className={styles.githubIcon} />
                    <span className="sr-only">{help.ariaLabel}</span>
                </Link>
            </Button>
            <ModeToggle />
        </div>
        <div className={styles.right}>
            <Button
                variant="outline"
                size="icon"
                className={styles.mobileButton}
                onClick={() => onOpenChange(!isOpen)}
            >
                {isOpen ? (
                    <X className={styles.menuIcon} />
                ) : (
                    <Menu className={styles.menuIcon} />
                )}
            </Button>
        </div>
    </div>
)
