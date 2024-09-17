<script lang="ts">
	import { type SvelteComponent } from 'svelte'
	import StudentIcon from '~icons/ph/student-bold'
	import RegisterIcon from '~icons/mdi/register'
	import RoleIcon from '~icons/eos-icons/admin-outlined'
	import SchoolIcon from '~icons/teenyicons/school-solid'
	import UserIcon from '~icons/fa-solid/user-cog'
	import { userRoleStore, SidebarContext } from '$lib/store'
	import { page } from '$app/stores'
	import CloseIcon from '~icons/ic/round-close'
	type SidebarData = {
		icon: typeof SvelteComponent
		children: string
		isDisabled?: boolean
		name: string
	}

	$: role = $userRoleStore

	const { subscribe } = SidebarContext

	let collapseMenu = false

	subscribe((context) => {
		collapseMenu = context.collapseMenu
	})

	const setCollapseMenu = (v: boolean) => {
		SidebarContext.update((context) => ({
			...context,
			collapseMenu: v
		}))
	}

	function handleToggleCollapse() {
		setCollapseMenu(!collapseMenu)
		console.log('check collapseMenu: ', collapseMenu)
	}

	const sidebarData: SidebarData[] = [
		{ children: 'Danh sách đăng ký', icon: RegisterIcon, name: 'registrations' },
		{ children: 'Danh sách học viên', icon: StudentIcon, isDisabled: false, name: 'students' },
		{
			children: 'Quản lý vai trò',
			icon: RoleIcon,
			isDisabled: role === 'teacher' || localStorage.getItem('role') === 'teacher' ? true : false,
			name: 'roles'
		},
		{
			children: 'Quản lý người dùng',
			icon: UserIcon,
			isDisabled: role === 'teacher' || localStorage.getItem('role') === 'teacher' ? true : false,
			name: 'users'
		},
		{
			children: 'Quản lý cơ sở',
			icon: SchoolIcon,
			isDisabled: role === 'teacher' || localStorage.getItem('role') === 'teacher' ? true : false,
			name: 'agencies'
		},
		{ children: 'Quản lý lớp học', icon: SchoolIcon, isDisabled: false, name: 'classes' }
	]
</script>

<aside
	id="layout-sidebar"
	class="fixed left-0 top-0 z-10 h-screen w-64 sm:translate-x-0 transition-transform { collapseMenu ? 'translate-x-0' : '-translate-x-full'}"
	aria-label="Sidebar"
>
	<div class="flex h-full flex-col justify-between overflow-y-auto bg-gray-50 px-3 py-4">
		<slot name="header" />
		<div class="h-full w-full">
			<div class="flex flex-row items-center justify-end gap-4 px-2 sm:hidden">
				<button 
					class="ml-3 mt-2 inline-flex items-center rounded-lg p-2 text-sm text-gray-700 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200"
					on:click={handleToggleCollapse}	
				>
					<CloseIcon />
				</button>
			</div>
			<ul class="menu gap-2">
				{#each sidebarData as { children, icon, isDisabled, name } (crypto.randomUUID())}
					<li class={isDisabled ? 'disabled' : ''}>
						<a
							class="{!isDisabled ? 'active' : 'pointer-events-none'} rounded {$page.url
								.pathname === `/admin/${name}`
								? 'my-2 !translate-x-2 !border !bg-white !text-black'
								: 'bg-black'}"
							href="/admin/{name}"
							on:click={handleToggleCollapse}
						>
							{#if isDisabled}
								<div
									class="tooltip tooltip-info tooltip-top p-0"
									data-tip="Tính năng đang phát triển"
								>
									<span class="flex items-center gap-2">
										<svelte:component this={icon} />
										<span class="ml-3">{children}</span>
									</span>
								</div>
							{:else}
								<svelte:component this={icon} />
								<span class="ml-3">{children}</span>
							{/if}
						</a>
					</li>
				{/each}
			</ul>
		</div>
		<slot name="footer" />
	</div>
</aside>
