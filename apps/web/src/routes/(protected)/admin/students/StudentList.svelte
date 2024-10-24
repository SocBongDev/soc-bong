<script lang="ts">
	import EllipsisIcon from '~icons/fa6-solid/ellipsis'
	import { fade } from 'svelte/transition'
	import ArrowRightIcon from '~icons/formkit/arrowright'
	import { dialogProps, Notify, openDialog } from '$lib/store'
	import dayjs from 'dayjs'
	import type { PageData } from './$types'
	import { invalidate } from '$app/navigation'
	import { PUBLIC_API_SERVER_URL } from '$env/static/public'
	import { classIdStore } from '$lib/store'
	import { onMount } from 'svelte'
	import { get } from 'svelte/store'

	export let data: PageData
	let isChecked: string[] = []
	let classId = get(classIdStore) || 1
	let isCheckedAll = false
	let loading = false
	let studentList = {
		data: data.students.data,
		page: data.students.page,
		pageSize: data.students.pageSize ?? '15'
	}
	export let onClick: (id: number) => void
	const token = localStorage.getItem('access_token')

	const studentClassMap = {
		seed: 'Lớp mầm',
		buds: 'Lớp chồi',
		leaf: 'Lớp lá',
		toddlers: 'Trẻ ( 18 - 24 tháng tuổi )'
	}

	function formatStudentDate(studentDate: string | null) {
		if (!studentDate) return 'Chưa điền'
		return dayjs(studentDate).format('DD/MM/YYYY')
	}

	function formatBirthYear(studentDate: string | null) {
		if (!studentDate) return 'Chưa điền'
		return dayjs(studentDate).format('YYYY')
	}

	function formatStudentGender(studentGender: string | boolean | null) {
		switch (studentGender) {
			case true:
				return 'Nam'
			case false:
				return 'Nữ'
			case null:
				return 'Chưa điền giới tính'
			default:
				return 'khác'
		}
	}

	function formatStudentClassId(classId: number) {
		const classRoomId = data?.classes?.data.find((cl) => cl.id == classId)
		switch (classRoomId?.grade) {
			case 'buds':
			case 'seed':
			case 'leaf':
			case 'toddlers':
				return studentClassMap[classRoomId?.grade as keyof typeof studentClassMap]
			default:
				return 'Lớp chưa đúng!'
		}
	}

	function formatAgencyName(agencyId: number) {
		const agency = data.agencies.data.find(
			(el) => (el.id && parseInt(el.id?.toString())) === agencyId
		)
		return agency?.name || 'N/A'
	}

	function handleCheckAll() {
		isCheckedAll = !isCheckedAll
		if (!isCheckedAll) {
			isChecked = []
			return
		}
		isChecked = studentList?.data?.map((student: any) => student?.id.toString())
	}

	function handleCheck(e: any) {
		const { id, checked } = e.currentTarget

		if (!checked) {
			const isValidCheckAll = isChecked.length === studentList.data.length
			if (isValidCheckAll) {
				isCheckedAll = false
			}

			isChecked = isChecked.filter((item) => item !== id)
			return
		}

		isChecked = [...isChecked, id]
		const isValidCheckAll = isChecked.length === studentList?.data.length
		if (isValidCheckAll) {
			isCheckedAll = true
		}
	}

	function clearSelected() {
		isCheckedAll = false
		isChecked = []
	}

	function refreshData() {
		loading = true
		invalidate('app:students')
	}

	async function batchDelete() {
		loading = true
		try {
			const ids = isChecked.map((el) => Number(el))
			const res = await fetch(`${PUBLIC_API_SERVER_URL}/students`, {
				method: 'DELETE',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ ids })
			}).then((res) => {
				if (res.status === 403) {
					Notify({
						type: 'error',
						id: crypto.randomUUID(),
						description: 'Bạn không đủ quyền hạn làm việc này!'
					})
				}
			})
			refreshStudentList()
			clearSelected()
		} catch (e) {
			console.error('Batch Delete error', e)
			Notify({ type: 'error', id: crypto.randomUUID(), description: 'Lỗi từ phía server' })
		} finally {
			loading = false
		}
	}

	async function loadStudentData(classId: number) {
		loading = true
		try {
			const response = await fetch(`${PUBLIC_API_SERVER_URL}/students?classId=${classId}`, {
				method: 'GET',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json'
				}
			})
			const studentData = await response.json()
			studentList = { ...studentList, data: studentData.data }
		} catch (err) {
			console.error('Error load student data: ', err)
		} finally {
			loading = false
		}
	}

	async function handleSelectClassId(event: any) {
		loading = true
		classId = parseInt((event.target as HTMLSelectElement).value)
		$classIdStore = classId
		await loadStudentData(classId)
	}

	onMount(() => {
		loading = true
		if ($classIdStore) {
			classId = $classIdStore
			loadStudentData(classId)
		}
	})

	export function refreshStudentList() {
		loading = true
		if (classId) {
			loadStudentData(classId)
		}
	}
