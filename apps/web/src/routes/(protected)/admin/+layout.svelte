<script lang="ts">
	import { goto } from '$app/navigation'
	import Sidebar from './Sidebar.svelte'
	import LogoutIcon from '~icons/ri/logout-circle-line'
	import ProfileIcon from '~icons/gg/profile'
	import { dialogProps, Notify, openDialog, SidebarContext } from '$lib/store'

	let modalRef: HTMLDialogElement | undefined
	$: if ($openDialog) {
		showModal()
	} else {
		closeModal()
	}

	async function handleSignOut() {
		if (localStorage.getItem('access_token')) {
			localStorage.removeItem('access_token')
			localStorage.removeItem('id_token')
			localStorage.removeItem('expires_at')
			return goto('/')
		} else {
			Notify({
				type: 'error',
				id: crypto.randomUUID(),
				description: 'Chưa thể thực hiện đăng xuất lúc này, vui lòng thử lại sau!'
			})
		}
	}

	function showModal() {
		modalRef?.showModal()
	}

	function closeModal() {
		modalRef?.close()
	}

	function openSidebar() {
		SidebarContext.update((context) => ({
			...context,
			collapseMenu: true
		}))
	}
</script>

<button
	data-drawer-target="layout-sidebar"
	data-drawer-toggle="layout-sidebar"
	aria-controls="layout-sidebar"
	type="button"
	class="ml-3 mt-2 inline-flex items-center rounded-lg p-2 text-sm text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 sm:hidden"
	on:click={openSidebar}
>
	<span class="sr-only">Open sidebar</span>
	<svg
		class="h-6 w-6"
		aria-hidden="true"
		fill="currentColor"
		viewBox="0 0 20 20"
		xmlns="http://www.w3.org/2000/svg"
	>
		<path
			clip-rule="evenodd"
			fill-rule="evenodd"
			d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"
		/>
	</svg>
</button>

<Sidebar>
	<footer slot="footer">
		<div class="dropdown dropdown-top">
			<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
			<label for="" tabindex="0" class="avatar btn btn-circle btn-ghost">
				<div class="w-10 rounded-full">
					<img src="https://api.dicebear.com/6.x/initials/svg?seed=" alt="User avatar" />
				</div>
			</label>
			<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
			<ul
				tabindex="0"
				class="menu dropdown-content rounded-box menu-sm z-50 mt-3 w-60 bg-base-100 p-2 shadow"
			>
				<li class="disabled p-2">
					<a class="justify-between rounded" href="/">
						<div class="tooltip tooltip-info tooltip-top p-0" data-tip="Tính năng đang phát triển">
							<span class="flex items-center gap-2">
								<ProfileIcon />
								Thông tin cá nhân
								<!-- <span class="badge">New</span> -->
							</span>
						</div>
					</a>
				</li>
				<li class="p-2">
					<button class="rounded" on:click={handleSignOut}>
						<LogoutIcon />
						Đăng xuất</button
					>
				</li>
			</ul>
		</div>
	</footer>
</Sidebar>

<div class="h-screen p-4 sm:ml-64">
	<div class="h-full rounded-lg p-4">
		<dialog id="delete_modal" class="modal" bind:this={modalRef}>
			<div class="modal-box">
				<h3 class="text-lg font-bold">{$dialogProps?.title}</h3>
				<p class="py-4">{$dialogProps?.description}</p>
				<div class="modal-action">
					<form class="flex items-center gap-5" method="dialog">
						<button
							class="btn btn-ghost rounded normal-case active:!translate-y-1"
							on:click={(e) => {
								$dialogProps?.onClose?.(e)
								openDialog.set(false)
								dialogProps.set(undefined)
							}}>{$dialogProps?.cancelLable ?? 'Đóng'}</button
						>
						<button
							class="btn rounded bg-red-500 px-10 normal-case text-white hover:bg-red-400 active:!translate-y-1"
							on:click={(e) => {
								$dialogProps?.onContinue?.(e)
								openDialog.set(false)
							}}>{$dialogProps?.okLabel ?? 'Xác nhận'}</button
						>
					</form>
				</div>
			</div>
		</dialog>
		<slot />
	</div>
</div>
