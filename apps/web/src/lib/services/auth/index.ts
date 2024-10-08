export function getSession() {
	const session = {
		accessToken: localStorage.getItem('access_token'),
		idToken: localStorage.getItem('id_token'),
		expiresAt: localStorage.getItem('expires_at')
	}

	if (!session.accessToken || !session.idToken || !session.expiresAt) {
		return null
	}
	return session
}

export function createSession(s: any) {
	const expiresAt = JSON.stringify(s.expiresIn * 1000 + new Date().getTime())
	localStorage.setItem('access_token', s.accessToken)
	localStorage.setItem('id_token', s.idToken)
	localStorage.setItem('expires_at', expiresAt)
}

export function isAuthenticated() {
	const session = getSession()
	if (!session) {
		return false
	}
	// Check whether the current time is past the access token's expiry time
	const expiresAt = JSON.parse(session.expiresAt || '')
	return new Date().getTime() < expiresAt
}
