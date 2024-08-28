<script lang="ts">
	import type { UserProps } from '$lib/common/type'
	import { PUBLIC_API_SERVER_URL } from '$env/static/public'
	import dayjs from 'dayjs'
	import type { PageData } from './$types'
	import PlusIcon from '~icons/ic/round-add'
	import EllipsisIcon from '~icons/fa6-solid/ellipsis'
	import ArrowRightIcon from '~icons/formkit/arrowright'
	import RefreshIcon from '~icons/ri/refresh-line'
	import TimesIcon from '~icons/uil/times'
	import TrashIcon from '~icons/lucide/trash'
	import TextField from '$lib/components/TextField.svelte'
	import DateField from '$lib/components/DateField.svelte'
	import SelectField from '$lib/components/SelectField.svelte'
	import { createForm } from 'felte'
	import { CreateUserSchema as schema } from './validate'
	import { validator } from '@felte/validator-zod'
	import { dialogProps, Notify, openDialog } from '$lib/store'
	import { invalidate } from '$app/navigation'
	import { blur, fade } from 'svelte/transition'
	import PasswordField from '$lib/components/PasswordField.svelte'
	export let data: PageData
	let drawerToggleRef: HTMLInputElement
	let isChecked: string[] = []
	let scrollClass = ''
	let isCheckedAll = false
	let isNew = true

	let recordData: UserProps | null = null
	let checked: boolean
	let loading = false
	let abortController: AbortController | undefined = undefined
	$: isNew = !recordData
	$: if (recordData !== null) {
		const { id, createdAt, ...initialValues } = recordData
		setInitialValues(initialValues)
		reset()
	}

	const token = localStorage.getItem('access_token')

	var agencyOptions = data?.agencies?.data?.map((el) => ({
		label: el.name,
		value: el.id ? el.id.toString() : '1'
	}))

	const defaultFormValues: UserProps = {
		email: '',
		first_name: '',
		last_name: '',
		password: '',
		connection: 'Username-Password-Authentication',
		phone_number: '',
		dob: dayjs().format('DD-MM-YYYY'),
		agencyId: parseInt(agencyOptions[0]?.value || '1'),
		auth0_user_id: ''
	}

	let formData: Record<string, string> = {}
	$: userSchema.forEach((field) => {
		if (!(field.name in formData)) {
			formData[field.name] = ''
		}
	})

	function formatAgencyName(agencyId: number) {
		const agency = data.agencies.data.find(
			(el) => (el.id && parseInt(el.id?.toString())) === agencyId
		)
		return agency?.name || 'N/A'
	}

	function handleCheck(e: any) {
		const { id, checked } = e.currentTarget

		if (!checked) {
			const isValidCheckAll = isChecked.length === data.users.data.length
			if (isValidCheckAll) {
				isCheckedAll = false
			}

			isChecked = isChecked.filter((item) => item !== id)
			return
		}

		isChecked = [...isChecked, id]
		const isValidCheckAll = isChecked.length === data?.users?.data.length
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
		isChecked = data.users?.data?.map((el: any) => el?.id.toString())
	}

	function resetDefaultForm() {
		setInitialValues(defaultFormValues)
		reset()
	}

	const { form, errors, setInitialValues, reset } = createForm({
		debounced: { timeout: 500 },
		extend: validator({ schema }),
		transform: (values: any) => ({
			...values,
			agencyId: parseInt(values.agencyId, 10)
		}),
		onSubmit: save
	})

	async function save(req: UserProps) {
		loading = true
		const body = {
			...req,
			agencyId: req.agencyId,
			first_name: req.first_name,
			last_name: req.last_name,
			phone_number: req.phone_number,
			dob: req.dob,
			connection: 'Username-Password-Authentication',
		}
		const bodyFormated = JSON.stringify(body)
		const method = isNew ? 'POST' : 'PUT'
		const url = isNew
			? `${PUBLIC_API_SERVER_URL}/users`
			: `${PUBLIC_API_SERVER_URL}/users/${recordData?.id}`

		const request = fetch(url, {
			method,
			headers: {
				'Content-Type': 'application/json',
				accept: 'application/json',
				Authorization: `Bearer ${token}`
			},
			body: bodyFormated
		}).then((res) => {
			if (res.status == 422) {
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'phía server đã tồn tại dữ liệu này!'
				})
			} else if (res.status == 403 || res.status == 401) {
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'Người dùng hiện không có quyền thực hiện này!'
				})
			}
		})

		try {
			const res = await request
			refreshData()
			resetDefaultForm()
			Notify({
				type: "success",
				id: crypto.randomUUID(),
				description: isNew ? 'Tạo người dùng thành công!' : 'Cập nhật thành công!'
			})
			hide()
		} catch (e) {
			console.error('Save Error: ', e)
			Notify({
				type: 'error',
				id: crypto.randomUUID(),
				description: 'Lỗi từ phía server'
			})
		}
		loading = false
	}

	async function loadData(id: number, signal: AbortSignal) {
		loading = true

		try {
			const res = await fetch(`${PUBLIC_API_SERVER_URL}/users/${id}`, {
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json',
					accept: 'application/json'
				},
				signal
			}).then((res) => res.json())
			res.dob = dayjs(res?.dob).format('YYYY-MM-DD')
			recordData = res
		} catch (e: any) {
			console.error('LoadData: ', e)
			if (e.name !== undefined && e.name === 'AbortError') {
				return
			}

			Notify({ description: 'Lỗi từ phía server', id: crypto.randomUUID(), type: 'error' })
		} finally {
			loading = false
		}
	}

	function refreshData() {
		invalidate('app:users')
	}

	let prevPromise: Promise<void>
	async function show(id?: number) {
		checked = true
		drawerToggleRef.checked = true

		if (abortController) {
			abortController.abort()
		}

		if (id !== undefined) {
			if (prevPromise) {
				await prevPromise
			}
			abortController = new AbortController()
			prevPromise = loadData(id, abortController.signal)
		}
	}

	function clearSelected() {
		isChecked = []
		isCheckedAll = false
	}

	async function batchDelete() {}

	function handleContentScroll(panel: HTMLElement) {
		const heightDiff = panel.scrollHeight - panel.offsetHeight
		if (heightDiff > 0) {
			scrollClass = 'scrolled'
		}

		if (panel.scrollTop === 0) {
			scrollClass = 'scroll-reached-top'
		}

		if (panel.scrollTop + panel.offsetHeight === panel.scrollHeight) {
			scrollClass = 'scroll-reached-bottom'
		}
	}

	const userSchema: {
		name: string
		type: 'text' | 'date' | 'select' | 'password'
		required: boolean
		options?: { label: string; value: string }[]
		disabled?: boolean
	}[] = [
		{
			name: 'email',
			type: 'text',
			required: false,
			disabled: true
		},
		{
			name: 'first_name',
			type: 'text',
			required: true
		},
		{
			name: 'last_name',
			type: 'text',
			required: true
		},
		{
			name: 'dob',
			type: 'date',
			required: true
		},
		{
			name: 'phone_number',
			type: 'text',
			required: true
		},
		{
			name: 'password',
			type: 'text' || 'password',
			required: true,
			disabled: true,
		},
		{
			name: 'agencyId',
			type: 'select',
			required: true,
			options: agencyOptions
		}
	]

	function hide() {
		checked = false
		drawerToggleRef.checked = false
		resetDefaultForm()
	}
