import auth0 from 'auth0-js'
import {
	PUBLIC_AUTH0_CLIENT_ID,
	PUBLIC_AUTH0_DOMAIN,
	PUBLIC_AUTH0_AUDIENCE,
	PUBLIC_AUTH0_CALLBACK_URL
} from '$env/static/public'
export const ssr = false
export const prerender = true

export async function load() {
	const options = {
		clientID: PUBLIC_AUTH0_CLIENT_ID,
		domain: PUBLIC_AUTH0_DOMAIN,
		responseType: 'token',
		audience: PUBLIC_AUTH0_AUDIENCE,
		redirectUri: PUBLIC_AUTH0_CALLBACK_URL,
		scope: 'openid profile email'
	}
	const webAuthClient = new auth0.WebAuth(options)
	const authenticationClient = new auth0.Authentication(options)

	return { webAuthClient, authenticationClient }
}
