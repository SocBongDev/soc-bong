<script>
	import { onMount } from 'svelte'
	import { page } from '$app/stores'
	import { goto } from '$app/navigation'
	import { createSession, getSession, isAuthenticated } from '$lib/services/auth/index'
	import { redirect } from '@sveltejs/kit'
	import { PUBLIC_AUTH0_CALLBACK_URL } from '$env/static/public'

	onMount(() => {
		;(async () => {
			if ($page.url.searchParams.get('code') && $page.url.searchParams.get('state')) {
				const code = $page.url.searchParams.get('code')
				const session = getSession()
				if (session && !code) {
					if (isAuthenticated()) {
						return goto('/admin/registrations')
					}
					throw redirect(303, '/')
				}

				const options = {
					grantType: 'authorization_code',
					code: code,
					redirectUri: `${PUBLIC_AUTH0_CALLBACK_URL}`
				}

				return $page.data.authenticationClient.oauthToken(options, (err, response) => {
					if (err) {
						return goto('/')
					}
					if (response && response.accessToken && response.idToken) {
						return $page.data.webAuthClient.validateToken(
							response.idToken,
							null,
							(err, payload) => {
								if (err) {
									return goto('/')
								}
								createSession(response)
								if (isAuthenticated()) {
									return goto('/admin/registrations')
								}
								return goto('/')
							}
						)
					}
					return goto('/')
				})
			} else {
				if (isAuthenticated()) {
					return goto('/admin/registrations')
				}
				return goto('/')
			}
		})()
	})
</script>

<div class="callback">
	<div class="loading-ellipsis">Loading...</div>
</div>

<style>
	.callback {
		top: 45%;
		position: relative;
		margin-left: auto;
		margin-right: auto;
		width: 400px;
		height: 100vh;
	}
</style>
