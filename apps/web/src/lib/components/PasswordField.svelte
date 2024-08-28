<script lang="ts">
	import Error from './FloatingLabel/Error.svelte'

	export let name: string
	export let required = false
	export let value: string | undefined = undefined
	export let error: string[] | null = null
    export let disabled: boolean = false
	let isPasswordField = name === 'password' || name === 'confirmPassword'
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
            class="peer w-full bg-neutral-100 px-4 pb-2.5 pt-1 font-normal text-black focus:bg-neutral-200 focus:outline-none disabled:hover:cursor-not-allowed disabled:select-none"
            value={value ?? ''}
            on:input={handleInput}
            id={name}
	        disabled={disabled}
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
                {name}
            </span>
        </label>
        <button
            type="button"
            class="absolute right-2 top-[calc(50%+0.75rem)] -translate-y-1/2"
            on:click={toggleVisibility}
        >
            {isShow ? '👁️' : '👁️‍🗨️'}
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
