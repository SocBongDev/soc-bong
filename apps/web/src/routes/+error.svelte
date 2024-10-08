<script lang="ts">
	import { onDestroy, onMount } from 'svelte'
	import { goto } from '$app/navigation'
	import ArrowRightIcon from '~icons/fa-solid/arrow-right'
	import { isAuthenticated } from '$lib/services/auth'
	/**
	 * If user enter any incorrect URL, generic error page will be triggered and will
	 * redirect user to root page
	 */

	let timeoutId: any

	onMount(() => {
		timeoutId = setTimeout(function () {
			if (isAuthenticated()) {
				goto('/admin/registrations')
			} else {
				goto('/')
			}
		}, 30000)
	})

	onDestroy(() => {
		if (timeoutId) {
			clearTimeout(timeoutId)
		}
	})

	function handleGoBack() {
		if (isAuthenticated()) {
			goto('/admin/registrations')
		} else {
			goto('/')
		}
	}
</script>

<svelte:head>
	<title>Error Page</title>
	<meta name="description" content="Error Page" />
</svelte:head>

<section class="flex h-screen flex-col items-center justify-center gap-8">
	<svg xmlns="http://www.w3.org/2000/svg" width="192" height="192" viewBox="0 0 24 24" {...$$props}>
		<path
			fill="none"
			stroke="currentColor"
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width="2"
			d="M9.172 16.172a4 4 0 0 1 5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"
		/>
	</svg>
	<div class="flex w-full flex-col items-center justify-center gap-4">
		<h1 class="w-full text-center text-6xl font-bold">404</h1>
		<p class="text-center">
			Có vẻ như trang này hiện đang không tồn tại hoặc bị lỗi
			<br />
			Xin hãy thông báo đến lập trình viên
		</p>
	</div>

	<button on:click={handleGoBack} class="group btn btn-primary">
		Trở về trang chủ
		<ArrowRightIcon class="transition group-hover:translate-x-2" />
	</button>
</section>
