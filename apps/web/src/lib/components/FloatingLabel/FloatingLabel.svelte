<script lang="ts">
	import type { SvelteHTMLElements } from 'svelte/elements'
	import Error from './Error.svelte'

	export let placeholder: string | undefined = undefined
	export let name: string | undefined = undefined
	export let id: string | undefined = undefined
	export let error: string[] | null = null
	export let type: SvelteHTMLElements['input']['type'] = undefined
</script>

<div class="floatingContainer">
	<input
		{id}
		{name}
		{type}
		on:change
		class="peer block w-full appearance-none border-0 border-b-2 border-green-300 bg-transparent px-0 py-2.5 text-sm text-gray-900 focus:border-green-600 focus:outline-none focus:ring-0 {error
			? 'border-red-600 focus:border-red-600'
			: ''}"
		placeholder=" "
	/>
	<label
		for="floating_standard"
		class="absolute top-3 -z-10 origin-[0] -translate-y-6 scale-75 transform text-sm text-gray-500 duration-300 peer-placeholder-shown:translate-y-0 peer-placeholder-shown:scale-100 peer-focus:left-0 peer-focus:-translate-y-6 peer-focus:scale-75 peer-focus:text-green-600 {error
			? 'text-red-600'
			: ''}">{placeholder}</label
	>
	{#if error !== null}
		{#each error as e}
			<Error>
				{e}
			</Error>
		{/each}
	{/if}
</div>

<style class="postcss">
	.floatingContainer {
		position: relative;
		z-index: 0;
	}
</style>
