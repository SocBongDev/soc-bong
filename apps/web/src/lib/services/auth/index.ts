import { browser } from '$app/environment'
import { get } from 'svelte/store'
import { accessToken, expiresAt, role, userId } from '../../store/session'
export function getSession() {
	if (!browser) return null

	const session = {
		accessToken: get(accessToken),
		expiresAt: get(expiresAt)
	}

	if (!session.accessToken || !session.expiresAt) {
		return null
	}

	return session
}

export function createSession(s: any) {
	const expiresAtValue = JSON.stringify(s.expiresIn * 1000 + new Date().getTime())
	accessToken.set(s.accessToken)
	expiresAt.set(expiresAtValue)
}

export function isAuthenticated() {
	const session = getSession()
	if (!session) {
		return false
	}
	// Check whether the current time is past the access token's expiry time
	const expiresAtTime = JSON.parse(session.expiresAt || '')
	return new Date().getTime() < expiresAtTime
}

export function clearSession() {
	accessToken.set(null)
	expiresAt.set(null)
	role.set(null)
	userId.set(null)

	if (browser) {
		localStorage.removeItem('access_token')
		localStorage.removeItem('expiresAt')
		localStorage.removeItem('role')
		localStorage.removeItem('user_id')
	}
}
