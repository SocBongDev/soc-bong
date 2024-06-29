<script lang="ts">
	import EllipsisIcon from '~icons/fa6-solid/ellipsis'
	import { fade } from 'svelte/transition'
	import ArrowRightIcon from '~icons/formkit/arrowright'
	import { dialogProps, Notify, openDialog } from '$lib/store'
	import dayjs from 'dayjs'
	import type { PageData } from './$types'
	import { invalidate } from '$app/navigation'
	import { PUBLIC_API_SERVER_URL } from '$env/static/public'

	export let data: PageData
	let isChecked: string[] = []
	let classId = 1
	let isCheckedAll = false
	let studentList = {
		data: data.students.data,
		page: data.students.page,
		pageSize: data.students.pageSize ?? '15'
	}
	export let onClick: (id: number) => void
	const token = localStorage.getItem('access_token')

	async function getClassId() {
		const res = await fetch(`${PUBLIC_API_SERVER_URL}/classes`, {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${token}`,
				'Content-Type': 'application/json'
			}
		})
		const data = await res.json()
		return data
	}

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

	function formatStudentGender(studentGender: string | null) {
		switch (studentGender) {
			case 'male':
			case 'female':
				return studentGenderMap[studentGender as keyof typeof studentGenderMap]
			case null:
				return 'Chưa điền giới tính'
			default:
				return 'khác'
		}
	}

	function formatStudentClassId(studentClass: number) {
		switch (studentClass) {
			case 1:
			case 2:
			case 3:
				return studentClassIdMap[studentClass as keyof typeof studentClassIdMap]
			default:
				return 'Lớp chưa đúng'
		}
	}

	function formatStudentClass(studentClass: string) {
		switch (studentClass) {
			case 'seed':
			case 'buds':
			case 'leaf':
			case 'toddlers':
				return studentClassMap[studentClass as keyof typeof studentClassMap]
			default:
				return 'N/A'
		}
	}

	function formatAgencyName(agencyId: number) {
		const agency = data.agencies.data.find((el) => parseInt(el.id.toString()) === agencyId)
		return agency?.agencyName || 'N/A'
	}

	function handleCheckAll() {
		isCheckedAll = !isCheckedAll
		if (!isCheckedAll) {
			isChecked = []
			return
		} else {
			isChecked = data.students?.data?.map((student: any) => student?.id)
		}
	}

	function handleCheck(e: any) {
		const { id, checked } = e.currentTarget

		if (!checked) {
			const isValidCheckAll = isChecked.length === data.students.data.length
			if (isValidCheckAll) {
				isCheckedAll = false
			}

			isChecked = isChecked.filter((item) => item !== id)
			return
		}

		isChecked = [...isChecked, id]
		const isValidCheckAll = isChecked.length === data.students.data.length
		if (isValidCheckAll) {
			isCheckedAll = true
		}
	}

	function clearSelected() {
		isCheckedAll = false
		isChecked = []
	}

	function refreshData() {
		invalidate('app:students')
	}

	async function batchDelete() {
		try {
			const ids = isChecked.map((el) => Number(el))
			await fetch('/api/students', { body: JSON.stringify({ ids }), method: 'DELETE' })
			refreshData()
			clearSelected()
		} catch (e) {
			console.error('Batch Delete error', e)
			Notify({ type: 'error', id: crypto.randomUUID(), description: 'Lỗi từ phía server' })
		}
	}
</script>

<div class="relative flex h-auto flex-col gap-10 overflow-x-auto">
	<table class="table">
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
				<th>Mã phòng học</th>
				<th>
					<button class="btn btn-square btn-ghost btn-sm active:!translate-y-1">
						<EllipsisIcon />
					</button>
				</th>
			</tr>
		</thead>
		<tbody>
			{#if data}
				{#each data?.students.data as student (student.id)}
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
						<td on:click={() => onClick(student.id)}>{formatStudentClass(student.grade)}</td>
						<th on:click={() => onClick(student.id)}>{student.firstName}</th>
						<th on:click={() => onClick(student.id)}>{student.lastName}</th>
						<td on:click={() => onClick(student.id)}
							>{formatStudentDate(student.enrollDate) || ''}</td
						>
						<td on:click={() => onClick(student.id)}>{formatStudentDate(student.dob) || ''}</td>
						<td on:click={() => onClick(student.id)}>{student.birthYear}</td>
						<td on:click={() => onClick(student.id)}>{formatStudentGender(student.gender)}</td>
						<td on:click={() => onClick(student.id)}
							>{formatAgencyName(student.agencyId)}</td
						>
						<td on:click={() => onClick(student.id)}>{formatStudentClassId(student.classRoomId)}</td
						>
						<td on:click={() => onClick(student.id)}>
							<div class="px-2 text-center align-middle">
								<ArrowRightIcon />
							</div>
						</td>
					</tr>
				{/each}
			{/if}
		</tbody>
	</table>
	<!-- Page Num -->
	<div class="join mt-auto self-center">
		<a
			class={data.students.page === 1 ? 'pointer-events-none cursor-default opacity-40' : ''}
			href={`/admin?page=${data.students.page - 1}&pageSize=${data.students.pageSize}`}
		>
			<button class="btn join-item">«</button>
		</a>
		<button class="btn join-item">Trang {data.students.page}</button>
		<a
			class={data.students.data.length < data.students.pageSize || data.students.data.length === 0
				? 'pointer-events-none cursor-default opacity-40'
				: ''}
			href={`/admin?page=${data.students.page + 1}&pageSize=${data.students.pageSize}`}
		>
			<button class="btn join-item">»</button>
		</a>
	</div>
	<!-- handle check -->
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
