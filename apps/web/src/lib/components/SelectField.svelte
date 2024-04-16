<script lang="ts">
	import Error from './FloatingLabel/Error.svelte'

	export let name: string
	export let options: { label: string; value: string }[] = []
	export let required = false
	export let error: string[] | null = null
</script>

<div class="form-control w-full flex-col-reverse overflow-hidden rounded">
	<select class="peer select w-full bg-neutral-100 focus:bg-neutral-200" {name}>
		{#each options as { label, value }, i (label)}
			<option selected={i === 0} {value}>{label}</option>
		{/each}
	</select>
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
