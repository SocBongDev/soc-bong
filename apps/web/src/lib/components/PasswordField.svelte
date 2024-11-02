<script lang="ts">
	import Error from './FloatingLabel/Error.svelte'

	export let name: string
	export let viName: string
	export let required = false
	export let value: string | undefined = undefined
	export let error: string[] | null = null
	export let disabled = false
	let isShow = false

	function handleInput(e: Event) {
		const target = e.target as HTMLInputElement
		value = target.value
	}
	function toggleVisibility() {
		isShow = !isShow
	}
</script>

<div class="form-control w-full overflow-hidden rounded">
	<div class="relative flex flex-col-reverse">
		<input
			{name}
			type={isShow ? 'text' : 'password'}
			class="peer w-full transform bg-neutral-100 px-4 pb-2.5 pt-1 font-normal text-black transition-opacity duration-300 ease-in-out focus:bg-neutral-200 focus:outline-none disabled:select-none disabled:hover:cursor-not-allowed"
			class:opacity-100={isShow}
			class:opacity-70={!isShow}
			value={value ?? ''}
			on:input={handleInput}
			id={name}
			{disabled}
		/>
		<label
			class="label bg-neutral-100 px-4 py-1 text-neutral-500 peer-focus:bg-neutral-200 peer-focus:text-black"
			for={name}
		>
			<span
				class="label-text font-semibold text-inherit {required
					? "-ml-px -mt-0.5 after:text-xs after:text-red-500 after:content-['_*']"
					: ''}"
			>
				{viName}
			</span>
		</label>
		<button
			type="button"
			class="absolute right-2 top-[calc(50%+0.75rem)] -translate-y-1/2"
			on:click={toggleVisibility}
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				width="24px"
				height="24px"
				viewBox="0 0 1024 1024"
				class="transform transition-transform duration-300 ease-in-out"
				class:rotate-0={isShow}
				class:rotate-180={!isShow}
				class:opacity-100={isShow}
				class:opacity-70={!isShow}
			>
				{#if isShow}
					<path
						fill="currentColor"
						d="M942.2 486.2C847.4 286.5 704.1 186 512 186c-192.2 0-335.4 100.5-430.2 300.3a60.3 60.3 0 0 0 0 51.5C176.6 737.5 319.9 838 512 838c192.2 0 335.4-100.5 430.2-300.3c7.7-16.2 7.7-35 0-51.5M512 766c-161.3 0-279.4-81.8-362.7-254C232.6 339.8 350.7 258 512 258s279.4 81.8 362.7 254C791.5 684.2 673.4 766 512 766m-4-430c-97.2 0-176 78.8-176 176s78.8 176 176 176s176-78.8 176-176s-78.8-176-176-176m0 288c-61.9 0-112-50.1-112-112s50.1-112 112-112s112 50.1 112 112s-50.1 112-112 112"
					/>
				{:else}
					<path
						fill="currentColor"
						d="M942.2 486.2Q889.47 375.11 816.7 305l-50.88 50.88C807.31 395.53 843.45 447.4 874.7 512C791.5 684.2 673.4 766 512 766q-72.67 0-133.87-22.38L323 798.75Q408 838 512 838q288.3 0 430.2-300.3a60.29 60.29 0 0 0 0-51.5m-63.57-320.64L836 122.88a8 8 0 0 0-11.32 0L715.31 232.2Q624.86 186 512 186q-288.3 0-430.2 300.3a60.3 60.3 0 0 0 0 51.5q56.69 119.4 136.5 191.41L112.48 835a8 8 0 0 0 0 11.31L155.17 889a8 8 0 0 0 11.31 0l712.15-712.12a8 8 0 0 0 0-11.32M149.3 512C232.6 339.8 350.7 258 512 258c54.54 0 104.13 9.36 149.12 28.39l-70.3 70.3a176 176 0 0 0-238.13 238.13l-83.42 83.42C223.1 637.49 183.3 582.28 149.3 512m246.7 0a112.11 112.11 0 0 1 146.2-106.69L401.31 546.2A112 112 0 0 1 396 512"
					/><path
						fill="currentColor"
						d="M508 624c-3.46 0-6.87-.16-10.25-.47l-52.82 52.82a176.09 176.09 0 0 0 227.42-227.42l-52.82 52.82c.31 3.38.47 6.79.47 10.25a111.94 111.94 0 0 1-112 112"
					/>
				{/if}
			</svg>
		</button>
	</div>
	{#if error !== null}
		{#each error as e}
			<Error>
				{e}
			</Error>
		{/each}
	{/if}
</div>