</script>

<div class="drawer drawer-end h-full">
	<input
		id="my-drawer"
		type="checkbox"
		class="drawer-toggle"
		bind:this={drawerToggleRef}
		on:change={(e) => {
			if (!e.currentTarget.checked) {
				recordData = null
				checked = false
				resetDefaultForm()
			} else {
				checked = true
			}
		}}
	/>
	<div class="drawer-content">
		<header class="mb-5 flex items-center justify-between">
			<div class="flex items-center gap-5">
				<div class="breadcrumbs text-sm">
					<ul>
						<li>Admin</li>
						<li>Danh sách người dùng</li>
					</ul>
				</div>
				<button class="btn btn-circle btn-ghost btn-sm active:!rotate-180" on:click={refreshData}>
					<RefreshIcon />
				</button>
			</div>

			<button
				class="btn btn-primary btn-sm rounded normal-case active:!translate-y-1"
				on:click={() => {
					recordData = null
					show()
				}}
			>
				<PlusIcon />
				Thêm mới
			</button>
		</header>

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
						<th>Email người dùng </th>
						<th>Tên người dùng</th>
						<th>Ngày sinh</th>
						<th>Số điện thoại</th>
						<th>Cơ sở hiện tại</th>
						<th>User Id </th>
						<th>
							<button class="btn btn-square btn-ghost btn-sm active:!translate-y-1">
								<EllipsisIcon />
							</button>
						</th>
					</tr>
				</thead>
				<tbody>
					{#each data.users.data as user (user.id)}
						{#if user.id}
							<tr class="hover cursor-pointer">
								<th>
									<label>
										<input
											id={user.id?.toString()}
											type="checkbox"
											class="checkbox checkbox-sm rounded"
											on:click={handleCheck}
											checked={isChecked.includes(user.id?.toString())}
										/>
									</label>
								</th>
								<th on:click={() => show(user.id)}>{user.email}</th>
								<th on:click={() => show(user.id)}>{user.first_name + ' ' + user.last_name}</th>
								<td on:click={() => show(user.id)}>{dayjs(user.dob).format('DD/MM/YYYY')}</td>
								<td on:click={() => show(user.id)}>{user.phone_number}</td>
								<td on:click={() => show(user.id)}>{formatAgencyName(user.agencyId)}</td>
								<td on:click={() => show(user.id)}>{user.auth0_user_id}</td>
								<td on:click={() => show(user.id)}>
									<div class="px-2">
										<ArrowRightIcon />
									</div>
								</td>
							</tr>
						{/if}
					{/each}
				</tbody>
			</table>

			<div class="join mt-auto self-center">
				<a
					class={data.users.page === 1 ? 'pointer-events-none cursor-default opacity-40' : ''}
					href={`/admin/users?page=${data.users.page - 1}&pageSize=${data.users.pageSize}`}
				>
					<button class="btn join-item">«</button>
				</a>
				<button class="btn join-item">Trang {data.users.page}</button>
				<a
					class={data.users.data.length < data.users.pageSize || data.users.data.length === 0
						? 'pointer-events-none cursor-default opacity-40'
						: ''}
					href={`/admin/users?page=${data.users.page + 1}&pageSize=${data.users.pageSize}`}
				>
					<button class="btn join-item">»</button>
				</a>
			</div>
			{#if isChecked.length > 0}
				<div class="absolute bottom-20 left-1/2 w-1/2 -translate-x-1/2" transition:fade>
					<div class="alert flex justify-between rounded-full bg-white py-2.5 text-sm shadow">
						<div class="flex w-1/2 items-center gap-3">
							<span>Đã chọn <strong>{isChecked.length}</strong> dòng</span>
							<button
								class="btn btn-outline btn-sm rounded border-2 bg-white normal-case"
								on:click={clearSelected}>Bỏ chọn</button
							>
						</div>
						<button
							class="btn btn-ghost btn-sm rounded normal-case text-red-500 hover:bg-red-100"
							on:click={() => {
								dialogProps.set({
									description: 'Hành vi này không thể hoàn tác. Bạn có muốn tiếp tục?',
									title: 'Yêu cầu xác nhận!',
									onContinue: batchDelete
								})
								openDialog.set(true)
							}}>Xóa dữ liệu đã chọn</button
						>
					</div>
				</div>
			{/if}
		</div>
	</div>

	<div class="drawer-side z-10 place-items-center">
		<label for="my-drawer" class="drawer-overlay" />
		<div class="flex h-full w-1/2 flex-col bg-white">
			{#if loading}
				<label
					for="drawer-content"
					class="absolute inset-0 z-10 grid place-items-center bg-black/60"
					transition:blur
				>
					<span class="loading loading-spinner loading-lg text-info" />
				</label>
			{/if}
			<header class="relative flex flex-shrink-0 items-center justify-between px-7 py-6">
				<h4 class="text-lg font-bold">{isNew ? 'Tạo mới' : 'Chỉnh sửa'}</h4>
				<div class="tooltip tooltip-warning tooltip-left" data-tip="Xóa">
					<button
						class="btn btn-circle btn-ghost btn-sm active:!translate-y-1"
						on:click={() => {
							dialogProps.set({
								description: 'Hành vi này không thể hoàn tác. Bạn có muốn tiếp tục?',
								title: 'Yêu cầu xác nhận!'
								// onContinue: handleDelete
							})
							openDialog.set(true)
						}}
					>
						<TrashIcon class="text-red-500" />
					</button>
				</div>
				<button
					class="btn btn-sm absolute {checked ? 'right-full' : ''} top-5 rounded-s-full"
					on:click={hide}
				>
					<TimesIcon />
				</button>
			</header>
			<section
				class="flex-1 overflow-scroll px-7 py-6 {scrollClass}"
				on:scroll={(e) => handleContentScroll(e.currentTarget)}
			>
					<form class="flex flex-col gap-8 text-sm" id="upsertForm" use:form>
						{#each userSchema as { type, name, required, options, disabled } (name)}
							{#if type === 'text' || type === 'password'}
								{#if name == 'password' || name == 'confirmPassword'}
									<PasswordField
										error={$errors[name]}
										{name}
										{required}
										{disabled}
										bind:value={formData[name]}
									/>
								{:else}
									<TextField error={$errors[name]} {name} {required} {disabled}/>
								{/if}
							{:else if type === 'select'}
								<SelectField error={$errors[name]} {name} {options} {required} />
							{:else if type === 'date'}
								<DateField error={$errors[name]} {name} {required} />
							{/if}
						{/each}
					</form>
			</section>
			<footer class="flex flex-shrink-0 items-center justify-end px-7 py-6">
				<div class="flex items-center gap-5">
					<button class="btn btn-ghost rounded normal-case active:!translate-y-1" on:click={hide}
						>Hủy bỏ</button
					>
					<button
						class="btn rounded px-10 normal-case active:!translate-y-1"
						type="submit"
						form="upsertForm">{isNew ? 'Tạo mới' : 'Lưu thay đổi'}</button
					>
				</div>
			</footer>
		</div>
	</div>
</div>
