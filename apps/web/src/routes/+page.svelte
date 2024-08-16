<script lang="ts">
	import type { PageData } from './$types'
	import ArrowRightIcon from '~icons/fa-solid/arrow-right'
	import { PUBLIC_AUTH0_CLIENT_ID, PUBLIC_AUTH0_CALLBACK_URL } from '$env/static/public'
	import { goto } from '$app/navigation'

	export let data: PageData

	function login() {
		const options = {
			clientID: PUBLIC_AUTH0_CLIENT_ID,
			redirectUri: PUBLIC_AUTH0_CALLBACK_URL,
			responseType: 'code'
		}
		data.webAuthClient.authorize(options)
	}

	function signup() {
		goto('/signup')
	}
</script>

<div class="flex h-screen w-full flex-col items-center justify-center gap-4">
	<button class="btn btn-primary btn-block group max-w-xs" on:click={() => login()}>
		<span class="capitalize">Đăng Nhập</span>
		<ArrowRightIcon class="transition group-hover:translate-x-2" /></button
	>

	<button class="btn btn-primary btn-block group max-w-xs" on:click={() => signup()}>
		<span class="capitalize">Đăng Kí Tài Khoản</span>
		<ArrowRightIcon class="transition group-hover:translate-x-2" /></button
	>
</div>
