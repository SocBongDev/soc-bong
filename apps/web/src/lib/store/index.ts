import { writable } from 'svelte/store'

export interface NotificationMessage {
	id: string
	title?: string
	description: string
	type: 'error' | 'warning' | 'success' | 'loading'
}

export interface AttendedStatus {
	id?: string
	date?: string
	attendedStatus: 'attended' | 'absented' | 'excused' | 'dayoff' | 'holiday' | 'unknown';
	classId?: string,
	studentId?: string,
}

export const classIdStore = writable<number | null>(1);

export const userRoleStore = writable<string | null>("");

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

export const statusChange = writable<Array<AttendedStatus>>([])


interface ISidebarContext {
	collapseMenu: boolean;
	setCollapseMenu: (v: boolean) => void;
}

export const collapseMenu = writable(false)

export const SidebarContext = writable<ISidebarContext>({
	collapseMenu: false,
	setCollapseMenu: (v) => {
		SidebarContext.update((context) => ({
			...context,
			collapseMenu: v,
		}));
	}
})