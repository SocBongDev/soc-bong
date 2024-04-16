<script lang="ts">
	import { text } from '@sveltejs/kit'

	let isDropdownOpen: boolean = false

	const status = [
		{ name: 'Có mặt', color: 'bg-green-500', letter: '🟢' },
		{ name: 'Vắng', color: 'bg-red-500', letter: '🔴' },
		{ name: 'Nghỉ', color: 'bg-gray-500', letter: '🟡' }
	]

	$: activeStatus = status[0]

	let active: string | null = ''

	function handleDropdownClick() {
		isDropdownOpen = !isDropdownOpen
	}

	function handleChangeState(index: number) {
		activeStatus = status[index]
		isDropdownOpen = false
	}

	function handleDropdownFocusLoss({
		relatedTarget,
		currentTarget
	}: {
		relatedTarget: EventTarget | null
		currentTarget: HTMLElement
	}) {
		if (relatedTarget instanceof HTMLElement && !currentTarget.contains(relatedTarget)) {
			isDropdownOpen = false
		}
	}
</script>

<div class="dropdown" on:focusout={handleDropdownFocusLoss}>
	<button class="btn btn-square btn-xs align-middle" on:click={handleDropdownClick}>
		{#if active == ''}
			<svg class="w-10 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"
				><path
					d="M21.64,13a1,1,0,0,0-1.05-.14,8.05,8.05,0,0,1-3.37.73A8.15,8.15,0,0,1,9.08,5.49a8.59,8.59,0,0,1,.25-2A1,1,0,0,0,8,2.36,10.14,10.14,0,1,0,22,14.05,1,1,0,0,0,21.64,13Zm-9.5,6.69A8.14,8.14,0,0,1,7.08,5.22v.27A10.15,10.15,0,0,0,17.22,15.63a9.79,9.79,0,0,0,2.1-.22A8.11,8.11,0,0,1,12.14,19.73Z"
				/></svg
			>
		{:else}
			<span class="max-h-full w-fit">{active}</span>
		{/if}
	</button>
	<ul
		class="dropdown-content menu bg-base-100 rounded-box z-10 p-2 shadow"
		style:visibility={isDropdownOpen ? 'visible' : 'hidden'}
	>
		{#each status as { name, color, letter }, index (letter)}
			<button
				class="flex items-center"
				on:click={(event) => {
					handleChangeState(index)
					active = event.currentTarget.textContent
				}}
			>
				<span class="btn btn-square text-gray-400">{letter}</span>
			</button>
		{/each}
	</ul>
</div>
