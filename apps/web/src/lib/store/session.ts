import { browser } from '$app/environment'
import { writable } from 'svelte/store'

const defaultValue = null

const initialAccessToken = browser
	? window.localStorage.getItem('access_token') ?? defaultValue
	: defaultValue
const initalExpiresAt = browser
	? window.localStorage.getItem('expiresAt') ?? defaultValue
	: defaultValue
const initalRoles = browser ? window.localStorage.getItem('role') ?? defaultValue : defaultValue
const initalUserId = browser ? window.localStorage.getItem('user_id') ?? defaultValue : defaultValue

export const accessToken = writable<string | null>(initialAccessToken)
export const expiresAt = writable<string | null>(initalExpiresAt)
export const role = writable<string | null>(initalRoles)
export const userId = writable<string | null>(initalUserId)

accessToken.subscribe((value) => {
	if (browser) {
		if (value === null) {
			localStorage.removeItem('access_token')
		} else {
			localStorage.setItem('access_token', value)
		}
	}
})

expiresAt.subscribe((value) => {
	if (browser) {
		if (value === null) {
			localStorage.removeItem('expiresAt')
		} else {
			localStorage.setItem('expiresAt', value)
		}
	}
})

role.subscribe((value) => {
	if (browser) {
		if (value === null) {
			localStorage.removeItem('role')
		} else {
			localStorage.setItem('role', value)
		}
	}
})

userId.subscribe((value) => {
	if (browser) {
		if (value === null) {
			localStorage.removeItem('user_id')
		} else {
			localStorage.setItem('user_id', value)
		}
	}
})
