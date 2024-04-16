import { writable } from 'svelte/store'

export interface NotificationMessage {
	id: string
	title?: string
	description: string
	type: 'error' | 'warning' | 'success' | 'loading'
}

export const notifications = writable<Array<NotificationMessage>>([])

export function Notify(notification: NotificationMessage) {
	notifications.update((n) => [...n, notification])
}

export function CloseNotification(id: string) {
	notifications.update((n) => n.filter((item) => item.id !== id))
}

export type CustomMouseEventHandler = (e?: MouseEvent) => void

export interface DialogProps {
	onContinue?: CustomMouseEventHandler
	onClose?: CustomMouseEventHandler
	open?: boolean
	cancelLable?: string
	okLabel?: string
	title?: string
	withCloseButton?: boolean
	description?: string
}

export const dialogProps = writable<DialogProps | undefined>(undefined)

export const openDialog = writable(false)
