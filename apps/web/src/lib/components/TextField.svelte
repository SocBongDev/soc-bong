<script lang="ts">
	import Error from './FloatingLabel/Error.svelte'

	export let name: string
	export let required = false
	export let value: string | undefined = undefined
	export let error: string[] | null = null
	export let disabled: boolean = false
</script>

<div class="form-control w-full flex-col-reverse overflow-hidden rounded">
	<input
		{name}
		type="text"
		class="peer w-full bg-neutral-100 px-4 pb-2.5 pt-1 font-normal text-black focus:bg-neutral-200 focus:outline-none disabled:select-none disabled:hover:cursor-not-allowed"
		bind:value
		{disabled}
	/>
	<label
		class="label bg-neutral-100 px-4 py-1 text-neutral-500 peer-focus:bg-neutral-200 peer-focus:text-black"
		for="text"
	>
		<span
			class="label-text font-semibold text-inherit {required
				? "-ml-px -mt-0.5 after:text-xs after:text-red-500 after:content-['_*']"
				: ''}"
		>
			{name}
		</span>
	</label>
	{#if error !== null}
		{#each error as e}
			<Error>
				{e}
			</Error>
		{/each}
	{/if}
</div>
