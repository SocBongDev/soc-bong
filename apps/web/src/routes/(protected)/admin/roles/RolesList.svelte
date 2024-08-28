<script lang="ts">
	import { onMount } from 'svelte'
	import ArrowRightIcon from '~icons/formkit/arrowright'
	import EllipsisIcon from '~icons/fa6-solid/ellipsis'
	import type { PageData } from './$types'

	export let data: PageData

	let isChecked: string[] = []
	let isCheckedAll = false
	let checked: boolean

	function handleCheck(e: any) {
		const { id, checked } = e.currentTarget

		if (!checked) {
			const isValidCheckAll = isChecked.length === data.roles.data.length
			if (isValidCheckAll) {
				isCheckedAll = false
			}

			isChecked = isChecked.filter((item) => item !== id)
			return
		}

		isChecked = [...isChecked, id]
		const isValidCheckAll = isChecked.length === data?.roles?.data.length
		if (isValidCheckAll) {
			isCheckedAll = true
		}
	}

	function handleCheckAll() {
		isCheckedAll = !isCheckedAll
		if (!isCheckedAll) {
			isChecked = []
			return
		}

		isChecked = data.roles?.data?.map((el: any) => el?.id.toString())
	}

	function clearSelected() {
		isCheckedAll = false
		isChecked = []
	}
</script>

<div class="relative flex h-full flex-col gap-10 overflow-x-auto">
	<table class="table">
		<thead>
			<tr>
				<th>
					<label>
						<input
							type="checkbox"
							class="checkbox checkbox-sm rounded"
							checked={isCheckedAll}
							on:click={handleCheckAll}
						/>
					</label>
				</th>
				<th>Id của vai trò</th>
				<th>Tên vai trò</th>
				<th>Miêu tả vai trò</th>
				<th>
					<button class="btn btn-square btn-ghost btn-sm active:!translate-y-1">
						<EllipsisIcon />
					</button>
				</th>
			</tr>
		</thead>
		<tbody>
			{#each data.roles.data as role (role.id)}
				{#if role.id}
					<tr class="hover cursor-pointer">
						<th>
							<label>
								<input
									id={role.id?.toString()}
									type="checkbox"
									class="checkbox checkbox-sm rounded"
									on:click={handleCheck}
									checked={isChecked.includes(role.id?.toString())}
								/>
							</label>
						</th>
						<th on:click={() => console.log(role.id)}>{role.id}</th>
						<td on:click={() => console.log(role.id)}
							>{role.name}</td
						>
						<td on:click={() => console.log(role.id)}>{role.description}</td>
						<td on:click={() => console.log(role.id)}>
							<div class="px-2">
								<ArrowRightIcon />
							</div>
						</td>
					</tr>
				{/if}
			{/each}
		</tbody>
	</table>
</div>
