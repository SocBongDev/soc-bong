<script lang="ts">
	import { onDestroy, onMount } from 'svelte'
	import { fade } from 'svelte/transition'

	export let autoClose: (() => NodeJS.Timeout) | undefined = undefined
	let timeout: NodeJS.Timeout | undefined

	function handleClearTimeout() {
		clearTimeout(timeout)
	}

	function mount() {
		timeout = autoClose?.()
	}

	onMount(mount)
	onDestroy(handleClearTimeout)
</script>

<label
	class="alert alert-error"
	on:mouseenter={handleClearTimeout}
	on:mouseleave={autoClose}
	transition:fade
>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		class="h-6 w-6 shrink-0 stroke-current"
		fill="none"
		viewBox="0 0 24 24"
		><path
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width="2"
			d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
		/></svg
	>
	<span>
		<slot />
	</span>
	<div>
		<button class="btn btn-ghost btn-sm active:translate-y-1" on:click>Close</button>
	</div>
</label>
