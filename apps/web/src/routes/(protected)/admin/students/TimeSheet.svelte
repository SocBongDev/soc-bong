<script lang="ts">
	import TrackerIcon from '~icons/fluent-mdl2/trackers'
	import ExportIcon from '~icons/mdi/export'
	import ClickDropdown from '$lib/components/ClickDropdown.svelte'
	import dayjs from 'dayjs'
	import { onDestroy, onMount } from 'svelte'
	import { fade } from 'svelte/transition'
	import { Notify, classIdStore, dialogProps, openDialog } from '$lib/store'
	import { statusChange } from '$lib/store'
	import type { PageData } from './$types'
	import type { AttendedStatus } from '$lib/store'
	import { PUBLIC_API_SERVER_URL } from '$env/static/public'
	import { get } from 'svelte/store'
	import type { StudentProps } from '$lib/common/type'
	let inputValue: string = dayjs().format('YYYY-MM') || '07-2024'
	let yearPicked: number = parseInt(inputValue.split('-')[0], 10)
	let monthPicked: number = parseInt(inputValue.split('-')[1], 10)
	let statusArray: AttendedStatus[] = []
	import utc from 'dayjs/plugin/utc'

	dayjs.extend(utc)
	const token = localStorage.getItem('access_token')

	const unsubscribe = statusChange.subscribe((value) => {
		statusArray = value
	})

	onDestroy(() => {
		unsubscribe()
	})

	export let data: PageData
	let classId = get(classIdStore) || 1
	let studentList: StudentProps[] = data.students.data ?? []
	let attendances: any = data.attendances ?? []
	const studentIds =
		studentList && studentList.map((student: any) => student.id).sort((a: any, b: any) => b - a)
	attendances = studentIds && studentIds?.map((id: any) => attendances[id])
	let loading = false
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

	async function refreshData() {
		loading = true
		let datePicked = dayjs(inputValue).format('MM-YYYY')
		loadAttendancesData(classId, datePicked)
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
			studentList = studentData.data
		} catch (err) {
			console.error('Error load student data: ', err)
		} finally {
			loading = false
		}
	}

	async function loadAttendancesData(classId: number, datePicked: string) {
		loading = true
		try {
			const response = await fetch(
				`${PUBLIC_API_SERVER_URL}/attendances?classId=${classId}&period=${datePicked}`,
				{
					method: 'GET',
					headers: {
						Authorization: `Bearer ${token}`,
						'Content-Type': 'application/json'
					}
				}
			)
			const preAttendances = await response.json()

			const studentIds = studentList
				.map((student: any) => student.id)
				.sort((a: any, b: any) => b - a)
			attendances = studentIds.map((id: any) => preAttendances[id])
		} catch (err) {
			console.error('Error load student attendances data: ', err)
		} finally {
			loading = false
		}
	}

	async function handleInput(event: any) {
		const value = (event.target as HTMLInputElement).value
		inputValue = value
		yearPicked = parseInt(value.split('-')[0], 10)
		monthPicked = parseInt(value.split('-')[1], 10)
		let datePicked = dayjs(value).format('MM-YYYY')
		loadAttendancesData(classId, datePicked)
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

	async function handleSelectClassId(event: any) {
		classId = parseInt((event.target as HTMLSelectElement).value)
		$classIdStore = classId
		const datePicked = dayjs(inputValue).format('MM-YYYY')
		loadStudentData(classId)
		loadAttendancesData(classId, datePicked)
	}

	function resetStatusArray() {
		statusChange.set([])
		statusArray = []
	}

	function clearStatusChanges() {
		resetStatusArray()
		isReset = true
		setTimeout(() => {
			isReset = false
		}, 500)
	}

	async function batchUpdate() {
		if (statusArray.length > 0) {
			const updatePromises = statusArray.map((status) => {
				if (status?.id) {
					return fetch(`${PUBLIC_API_SERVER_URL}/attendances`, {
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
					return fetch(`${PUBLIC_API_SERVER_URL}/attendances`, {
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
								studentId: status.studentId && parseInt(status.studentId)
							}
						])
					})
				}
			})

			try {
				await Promise.all(updatePromises)
				Notify({
					type: 'success',
					id: crypto.randomUUID(),
					description: `Đã cập nhật điểm danh thành công cho ${statusArray.length} ngày`
				})
			} catch (error) {
				Notify({
					type: 'error',
					id: crypto.randomUUID(),
					description: 'Lỗi không thể thực hiện chức năng này'
				})
			}
			resetStatusArray()
			await refreshData()
		} else {
			resetStatusArray()
			refreshData()
			Notify({
				type: 'error',
				id: crypto.randomUUID(),
				description: 'Lỗi không thể thực hiện chức năng này'
			})
		}
	}

	async function handleExportAttendances() {
		const response = await fetch(
			`${PUBLIC_API_SERVER_URL}/attendances/${classId}/export-excel?period=${`${
				monthPicked < 10 ? `0${monthPicked}` : monthPicked
			}-${yearPicked}`}`,
			{
				method: 'GET',
				headers: {
					Authorization: `Bearer ${token}`,
					Accept: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
				}
			}
		)
		if (response.ok) {
			const blob = await response.blob()
			// Handle the blob, e.g., download the file or process it further
			const url = window.URL.createObjectURL(blob)
			const a = document.createElement('a')
			a.style.display = 'none'
			a.href = url
			a.download = `Workbook-${classId}-.xlsx` // Specify the file name you want to save as
			document.body.appendChild(a)
			a.click()
			window.URL.revokeObjectURL(url)
		} else {
			Notify({
				type: 'error',
				id: crypto.randomUUID(),
				description: 'Đã có lỗi xảy ra khi thực hiện tải xuống!'
			})
			console.error('Failed to fetch the .xlsx file:', response.statusText)
		}
	}

	async function handleTrackAttendancesDay() {
		let date = dayjs('2024-10-02').utc().local().format('YYYY-MM-DD')
		if (!statusArray.length) {
			//Case 1: Create new att for all students when no elements in statusChange
			const newAttendances: AttendedStatus[] = studentList
				.map((student, index) => {
					const attendancesIndex = attendances[`${index}`]?.attendances.find((att: any) => {
						const attendedDate = dayjs(String(Object(att).attendedAt))
							.utc()
							.local()
							.format('YYYY-MM-DD')
						return attendedDate === date
					})

					if (!attendancesIndex) {
						return {
							date: date,
							studentId: student?.id?.toString(),
							attendedStatus: 'attended',
							classId: classId.toString()
						}
					} else if (attendancesIndex.attendedStatus !== 'attended') {
						return {
							date,
							studentId: attendancesIndex.studentId,
							attendedStatus: 'attended',
							classId: classId.toString(),
							id: attendancesIndex?.id
						}
					}
					return null
				})
				.filter((att) => att !== null) as AttendedStatus[]

			if (newAttendances.length === 0) {
				Notify({
					type: 'success',
					id: crypto.randomUUID(),
					description: `Các học sinh ở ngày ${dayjs(date).utc().local().format('DD')} tháng ${dayjs(
						date
					)
						.utc()
						.local()
						.format('MM')} đã được điểm danh đầy đủ và hệ thống đã ghi lại!`
				})
			}

			statusChange.update((status) => [...status, ...newAttendances])
		} else {
			statusChange.update((existingStatus) => {
				const updatedStatus: AttendedStatus[] = studentList
					.map((student, index) => {
						const existingRecord = existingStatus.find(
							(record) =>
								record.studentId === student.id &&
								dayjs(record.date).utc().local().format('YYYY-MM-DD') === date
						)

						const attendanceForDay = attendances[`${index}`]?.attendances.find((att: any) => {
							const attendedDate = dayjs(String(Object(att).attendedAt))
								.utc()
								.local()
								.format('YYYY-MM-DD')
							return attendedDate === date
						})
						// If there is already an existing record marked as 'attended', skip it
						if (attendanceForDay && attendanceForDay.attendedStatus === 'attended') {
							return null // Skip this record
						}
						if (existingRecord) {
							// If there is an existing record for this student, update the status for today
							return {
								...existingRecord,
								attendedStatus: attendanceForDay ? 'attended' : existingRecord.attendedStatus,
								id: attendanceForDay?.id || existingRecord.id // Retain the same ID or use the found one
							}
						} else if (attendanceForDay) {
							// If no existing record, but attendance is found for today, create a new one
							return {
								date,
								studentId: student.id?.toString(),
								attendedStatus: 'attended',
								classId: classId.toString(),
								id: attendanceForDay?.id
							}
						} else {
							// If no existing record and no attendance found for today, create a default record
							return {
								date,
								studentId: student.id?.toString(),
								attendedStatus: 'attended', // Or set to a different default if needed
								classId: classId.toString()
							}
						}
					})
					.filter((att) => att !== null) as AttendedStatus[]

				const mergedStatus = [
					...existingStatus.filter((record) => {
						return !studentList.some((student) => {
							return (
								student.id == record.studentId &&
								dayjs(record.date).utc().local().format('YYYY-MM-DD') == date
							)
						})
					}),
					...updatedStatus
				]

				return mergedStatus
			})
		}
	}

	onMount(() => {
		loading = true
		let datePicked = dayjs(inputValue).format('MM-YYYY')
		if ($classIdStore) {
			classId = $classIdStore
			loadStudentData(classId)
			loadAttendancesData(classId, datePicked)
		}
	})

	export function refreshStudentAttendances() {
		loading = true
		let datePicked = dayjs(inputValue).format('MM-YYYY')
		if (classId) {
			loadStudentData(classId)
			loadAttendancesData(classId, datePicked)
		}
	}
