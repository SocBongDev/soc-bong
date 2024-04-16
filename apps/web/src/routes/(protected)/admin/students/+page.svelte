<script lang="ts">
	import TimeSheet from './TimeSheet.svelte'
	import StudentList from './StudentList.svelte'
	import RefreshIcon from '~icons/ri/refresh-line'
	import type { PageData } from './$types'
	import type { CreateParentBody, CreateStudentBody, Parent, Student } from '$lib'
	import { dialogProps, Notify, openDialog } from '$lib/store'
	import PlusIcon from '~icons/ic/round-add'
	import { createForm } from 'felte'
	import { CreateStudentSchema as schema } from './validate'
	import { validator } from '@felte/validator-zod'
	import TextField from '$lib/components/TextField.svelte'
	import SelectField from '$lib/components/SelectField.svelte'
	import DateField from '$lib/components/DateField.svelte'
	import { blur } from 'svelte/transition'
	import TrashIcon from '~icons/fa-solid/trash-alt'
	import TimesIcon from '~icons/uil/times'
	import { invalidate } from '$app/navigation'

	export let data: PageData
	let drawerToggleRef: HTMLInputElement
	let scrollClass = ''
	let isNew = true

	let recordData: (Student & Parent) | null = null
	let checked: boolean
	let loading = false
	let abortController: AbortController | undefined = undefined
	$: isNew = !recordData
	$: if (recordData !== null) {
		const { id, ...restOfStudentData } = recordData
		if (id !== undefined) {
			setInitialValues(restOfStudentData)
			reset()
		} else {
			Notify({
				type: 'error',
				id: crypto.randomUUID(),
				description: 'Lỗi không tìm thấy học sinh hoặc gia đình của học sinh này!'
			})
		}
	}

	let activeTabValue: number = 1
	function handleClick(tabValue: number) {
		activeTabValue = tabValue
	}

	var agencyOptions = data?.agencies?.data?.map((el) => ({ label: el.agencyName, value: el.id.toString() }))

	const defaultFormValues = {
		firstName: '',
		lastName: '',
		grade: 'seed',
		enrollDate: '',
		dob: '',
		birthYear: '',
		gender: 'male',
		ethnic: '',
		birthPlace: '',
		tempRes: '',
		permResProvince: '',
		permResDistrict: '',
		permResCommune: '',
		agencyId: parseInt(agencyOptions[0]?.value),
		classRoomId: 1,
		studentId: 1,
		parentName: '',
		parentDob: '',
		parentGender: 'male',
		parentPhoneNumber: '',
		parentZalo: '',
		parentOccupation: '',
		parentLandlord: '',
		parentRoi: '',
		parentBirthPlace: '',
		parentResRegistration: ''
	}

	const studentSchema: {
		name: string
		type: 'text' | 'date' | 'select'
		required: boolean
		options?: { label: string; value: string }[]
	}[] = [
		{
			name: 'firstName',
			type: 'text',
			required: true
		},
		{
			name: 'lastName',
			type: 'text',
			required: true
		},
		{
			name: 'grade',
			type: 'select',
			required: true,
			options: [
				{ label: 'Lớp mầm', value: 'seed' },
				{ label: 'Lớp chồi', value: 'buds' },
				{ label: 'Lớp lá', value: 'leaf' },
				{ label: 'Trẻ ( 18 - 24 tháng tuổi )', value: 'toddlers' }
			]
		},
		{
			name: 'enrollDate',
			type: 'date',
			required: false
		},
		{
			name: 'dob',
			type: 'date',
			required: false
		},
		{
			name: 'gender',
			type: 'select',
			required: true,
			options: [
				{ label: 'Nam', value: 'male' },
				{ label: 'Nữ', value: 'female' }
			]
		},
		{
			name: 'birthYear',
			type: 'text',
			required: false
		},
		{
			name: 'ethnic',
			type: 'text',
			required: false
		},
		{
			name: 'birthPlace',
			type: 'text',
			required: false
		},
		{
			name: 'tempRes',
			type: 'text',
			required: false
		},
		{
			name: 'permResProvince',
			type: 'text',
			required: false
		},
		{
			name: 'permResDistrict',
			type: 'text',
			required: false
		},
		{
			name: 'permResCommune',
			type: 'text',
			required: false
		},
		{
			name: 'agencyId',
			type: 'select',
			required: true,
			options: agencyOptions,
		},
		{
			name: 'classRoomId',
			type: 'select',
			required: true,
			options: [
				{ label: 'Lớp 01', value: '1' },
				{ label: 'Lớp 02', value: '2' },
				{ label: 'Lớp 03', value: '3' }
			]
		}
	]

	const parentSchema: {
		name: string
		type: 'text' | 'date' | 'select'
		required: boolean
		options?: { label: string; value: string }[]
	}[] = [
		{
			name: 'parentName',
			type: 'text',
			required: true
		},
		{
			name: 'parentDob',
			type: 'date',
			required: false
		},
		{
			name: 'parentGender',
			type: 'select',
			required: true,
			options: [
				{ label: 'Nam', value: 'male' },
				{ label: 'Nữ', value: 'female' }
			]
		},
		{
			name: 'parentPhoneNumber',
			type: 'text',
			required: false
		},
		{
			name: 'parentZalo',
			type: 'text',
			required: false
		},
		{
			name: 'parentOccupation',
			type: 'text',
			required: false
		},
		{
			name: 'parentLandlord',
			type: 'text',
			required: false
		},
		{
			name: 'parentRoi',
			type: 'text',
			required: false
		},
		{
			name: 'parentBirthPlace',
			type: 'text',
			required: false
		},
		{
			name: 'parentResRegistration',
			type: 'text',
			required: false
		}
	]

	const { form, errors, setInitialValues, reset } = createForm({
		debounced: { timeout: 500 },
		extend: validator({ schema }),
		transform: (values: any) => ({
			...values,
			agencyId: parseInt(values.agencyId, 10),
			classRoomId: parseInt(values.classRoomId, 10)
		}),
		onSubmit: save
	})

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

	async function save(req: CreateStudentBody & CreateParentBody) {
		loading = true
		const body = JSON.stringify(req)
		const method = isNew ? 'POST' : 'PUT'
		const url = isNew ? '/api/students' : `/api/students/${recordData?.id}`
		const request = fetch(url, {
			method,
			body
		}).then((res) => res.json())

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

	const tabData = [
		{ name: 'Time Sheet', section: 'timeSheet', value: 0 },
		{ name: 'Student List', section: 'studentlist', value: 1 }
	]

	function refreshData() {
		invalidate('app:students')
		invalidate('app:agencies')
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

	async function loadData(id: number, signal: AbortSignal) {
		loading = true
		try {
			const res = await fetch(`/api/students/${id}`, {
				signal
			}).then((res) => res.json())
			recordData = res
		} catch (e) {
			console.error('LoadData: ', e)
			if (e.name !== undefined && e.name === 'AbortError') {
				return
			}

			Notify({ description: 'Lỗi từ phía server', id: crypto.randomUUID(), type: 'error' })
		} finally {
			loading = false
		}
	}

	function resetDefaultForm() {
		setInitialValues(defaultFormValues)
		reset()
	}

	function hide() {
		checked = false
		drawerToggleRef.checked = false
		resetDefaultForm()
	}

	async function handleDelete() {
		if (recordData === undefined || recordData?.id === undefined) {
			return
		}

		loading = true

		try {
			const res = await fetch(`/api/students`, {
				body: JSON.stringify({ ids: [Number(recordData.id)] }),
				method: 'DELETE'
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
						<li>Danh sách học viên</li>
					</ul>
				</div>
				<button class="btn btn-circle btn-ghost btn-sm active:!rotate-180" on:click={refreshData}>
					<RefreshIcon />
				</button>
			</div>

			{#if activeTabValue === 1}
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
			{/if}
		</header>
		<div class="tabs tab-lg w-full">
			<ul class="tabs tab-lg w-full border-b border-gray-400">
				{#each tabData as item}
					<li class={` tab ${activeTabValue === item.value ? 'tab-bordered tab-active' : ''}`}>
						<span
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

		<div class="mt-4">
			{#if activeTabValue === 0}
				<TimeSheet />
			{:else if activeTabValue === 1}
				<StudentList {data} onClick={(id) => show(id)} />
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
				<h4 class="text-lg font-bold">{isNew ? 'Tạo mới học sinh' : 'Chỉnh sửa'}</h4>
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
					<h1 class="text-xl font-bold">1. Thông tin học sinh</h1>

					{#each studentSchema as { type, name, required, options } (name)}
						{#if type === 'text'}
							<TextField error={$errors[name]} {name} {required} />
						{:else if type === 'select'}
							<SelectField error={$errors[name]} {name} {options} {required} />
						{:else if type === 'date'}
							<DateField error={$errors[name]} {name} {required} />
						{/if}
					{/each}
					<h1 class="text-xl font-bold">2. Thông tin Phụ huynh học sinh</h1>
					{#each parentSchema as { type, name, required, options } (name)}
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