</script>

<div class="relative flex h-auto flex-col gap-10 overflow-x-auto">
	<div class="flex items-center justify-end gap-4 pl-4">
		<h3 class="">Chọn Lớp:</h3>
		{#if data.classes?.data.length > 0}
			<select
				on:change={handleSelectClassId}
				id="classId"
				class="select select-ghost h-fit min-h-0 w-fit max-w-xs font-bold"
			>
				{#each data.classes?.data as classroom, index}
					<option value={`${classroom?.id}`} selected={classroom?.id === classId}
						>{classroom?.name}</option
					>
				{/each}
			</select>
		{:else}
			<select class="disabled select select-ghost h-fit min-h-0 w-fit max-w-xs font-bold">
				<option value="1">Không tồn tại lớp nào</option>
			</select>
		{/if}
	</div>

	{#if loading}
		<div class="my-4 flex h-full w-full items-center justify-center">
			<span class="loading loading-infinity loading-lg" />
		</div>
	{:else}
		<table class="table -mt-8">
			<thead>
				<tr class="text-center">
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
					<th>Khối lớp</th>
					<th>Tên học sinh</th>
					<th>Họ học sinh</th>
					<th>Ngày nhập học</th>
					<th>Ngày sinh nhật</th>
					<th>Năm sinh</th>
					<th>Giới tính</th>
					<th>Cơ Sở Trường Học</th>
					<th>
						<button class="btn btn-square btn-ghost btn-sm active:!translate-y-1">
							<EllipsisIcon />
						</button>
					</th>
				</tr>
			</thead>
			<tbody>
				{#if studentList.data.length > 0}
					{#each studentList?.data as student (student.id)}
						{#if student.id !== undefined}
							<tr class="hover cursor-pointer text-center">
								<th>
									<label>
										<input
											id={student.id.toString()}
											type="checkbox"
											class="checkbox checkbox-sm rounded"
											checked={isChecked.includes(student.id?.toString())}
											on:click={handleCheck}
										/>
									</label>
								</th>
								<td on:click={() => onClick(Number(student.id))}
									>{formatStudentClassId(student.classId)}</td
								>
								<th on:click={() => onClick(Number(student.id))}>{student.firstName}</th>
								<th on:click={() => onClick(Number(student.id))}>{student.lastName}</th>
								<td on:click={() => onClick(Number(student.id))}
									>{formatStudentDate(student.enrolledAt) || ''}</td
								>
								<td on:click={() => onClick(Number(student.id))}
									>{formatStudentDate(student.dob) || ''}</td
								>
								<td on:click={() => onClick(Number(student.id))}>{formatBirthYear(student.dob)}</td>
								<td on:click={() => onClick(Number(student.id))}
									>{formatStudentGender(student.gender)}</td
								>
								<td on:click={() => onClick(Number(student.id))}
									>{formatAgencyName(student.agencyId)}</td
								>
								<td on:click={() => onClick(Number(student.id))}>
									<div class="px-2 text-center align-middle">
										<ArrowRightIcon />
									</div>
								</td>
							</tr>
						{/if}
					{/each}
				{:else}
					<tr class="flex h-12 w-full flex-row items-center justify-center border-none">
						<p class="w-full text-base font-medium">Không có dữ liệu...</p>
					</tr>
				{/if}
			</tbody>
		</table>
	{/if}

	<!-- Page Num -->
	<div class="join mt-auto self-center">
		<a
			class={studentList.page === 1 ? 'pointer-events-none cursor-default opacity-40' : ''}
			href={`/admin/students/?page=${studentList.page - 1}&pageSize=${studentList.pageSize}`}
		>
			<button class="btn join-item">«</button>
		</a>
		<button class="btn join-item">Trang {studentList.page}</button>
		<a
			class={studentList.data.length < studentList.pageSize || studentList.data.length === 0
				? 'pointer-events-none cursor-default opacity-40'
				: ''}
			href={`/admin/students?page=${studentList.page + 1}&pageSize=${studentList.pageSize}`}
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