</script>

<div class="relative flex h-full w-full flex-col justify-start gap-4">
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
			<div class="group-button flex w-fit items-center gap-2">
				<button
					class="flex w-fit items-center justify-center gap-1.5 rounded border border-gray-400 bg-gray-300 px-1.5 py-0.5 text-center text-sm font-semibold"
					on:click={handleExportAttendances}
				>
					<ExportIcon class="h-4 w-4" />
					Export
				</button>

				<button
					class="flex w-fit items-center justify-center gap-1.5 rounded border border-gray-400 bg-gray-300 px-1.5 py-0.5 text-center text-sm font-semibold"
					on:click={handleTrackAttendancesDay}
				>
					<TrackerIcon class="h-4 w-4" />
					Track Attended A Day
				</button>
			</div>
		</div>
	</div>
	{#if loading}
		<div class="my-4 flex h-full w-full items-center justify-center">
			<span class="loading loading-infinity loading-lg" />
		</div>
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
								{#if attendances[`${id}`]?.attendances.some((attendance) => {
									const attendedDate = dayjs(String(Object(attendance).attendedAt))
									return attendedDate.date() === date.day && attendedDate.month() + 1 === monthPicked
								})}
									<td class="w-fit p-0.5">
										{#each attendances[`${id}`]?.attendances as attendance}
											{#if dayjs(String(Object(attendance).attendedAt)).date() === date.day && dayjs(String(Object(attendance).attendedAt)).month() + 1 === monthPicked}
												{#key isReset}
													<ClickDropdown
														data={attendance}
														date={date.day.toString()}
														studentId={student?.id?.toString()}
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
												studentId={student?.id?.toString()}
												classId={classId.toString()}
												{monthPicked}
												{yearPicked}
											/>
										{/key}
									</td>
								{/if}
							{/each}
							<td class="min-w-[4rem] px-1"
								>{attendances[`${id}`]?.attendances?.filter(
									(att) => att.attendedStatus === 'attended'
								).length || `0`} buổi</td
							>
						</tr>
					{/each}
				{/if}
			</tbody>
		</table>

		<div class="join mt-auto self-center">
			<a
				class={data.students.page === 1 ? 'pointer-events-none cursor-default opacity-40' : ''}
				href={`/admin/students?page=${data.students.page - 1}&pageSize=${data.students.pageSize}`}
			>
				<button class="btn join-item">«</button>
			</a>
			<button class="btn join-item">Trang {data.students.page}</button>
			<a
				class={data.students.data.length < (data.students.pageSize || 15) ||
				data.students.data.length === 0
					? 'pointer-events-none cursor-default opacity-40'
					: ''}
				href={`/admin/students?page=${data.students.page + 1}&pageSize=${data.students.pageSize}`}
			>
				<button class="btn join-item">»</button>
			</a>
		</div>

		{#if statusArray.length > 0}
			<div class="absolute bottom-6 left-1/2 w-1/2 -translate-x-1/2" transition:fade>
				<div class="alert flex justify-between rounded-full bg-white py-2.5 text-sm shadow">
					<div class="flex w-1/2 items-center gap-3">
						<span>Đã điểm danh <strong>{statusArray.length}</strong> ngày</span>
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
	{/if}
</div>
