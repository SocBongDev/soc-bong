<script lang="ts">
	import { onMount } from 'svelte'
	import { page } from '$app/stores'
	import { goto } from '$app/navigation'
	import { createSession, getSession, isAuthenticated } from '$lib/services/auth/index'
	import { redirect } from '@sveltejs/kit'
	import { PUBLIC_AUTH0_CALLBACK_URL } from '$env/static/public'
	import { Notify, userRoleStore } from '$lib/store'

	function hasRequiredPermission(access_token: string): boolean {
		try {
			const decodedToken = JSON.parse(atob(access_token.split('.')[1]))
			const permissions = decodedToken.permissions || []
			const roles = decodedToken['user/roles']
			localStorage.setItem("role", roles[0])
			localStorage.setItem("user_id", decodedToken.sub)
			userRoleStore.set(roles[0])
			//define the minimum required permissions here:
			const requiredPermission = ['read:registrations']
			return requiredPermission.every((permission) => permissions.includes(permission))
		} catch (error) {
			console.error('check hasRequiredPer error: ', error)
			return false
		}
	}

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

				return $page.data.authenticationClient.oauthToken(options, (err: any, response: any) => {
					if (err) {
						Notify({
							type: 'error',
							id: crypto.randomUUID(),
							description: 'Chương trình xảy ra lỗi ngoài mong muốn hãy thử lại!'
						})
						return goto('/')
					}
					if (response && response.accessToken && response.idToken) {
						return $page.data.webAuthClient.validateToken(
							response.idToken,
							null,
							(err: any, payload: any) => {
								if (err) {
									return goto('/')
								}
								createSession(response)
								if (isAuthenticated()) {
									if (hasRequiredPermission(response.accessToken)) {
										return goto('/admin/registrations')
									} else {
										Notify({
											type: 'warning',
											id: crypto.randomUUID(),
											description: 'Tài khoản của bạn đang chờ được phê duyệt.'
										})
										return goto('/')
									}
								}
								Notify({
									type: 'error',
									id: crypto.randomUUID(),
									description: 'Bạn cần đăng nhập để thực hiện các hoạt động kế tiếp!'
								})
								return goto('/')
							}
						)
					}
					Notify({
						type: 'error',
						id: crypto.randomUUID(),
						description: 'Bạn cần đăng nhập để thực hiện các hoạt động kế tiếp!'
					})
					return goto('/')
				})
			} else {
				if (isAuthenticated()) {
					const session = getSession()
					if (session && hasRequiredPermission(session.accessToken || '')) {
						return goto('/admin/registrations')
					}
				} else {
					Notify({
						type: 'warning',
						id: crypto.randomUUID(),
						description: 'Tài khoản của bạn đang chờ được phê duyệt.'
					})
					return goto('/')
				}
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'Bạn cần đăng nhập để thực hiện các hoạt động kế tiếp!'
				})
				return goto('/')
			}
		})()
	})
</script>

<div class="callback">
	<div class="loading-container">
		<!-- <div class="loading-text">Loading</div> -->
		<div class="loading-animation">
			<div class="dot" />
			<div class="dot" />
			<div class="dot" />
			<div class="dot" />
		</div>
	</div>
</div>

<style>
	.callback {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
		background-color: #f0f0f0;
	}

	.loading-container {
		text-align: center;
	}

	.loading-animation {
		display: flex;
		justify-content: center;
	}

	.dot {
		width: 20px;
		height: 20px;
		background-color: #121212;
		border-radius: 50%;
		margin: 0 5px;
		animation: bounce 1.4s infinite ease-in-out both;
	}

	.dot:nth-child(1) {
		animation-delay: -0.32s;
	}

	.dot:nth-child(2) {
		animation-delay: -0.16s;
	}

	@keyframes bounce {
		0%,
		80%,
		100% {
			transform: scale(0);
		}
		40% {
			transform: scale(1);
		}
	}
</style>
