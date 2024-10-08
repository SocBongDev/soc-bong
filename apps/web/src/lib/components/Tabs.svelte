<script lang="ts">
	type item = { name: string; value: number; section: string; component: any }
	export let items: item[] = []
	export let activeTabValue: number = 1
	function handleClick(tabValue: number) {
		activeTabValue = tabValue
	}

	function handleDropdown(tabValue: number, event: KeyboardEvent) {
		if (event.key === 'Tab' || event.key === 'Enter' || event.key === 'ArrowRight') {
			activeTabValue = tabValue
		}
	}
</script>

<section id="students" class="h-full w-full">
	<ul class="tabs tab-lg border-b border-gray-400">
		{#each items as item}
			<li class={` tab ${activeTabValue === item.value ? 'tab-active tab-bordered' : ''}`}>
				<span
					class="mb-1 font-bold"
					on:click={() => handleClick(item.value)}
					on:keydown={(e) => handleDropdown(item.value, e)}
					role="button"
					tabindex="0"
				>
					{item.name}
				</span>
			</li>
		{/each}
	</ul>

	{#each items as item}
		{#if activeTabValue === item.value}
			<svelte:component this={item.component} />
		{/if}
	{/each}
</section>
