<script lang="ts">
	import LegendIcon from '~icons/ic/outline-watch-later'
	import ExportIcon from '~icons/mdi/export'
	import ClickDropdown from '$lib/components/ClickDropdown.svelte'
	import dayjs from 'dayjs'
	import { onDestroy, onMount } from 'svelte'
	import { fade } from 'svelte/transition'
	import { Notify, dialogProps, openDialog } from '$lib/store'
	import { statusChange } from '$lib/store'
	import type { AttendedStatus } from '$lib/store'
	import { PUBLIC_API_SERVER_URL } from '$env/static/public'
	let inputValue: string = dayjs().format('YYYY-MM') || '2023-05'
	let yearPicked: number = parseInt(inputValue.split('-')[0], 10)
	let monthPicked: number = parseInt(inputValue.split('-')[1], 10)
	let statusArray: AttendedStatus[] = []

	const token = localStorage.getItem('access_token')

	const unsubscribe = statusChange.subscribe((value) => {
		statusArray = value
	})

	onDestroy(() => {
		unsubscribe()
	})

	export let data
	let classId = 1
	let attendances: any = {}
	let studentList: any = []
	let loading = true
	$: isReset = false

	const status = [
		{ name: 'Nghỉ không phép', color: 'bg-red-600', letter: '🔴' },
		{ name: 'Có mặt', color: 'bg-green-500', letter: '🟢' },
		{ name: 'Nghỉ có phép', color: 'bg-yellow-400', letter: '🟡' },
		{ name: 'Ngày nghỉ ở trường', color: 'bg-gray-700', letter: '⚫' },
		{ name: 'Ngày nghỉ lễ', color: 'bg-blue-600', letter: '🔵' }
	]

	function generateWeekDays(day: number) {
		switch (day) {
			case 0:
				return 'S'
			case 1:
				return 'M'
			case 2:
				return 'T'
			case 3:
				return 'W'
			case 4:
				return 'T'
			case 5:
				return 'F'
			case 6:
				return 'S'
			default:
				return ''
		}
	}

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

	async function handleInput(event: any) {
		const value = (event.target as HTMLInputElement).value
		inputValue = value
		yearPicked = parseInt(value.split('-')[0], 10)
		monthPicked = parseInt(value.split('-')[1], 10)
		let datePicked = dayjs(value).format('MM-YYYY')
		const res = await fetch(
			`${PUBLIC_API_SERVER_URL}/attendances?classId=${classId}&period=${datePicked}`,
			{
				method: 'GET',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json'
				}
			}
		)
		attendances = await res.json()
	}

	function generateCalendar(year: number, month: number) {
		const daysInMonth = new Date(year, month, 0).getDate()
		let thRow = []
		for (let day = 1; day <= daysInMonth; day++) {
			let weekDay = new Date(year, month - 1, day).getDay()
			thRow[day - 1] = {
				day: day,
				weekDay: generateWeekDays(weekDay)
			}
		}
		return thRow
	}

	async function fetchData() {
		const getStudent = await fetch(`${PUBLIC_API_SERVER_URL}/students?classId=${classId}`, {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${token}`,
				'Content-Type': 'application/json'
			}
		})
		const studentData = await getStudent.json()
		studentList = studentData.data
		const res = await fetch(
			`${PUBLIC_API_SERVER_URL}/attendances?classId=${classId}&period=${dayjs().format('MM-YYYY')}`,
			{
				method: 'GET',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json'
				}
			}
		)
		attendances = await res.json()
		loading = false
	}

	onMount(() => {
		fetchData()
	})

	async function handleSelectClassId(event: any) {
		classId = parseInt((event.target as HTMLSelectElement).value)
		const datePicked = dayjs(inputValue).format('MM-YYYY')

		const studentsList = await fetch(`${PUBLIC_API_SERVER_URL}/students?classId=${classId}`, {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${token}`,
				'Content-Type': 'application/json'
			}
		})
		const attendancesList = await fetch(
			`${PUBLIC_API_SERVER_URL}/attendances?classId=${classId}&period=${datePicked}`,
			{
				method: 'GET',
				headers: {
					Authorization: `Bearer ${token}`,
					'Content-Type': 'application/json'
				}
			}
		)
		const studentData = await studentsList.json()

		studentList = studentData.data
		attendances = await attendancesList.json()
	}

	function clearStatusChanges() {
		statusChange.set([])
		statusArray = []
		isReset = true
		setTimeout(() => {
			isReset = false
		}, 500)
	}

	async function batchUpdate() {
		if (statusArray.length > 0) {
			statusArray.forEach(async (status) => {
				if (status?.id) {
					const res = await fetch(`${PUBLIC_API_SERVER_URL}/attendances`, {
						method: 'PATCH',
						headers: {
							Authorization: `Bearer ${token}`,
							'Content-Type': 'application/json'
						},
						body: JSON.stringify([
							{
								id: status.id,
								attendedStatus: status.attendedStatus
							}
						])
					})
				} else {
					const res = await fetch(`${PUBLIC_API_SERVER_URL}/attendances`, {
						method: 'POST',
						headers: {
							Authorization: `Bearer ${token}`,
							'Content-Type': 'application/json'
						},
						body: JSON.stringify([
							{
								attendedAt: status.date,
								attendedStatus: status.attendedStatus,
								classId: status.classId && parseInt(status.classId),
								studentId: status.studentId
							}
						])
					})
				}
			})
			Notify({
				type: 'success',
				id: crypto.randomUUID(),
				description: `Đã cập nhật điểm danh thành công cho ${statusArray.length} ngày`
			})
		} else {
			Notify({
				type: 'error',
				id: crypto.randomUUID(),
				description: 'Lỗi không thể thực hiện chức năng này'
			})
		}
		statusChange.set([])
		statusArray = []
		fetchData()
	}
