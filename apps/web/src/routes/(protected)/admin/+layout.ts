import { getSession, isAuthenticated } from '$lib/services/auth'
import { redirect } from '@sveltejs/kit'

export async function load({ params, parent, url, data }) {
	if (url.pathname.startsWith('/')) {
		const session = getSession()
		if (session) {
			if (isAuthenticated()) {
				return {}
			}

			throw redirect(303, '/')
		}

		throw redirect(303, '/')
	}
}
