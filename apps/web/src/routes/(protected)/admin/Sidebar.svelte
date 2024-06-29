<script lang="ts">
	import type { SvelteComponent } from 'svelte'
	import StudentIcon from '~icons/ph/student-bold'
	import RegisterIcon from '~icons/mdi/register'
	import RoleIcon from '~icons/eos-icons/admin-outlined'
	import SchoolIcon from '~icons/teenyicons/school-solid'
	import UserIcon from '~icons/fa-solid/user-cog'

	type SidebarData = {
		icon: typeof SvelteComponent
		children: string
		isDisabled?: boolean
		name: string
	}
	const sidebarData: SidebarData[] = [
		{ children: 'Danh sách đăng ký', icon: RegisterIcon, name: 'registrations' },
		{ children: 'Danh sách học viên', icon: StudentIcon, isDisabled: false, name: 'students' },
		{ children: 'Quản lý vai trò', icon: RoleIcon, isDisabled: true, name: 'roles'},
		{ children: 'Quản lý người dùng', icon: UserIcon, isDisabled: true, name: 'users'},
		{ children: 'Quản lý cơ sở', icon: SchoolIcon, isDisabled: false, name: 'agencies'}
	]
</script>

<aside
	id="default-sidebar"
	class="fixed left-0 top-0 z-10 h-screen w-64 -translate-x-full transition-transform sm:translate-x-0"
	aria-label="Sidebar"
>
	<div class="flex h-full flex-col justify-between overflow-y-auto bg-gray-50 px-3 py-4">
		<slot name="header" />
		<ul class="menu gap-2">
			{#each sidebarData as { children, icon, isDisabled, name } (crypto.randomUUID())}
				<li class={isDisabled ? 'disabled' : ''}>
					<a class="{!isDisabled ? 'active' : 'pointer-events-none'} rounded" href="/admin/{name}">
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
		<slot name="footer" />
	</div>
</aside>