</script>

<div class="h-full w-full flex flex-col justify-start gap-4">
	<div class="row-span-full mt-1 grid grid-cols-5 px-2">
		{#each status as { name, color, letter } (name)}
			<div class="flex items-center gap-1">
				<div class={`h-5 w-5 ${color} rounded-full border border-black/50`} />

				<p>- {name}</p>
			</div>
		{/each}
	</div>
	<div class="mt-2 flex justify-between">
		<div class="flex items-center justify-start">
			<select
				on:change={handleSelectClassId}
				id="classId"
				class="select select-ghost h-fit min-h-0 w-fit max-w-xs pl-2 font-bold"
			>
				{#await getClassId()}
					Loading Classroom...
				{:then classes}
					{#if classes.data.length > 0}
						{#each classes?.data as classroom, index}
							<option value={`${classroom?.id}`}>{classroom?.name}</option>
						{/each}
					{/if}
				{:catch error}
					System error: {error.message}
				{/await}
			</select>
			<div class="dropdown-calendar text-sm">
				<input
					type="month"
					id="month"
					name="month"
					bind:value={inputValue}
					on:input={handleInput}
					class="input input-ghost h-fit min-h-0 w-fit max-w-xs font-bold"
				/>
			</div>
		</div>

		<div class="flex items-center justify-end gap-2">
			<select class="select select-ghost h-fit min-h-0 w-fit max-w-xs pl-2 font-bold">
				<option value="1" disabled>All tracked hours</option>
				<option value="2">All tracked hours</option>
				<option value="3">All tracked hours</option>
				<option value="4">All tracked hours 2</option>
			</select>

			<select class="select select-ghost h-fit min-h-0 w-fit max-w-xs pl-2 font-bold">
				<option value="1" disabled>All Group</option>
				<option value="2">All Group</option>
				<option value="3">All Group</option>
				<option value="4">All Group</option>
			</select>

			<select class="select select-ghost h-fit min-h-0 w-fit max-w-xs pl-2 font-bold">
				<option value="1" disabled>All Schedules</option>
				<option value="2">All Schedules</option>
				<option value="3">All Schedules</option>
				<option value="4">All Schedules</option>
			</select>

			<div class="group-button flex w-fit items-center gap-2">
				<button
					class="flex w-fit items-center justify-center gap-1.5 rounded border border-gray-400 bg-gray-300 px-1.5 py-0.5 text-center text-sm font-bold"
				>
					<LegendIcon class="h-4 w-4" />
					Legend
				</button>
				<button
					class="flex w-fit items-center justify-center gap-1.5 rounded border border-gray-400 bg-gray-300 px-1.5 py-0.5 text-center text-sm font-bold"
				>
					<ExportIcon class="h-4 w-4" />
					Export
				</button>
			</div>
		</div>
	</div>
	{#if loading}
		<p>Loading...</p>
	{:else}
		<table class="table">
			<thead class="text-center">
				<tr>
					<th class="px-0">
						<label>
							<input
								type="text"
								placeholder="🔎 Tìm kiếm"
								class="input input-ghost w-fit max-w-xs font-bold"
							/>
						</label>
					</th>
					{#each generateCalendar(yearPicked, monthPicked) as { day, weekDay } (day)}
						<th class="group p-1">
							{weekDay}
							<br />
							{day}
						</th>
					{/each}
					<th class="min-w-[4rem] px-1">Tổng Kết</th>
				</tr>
			</thead>
			<tbody>
				<!-- thêm studentList -->
				<!-- {#each Object.entries(data) as student, index (index)} -->
				{#if studentList.length === 0}
					<tr class="hover cursor-pointer text-center">
						<td class="max-w-xs px-0">Không có dữ liệu</td>
					</tr>
				{:else}
					{#each studentList as student, id (id)}
						<tr class="hover cursor-pointer text-center">
							<th class="max-w-xs px-0">
								<div class="flex items-center justify-start gap-2">
									<div class="avatar">
										<div class="w-12 rounded-full">
											<img
												alt="avatar"
												src="https://www.shutterstock.com/image-vector/young-smiling-man-avatar-brown-600w-2261401207.jpg"
											/>
										</div>
									</div>
									{#if student}
										<p>
											{student.firstName}
											{student.lastName}
										</p>
									{/if}
								</div>
							</th>

							{#each generateCalendar(yearPicked, monthPicked) as date, index (index)}
								{#if attendances[`${id + 1}`]?.some((attendance) => {
									const attendedDate = dayjs(String(Object(attendance).attendedAt))
									return attendedDate.date() === date.day && attendedDate.month() + 1 === monthPicked
								})}
									<td class="w-fit p-0.5">
										{#each attendances[`${id + 1}`] as attendance}
											{#if dayjs(String(Object(attendance).attendedAt)).date() === date.day && dayjs(String(Object(attendance).attendedAt)).month() + 1 === monthPicked}
												{#key isReset}
													<ClickDropdown
														data={attendance}
														date={date.day.toString()}
														studentId={student?.id}
														classId={classId.toString()}
														{monthPicked}
														{yearPicked}
													/>
												{/key}
											{/if}
										{/each}
									</td>
								{:else}
									<td class="w-fit p-0.5">
										{#key isReset}
											<ClickDropdown
												data={null}
												date={date.day.toString()}
												studentId={student?.id}
												classId={classId.toString()}
												{monthPicked}
												{yearPicked}
											/>
										{/key}
									</td>
								{/if}
							{/each}
							<td class="min-w-[4rem] px-1">30 buổi</td>
						</tr>
					{/each}
				{/if}
			</tbody>
		</table>

		{#if statusArray.length > 0}
			<div class="absolute bottom-10 left-1/2 w-1/2 -translate-x-1/2" transition:fade>
				<div class="alert flex justify-between rounded-full bg-white py-2.5 text-sm shadow">
					<div class="flex w-1/2 items-center gap-3">
						<span>Đã sửa <strong>{statusArray.length}</strong> ngày</span>
						<button
							class="btn btn-outline btn-sm rounded border-2 bg-white normal-case"
							on:click={clearStatusChanges}>Huỷ thay đổi</button
						>
					</div>
					<button
						class="btn btn-ghost btn-sm rounded normal-case text-red-500 hover:bg-red-100"
						on:click={() => {
							dialogProps.set({
								description: 'Tiến hành điểm danh các ngày này?',
								title: 'Yêu cầu xác nhận!',
								onContinue: batchUpdate
							})
							openDialog.set(true)
						}}>Xác nhận điểm danh</button
					>
				</div>
			</div>
		{/if}
		<!-- <div class="join mt-auto self-center">
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
	</div> -->
	{/if}
</div>
