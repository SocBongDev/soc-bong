<script lang="ts">
	import PlusIcon from '~icons/ic/round-add'
	import ArrowLeftIcon from '~icons/formkit/arrowleft'
	import { fade, blur, fly, slide, scale } from 'svelte/transition'
	let activeAdd = false

	function toggleAdd() {
		activeAdd = !activeAdd
	}
</script>

{#if activeAdd}
	<div class="relative mb-2 h-full w-full" transition:fade={{ duration: 300 }}>
		<div class="grid grid-cols-5 items-center">
			<button class="btn btn-circle col-span-2 text-left" on:click={toggleAdd}>
				<ArrowLeftIcon class="font-bold" />
			</button>
			<div class="col-span-3 mt-4 h-full w-full text-left">
				<h1 class="text-2xl font-bold">Thêm mới vai trò</h1>
			</div>
		</div>

		<div class="mt-4 w-full">
			<h1 class="mt-2 text-xl font-bold">1. Đã tạo trước đó</h1>
			<div
				class="mt-2 grid h-full grid-cols-3 place-items-center rounded-xl border border-black/10 bg-black/5 p-4"
			>
				<div class="group card w-fit min-w-[9rem] rounded-xl p-4 shadow-md hover:bg-black">
					<figure>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="48"
							height="48"
							viewBox="0 0 24 24"
							class="group-hover:text-white"
							><path
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								d="M8 11A5 5 0 1 0 8 1a5 5 0 0 0 0 10Zm5.023 2.023C11.772 11.76 10.013 11 8 11c-4 0-7 3-7 7v5h7m2-3.5a2.5 2.5 0 1 0 5.002-.002A2.5 2.5 0 0 0 10 19.5ZM23 15l-3-3l-6 6m3.5-3.5l3 3l-3-3Z"
							/></svg
						>
					</figure>
					<div class="card-body p-4">
						<h2 class="card-title justify-center text-lg group-hover:text-white">User</h2>
					</div>
				</div>

				<div class="group card w-fit min-w-[9rem] rounded-xl p-4 shadow-md hover:bg-black">
					<figure>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="48"
							height="48"
							viewBox="0 0 24 24"
							class="group-hover:text-white"
							><path
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								d="M8 11A5 5 0 1 0 8 1a5 5 0 0 0 0 10Zm5.023 2.023C11.772 11.76 10.013 11 8 11c-4 0-7 3-7 7v5h7m2-3.5a2.5 2.5 0 1 0 5.002-.002A2.5 2.5 0 0 0 10 19.5ZM23 15l-3-3l-6 6m3.5-3.5l3 3l-3-3Z"
							/></svg
						>
					</figure>
					<div class="card-body p-4">
						<h2 class="card-title justify-center text-lg group-hover:text-white">User</h2>
					</div>
				</div>

				<div class="group card w-fit min-w-[9rem] rounded-xl p-4 shadow-md hover:bg-black">
					<figure>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="48"
							height="48"
							viewBox="0 0 24 24"
							class="group-hover:text-white"
							><path
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								d="M8 11A5 5 0 1 0 8 1a5 5 0 0 0 0 10Zm5.023 2.023C11.772 11.76 10.013 11 8 11c-4 0-7 3-7 7v5h7m2-3.5a2.5 2.5 0 1 0 5.002-.002A2.5 2.5 0 0 0 10 19.5ZM23 15l-3-3l-6 6m3.5-3.5l3 3l-3-3Z"
							/></svg
						>
					</figure>
					<div class="card-body p-4">
						<h2 class="card-title justify-center text-lg group-hover:text-white">User</h2>
					</div>
				</div>
			</div>
			<h1 class="mt-2 text-xl font-bold">2. Tạo thêm vai trò</h1>
			<div class="mt-2 h-fit rounded-xl border border-black/10 bg-black/5 p-4">
				<div class="flex flex-col justify-start gap-1">
					<label for="role">Role: </label>
					<div class="flex flex-row items-center justify-start gap-4">
						<input name="role" class="input border focus:border-black focus:outline-none" />
						<button class="btn btn-outline">Add</button>
					</div>
				</div>
			</div>
		</div>
	</div>
{:else if !activeAdd}
	<div
		class="relative flex h-full w-full flex-col items-center justify-center"
		transition:slide={{ duration: 300 }}
	>
		<button class="btn btn-square btn-outline btn-lg rounded-xl border-2" on:click={toggleAdd}>
			<PlusIcon class="h-12 w-12" />
		</button>
		<p class="mt-3 text-2xl font-semibold">Thêm vai trò mới</p>
	</div>
{/if}
