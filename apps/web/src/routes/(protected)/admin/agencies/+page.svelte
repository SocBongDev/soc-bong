<script lang="ts">
	import { createForm } from 'felte'
	import type { PageData } from './$types'
	import { validator } from '@felte/validator-zod'
	import { CreateAgencySchema as schema } from './validate'
	import PlusIcon from '~icons/ic/round-add'
	import EllipsisIcon from '~icons/fa6-solid/ellipsis'
	import ArrowRightIcon from '~icons/formkit/arrowright'
	import RefreshIcon from '~icons/ri/refresh-line'
	import TimesIcon from '~icons/uil/times'
	import TrashIcon from '~icons/lucide/trash'
	import TextField from '$lib/components/TextField.svelte'
	import DateField from '$lib/components/DateField.svelte'
	import SelectField from '$lib/components/SelectField.svelte'
	import { invalidate } from '$app/navigation'
	import { dialogProps, Notify, openDialog } from '$lib/store'
	import { blur, fade } from 'svelte/transition'
	import { onMount } from 'svelte'
	import type { AgencyProps } from '../type'
	import dayjs from 'dayjs'
	import { PUBLIC_API_SERVER_URL } from '$env/static/public'

	export let data: PageData
	let drawerToggleRef: HTMLInputElement
	let isChecked: string[] = []
	let scrollClass = ''
	let isCheckedAll = false
	let isNew = true
	let recordData: AgencyProps | null = null
	let checked: boolean
	let loading = false
	let abortController: AbortController | undefined = undefined
	const token = localStorage.getItem('access_token')

	$: isNew = !recordData
	$: if (recordData !== null) {
		const { id, createdAt, updatedAt, ...initialValues } = recordData
		setInitialValues(initialValues)
		reset()
	}

	const defaultFormValues = {
		name: '',
		address: '',
		phone: '',
		email: '',
		createdAt: dayjs().format('dd/MM/YYYY'),
		updatedAt: dayjs().format('dd/MM/YYYY')
	}

	function resetDefaultForm() {
		setInitialValues(defaultFormValues)
		reset()
	}

	function handleCheck(e: any) {
		const { id, checked } = e.currentTarget

		if (!checked) {
			const isValidCheckAll = isChecked.length === data?.agencies?.data.length
			if (isValidCheckAll) {
				isCheckedAll = false
			}

			isChecked = isChecked.filter((item) => item !== id)
			return
		}

		isChecked = [...isChecked, id]
		const isValidCheckAll = isChecked.length === data?.agencies?.data.length
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

		isChecked = data.agencies?.data?.map((el: any) => el?.id.toString())
	}

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

	function hide() {
		checked = false
		drawerToggleRef.checked = false
		resetDefaultForm()
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

	async function save(req: AgencyProps) {
		loading = true
		const body = JSON.stringify(req)
		const method = isNew ? 'POST' : 'PUT'
		const url = isNew
			? `${PUBLIC_API_SERVER_URL}/agencies`
			: `${PUBLIC_API_SERVER_URL}/agencies/${recordData?.id}`
		const request = fetch(url, {
			method,
			headers: {
				Authorization: `Bearer ${token}`,
				'Content-Type': 'application/json',
				accept: 'application/json'
			},
			body
		}).then((res) => {
			if (res.status == 422) {
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'phía server đã tồn tại dữ liệu này!'
				})
			} else if (res.status == 403) {
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'Người dùng hiện không có quyền thực hiện này!'
				})
			}
		})
		// note: all 4 fields need to be unique when insert
		try {
			const res = await request
			refreshData()
			resetDefaultForm()
			hide()
		} catch (e) {
			console.error('Save error', e)
			Notify({ type: 'error', id: crypto.randomUUID(), description: 'Lỗi từ phía server' })
		}

		loading = false
	}

	function refreshData() {
		invalidate('app:agencies')
	}

	async function loadData(id: number, signal: AbortSignal) {
		loading = true

		try {
			const res = await fetch(`${PUBLIC_API_SERVER_URL}/agencies/${id}`, {
				method: "GET",
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json',
					accept: 'application/json'
				},
				signal
			}).then((res) => res.json())
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

	const { form, errors, setInitialValues, reset } = createForm({
		debounced: { timeout: 500 },
		extend: validator({ schema }),
		onSubmit: save
	})

	const agencySchema: {
		name: string
		type: 'text' | 'date' | 'select'
		required: boolean
		options?: { label: string; value: string }[]
	}[] = [
		{
			name: 'name',
			type: 'text',
			required: true
		},
		{
			name: 'address',
			type: 'text',
			required: true
		},
		{
			name: 'phone',
			type: 'text',
			required: true
		},
		{
			name: 'email',
			type: 'text',
			required: true
		}
	]

	async function handleDelete() {
		if (recordData === undefined || recordData?.id === undefined) {
			return
		}

		loading = true

		try {
			const res = await fetch(`${PUBLIC_API_SERVER_URL}/agencies`, {
				body: JSON.stringify({ ids: [Number(recordData.id)] }),
				method: 'DELETE',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json',
					accept: 'application/json'
				}
			}).then((res) => res.json())
			refreshData()
			resetDefaultForm()
			hide()
			recordData = null
		} catch (e) {
			console.error(e)
			Notify({ type: 'error', id: crypto.randomUUID(), description: 'Lỗi từ phía server' })
		} finally {
			loading = false
		}
	}

	function clearSelected() {
		isChecked = []
	}

	async function batchDelete() {
		try {
			const ids = isChecked.map((el) => Number(el))
			await fetch(`${PUBLIC_API_SERVER_URL}/agencies`, {
				body: JSON.stringify({ ids }),
				method: 'DELETE',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json',
					accept: 'application/json'
				}
			})
			refreshData()
			clearSelected()
		} catch (e) {
			console.error('Batch Delete error', e)
			Notify({ type: 'error', id: crypto.randomUUID(), description: 'Lỗi từ phía server' })
		}
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
						<li>Quản Lí Cơ Sở</li>
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
						<th>Tên Cơ Sở</th>
						<th>Địa chỉ Cơ Sở</th>
						<th>Số Điện Thoại</th>
						<th>Địa chỉ Mail</th>
						<th>
							<button class="btn btn-square btn-ghost btn-sm active:!translate-y-1">
								<EllipsisIcon />
							</button>
						</th>
					</tr>
				</thead>
				<tbody>
					{#each data?.agencies?.data as agency (agency.id)}
						{#if agency.id}
							<tr class="hover cursor-pointer">
								<th>
									<label>
										<input
											id={agency.id?.toString()}
											type="checkbox"
											class="checkbox checkbox-sm rounded"
											on:click={handleCheck}
											checked={isChecked.includes(agency.id?.toString())}
										/>
									</label>
								</th>
								<th on:click={() => show(agency.id)}>{agency.name}</th>
								<td on:click={() => show(agency.id)}>{agency.address}</td>
								<td on:click={() => show(agency.id)}>{agency.phone}</td>
								<td on:click={() => show(agency.id)}>{agency.email}</td>
								<td on:click={() => show(agency.id)}>
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
					class={data?.agencies?.page === 1 ? 'pointer-events-none cursor-default opacity-40' : ''}
					href={`/admin?page=${data?.agencies?.page - 1}&pageSize=${data?.agencies?.pageSize}`}
				>
					<button class="btn join-item">«</button>
				</a>
				<button class="btn join-item">Trang {data?.agencies?.page}</button>
				<a
					class={data?.agencies?.data.length < data?.agencies?.pageSize ||
					data?.agencies?.data.length === 0
						? 'pointer-events-none cursor-default opacity-40'
						: ''}
					href={`/admin?page=${data?.agencies?.page + 1}&pageSize=${data?.agencies?.pageSize}`}
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
								title: 'Yêu cầu xác nhận!',
								onContinue: handleDelete
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
					{#each agencySchema as { type, name, required, options } (name)}
						{#if type === 'text'}
							<TextField error={$errors[name]} {name} {required} />
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
