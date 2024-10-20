<script lang="ts">
	import RefreshIcon from '~icons/ri/refresh-line'
	import { invalidate } from '$app/navigation'
	import RolesList from './RolesList.svelte'
	import AddRole from './AddRole.svelte'
	import type { PageData } from './$types'
	let drawerToggleRef: HTMLInputElement

	export let data: PageData
	function refreshData() {
		invalidate('app:students')
	}

	const tabData = [
		{ name: 'Danh sách vai trò đã có', section: 'roleList', value: 0 },
		{ name: 'Tạo mới vai trò', section: 'newRole', value: 1 }
	]

	let activeTabValue = 0
	function handleClick(tabValue: number) {
		activeTabValue = tabValue
	}

	function handleKeydown(event: KeyboardEvent, tabValue: number) {
		if (event.key === 'Enter' || event.key === ' ') {
			event.preventDefault()
			handleClick(tabValue)
		}
	}
</script>

<div class="drawer drawer-end h-full">
	<input id="my-drawer" type="checkbox" class="drawer-toggle" bind:this={drawerToggleRef} />
	<div class="drawer-content">
		<header class="mb-5 flex items-center justify-between">
			<div class="flex items-center gap-5">
				<div class="breadcrumbs text-sm">
					<ul>
						<li>Admin</li>
						<li>Quản lý vai trò</li>
					</ul>
				</div>
				<button class="btn btn-circle btn-ghost btn-sm active:!rotate-180" on:click={refreshData}>
					<RefreshIcon />
				</button>
			</div>
		</header>

		<div class="tabs tab-lg w-full">
			<ul class="tabs tab-lg w-full border-b border-gray-400">
				{#each tabData as item}
					<li class={` tab ${activeTabValue === item.value ? 'tab-active tab-bordered' : ''}`}>
						<span
							role="button"
							tabindex="0"
							on:keydown={(event) => handleKeydown(event, item.value)}
							class="mb-1 font-bold"
							on:click={() => {
								handleClick(item.value)
							}}
						>
							{item.name}
						</span>
					</li>
				{/each}
			</ul>
		</div>

		<div class="mt-4 h-full">
			{#if activeTabValue === 0}
				<RolesList {data} />
			{:else if activeTabValue === 1}
				<AddRole />
			{/if}
		</div>
	</div>
</div>
