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
	import type { StudentProps, ParentProps } from '../type'
	import dayjs from 'dayjs'

	export let data: PageData

	const API_URL = 'http://127.0.0.1:5000/api/v1'

	let drawerToggleRef: HTMLInputElement
	let scrollClass = ''
	let isNew = true

	let recordData: (StudentProps & ParentProps) | null = null
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

	var agencyOptions = data?.agencies?.data?.map((el) => ({
		label: el.name,
		value: el.id ? el.id.toString() : '1'
	}))

	var classOptions = data?.classes?.data?.map((cl) => ({
		label: cl.name,
		value: cl.id ? cl.id.toString() : '1'
	}))

	var genderMap = (value: string) => {
		if (value === '1') {
			return true
		} else {
			return false
		}
	}

	function genderBooleantoString(gender: boolean) {
		return gender ? '1' : '2'
	}

	const defaultFormValues = {
		firstName: '',
		lastName: '',
		enrolledAt: '', //enrolledAt
		dob: '',
		gender: '1', //boolean
		ethnic: '',
		birthPlace: '',
		tempAdress: '',
		permanentAddressProvince: '',
		permanentAddressDistrict: '',
		permanentAddressCommune: '',
		agencyId: parseInt(agencyOptions[0]?.value || '1'),
		classId: parseInt(classOptions[0]?.value || '1'),
		studentId: 1,
		parentName: '',
		parentDob: '',
		parentGender: genderMap('1'),
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
		options?: { label: string; value: string | boolean }[]
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
			name: 'enrolledAt', //enrolledAt
			type: 'date',
			required: false
		},
		{
			name: 'dob',
			type: 'date',
			required: false
		},
		{
			name: 'gender', //boolean
			type: 'select',
			required: true,
			options: [
				{ label: 'Nam', value: '1' },
				{ label: 'Nữ', value: '2' }
			]
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
			name: 'tempAdress',
			type: 'text',
			required: false
		},
		{
			name: 'permanentAddressProvince',
			type: 'text',
			required: false
		},
		{
			name: 'permanentAddressDistrict',
			type: 'text',
			required: false
		},
		{
			name: 'permanentAddressCommune',
			type: 'text',
			required: false
		},
		{
			name: 'agencyId',
			type: 'select',
			required: true,
			options: agencyOptions
		},
		{
			name: 'classId',
			type: 'select',
			required: true,
			options: classOptions
		}
	]

	const parentSchema: {
		name: string
		type: 'text' | 'date' | 'select'
		required: boolean
		options?: { label: string; value: string | boolean }[]
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
				{ label: 'Nam', value: '1' },
				{ label: 'Nữ', value: '2' }
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
			gender: genderMap(values.gender),
			parentGender: genderMap(values.parentGender),
			agencyId: parseInt(values.agencyId, 10),
			classId: parseInt(values.classRoomId, 10)
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

	async function save(req: ParentProps & StudentProps) {
		loading = true
		const body = JSON.stringify(req)
		console.log('body student: ', body)
		const method = isNew ? 'POST' : 'PUT'
		const url = isNew ? `${API_URL}/students` : `${API_URL}/students/${recordData?.id}`
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
			const res = await fetch(`${API_URL}/students/${id}`, {
				signal
			})
			const studentData = await res.json()
			
			if (studentData.id) {
				recordData = {
					...studentData,
					gender: genderBooleantoString(studentData.gender),
					enrolledAt: dayjs(studentData.enrolledAt).format('YYYY-MM-DD'),
					dob: dayjs(studentData.dob).format('YYYY-MM-DD'),
					parentBirthPlace: studentData.parent_birth_place,
					parentDob: dayjs(studentData.parent_dob).format('YYYY-MM-DD'),
					parentGender: genderBooleantoString(studentData.parent_gender),
					parentLandlord: studentData.landlord,
					parentName: studentData.parent_name,
					parentOccupation: studentData.occupation,
					parentPhoneNumber: studentData.phone_number,
					parentResRegistration: studentData.res_registration,
					parentRoi: studentData.roi,
					parentZalo: studentData.zalo,
					studentId: studentData.student_id,
				}
			} else {
				throw new Error('Student ID not found can not find Parent of Student')
			}
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
