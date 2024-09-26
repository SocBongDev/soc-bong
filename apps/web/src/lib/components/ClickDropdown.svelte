<script lang="ts">
	import { statusChange } from '$lib/store'
	import dayjs from 'dayjs'
	import DefaultIcon from './DefaultIcon.svelte'
	export let data: any

	export let date: string
	export let studentId: string | undefined
	export let classId: string
	export let monthPicked: number
	export let yearPicked: number
	export const trackADay: any[] = []
	let isDropdownOpen: boolean = false

	type attendedStatus = 'attended' | 'absented' | 'excused' | 'dayoff' | 'holiday'
	let active: string | null = ''
	type statusType = {
		name: attendedStatus
		color: string
		letter: string
	}[]

	const status: statusType = [
		{ name: 'absented', color: 'text-red-500', letter: '🔴' },
		{ name: 'attended', color: 'text-green-500', letter: '🟢' },
		{ name: 'excused', color: 'text-gray-500', letter: '🟡' },
		{ name: 'dayoff', color: 'text-black', letter: '⚫' },
		{ name: 'holiday', color: 'text-blue-500', letter: '🔵' }
	]

	$: activeStatus = status[0]

	$: activeTrackStatus = ''

	$: {
		const selectedDate = dayjs(`${yearPicked}-${monthPicked}-${date}`).format('YYYY-MM-DD')
		statusChange.subscribe((values) => {
			const statusFind = values.find(
				(status) =>
					dayjs(status.date).format('YYYY-MM-DD') == selectedDate && status.studentId == studentId
			)
			if (statusFind) {
				activeTrackStatus = statusFind.attendedStatus
			} else {
				activeTrackStatus = ''
			}
		})
	}

	function handleDropdownClick() {
		isDropdownOpen = !isDropdownOpen
	}

	function handleChangeState(index: number) {
		activeStatus = status[index]
		activeTrackStatus = activeStatus.name;
		statusChange.update((status) => {
			let id = data?.id
			if (!id) {
				const existingIndex = status.findIndex(
					(s) =>
						s.date === dayjs(`${yearPicked}-${monthPicked}-${date}`).format('YYYY-MM-DD') &&
						s.studentId === studentId
				)
				if (existingIndex !== -1) {
					return status.map((s, i) => {
						if (i === existingIndex) {
							return { ...s, attendedStatus: activeStatus.name }
						}
						return s
					})
				}
				return [
					...status,
					{
						date: dayjs(`${yearPicked}-${monthPicked}-${date}`).format('YYYY-MM-DD'),
						studentId: studentId,
						attendedStatus: activeStatus.name,
						classId: classId
					}
				]
			} else {
				const existingIndex = status.findIndex((s) => s?.id === id)
				if (existingIndex !== -1) {
					return status.map((s, i) => {
						if (i === existingIndex) {
							return { ...s, attendedStatus: activeStatus.name }
						}
						return s
					})
				}
				return [
					...status,
					{
						id: id,
						date: dayjs(
							`${yearPicked}-${monthPicked}-${date} ${dayjs().format('HH:mm:ss')}`
						).format('YYYY-MM-DD HH:mm:ss.SSS[Z]'),
						studentId: studentId,
						attendedStatus: activeStatus.name,
						classId: classId
					}
				]
			}
		})
		isDropdownOpen = false
	}

	function handleDropdownFocusLoss({
		relatedTarget,
		currentTarget
	}: {
		relatedTarget: EventTarget | null
		currentTarget: HTMLElement
	}) {
		if (relatedTarget instanceof HTMLElement && !currentTarget.contains(relatedTarget)) {
			isDropdownOpen = false
		}
	}

	function getAttendanceLetter(
		active: string | null,
		data: any,
		activeTrackStatus?: string | null
	) {
		const attendedStatus = data?.attendedStatus

		// Prioritize the activeTrackStatus, since it represents the latest change
		if (activeTrackStatus) {
			return status.find((s) => s.name === activeTrackStatus)?.letter || 'default'
		}
		if (active !== '') {
			return active
		}
		switch (data?.attendedStatus) {
			case 'attended':
				return status[1].letter
			case 'absented':
				return status[0].letter
			case 'excused':
				return status[2].letter
			case 'dayoff':
				return status[3].letter
			case 'holiday':
				return status[4].letter
			default:
				return 'default'
		}
	}
</script>

<div class="dropdown" on:focusout={handleDropdownFocusLoss}>
	<button
		class="btn btn-square btn-xs align-middle active:bg-slate-500"
		on:click={handleDropdownClick}
	>
		{#if getAttendanceLetter(active, data, activeTrackStatus) == 'default'}
			<DefaultIcon />
		{:else}
			<span class="max-h-full w-fit">
				{getAttendanceLetter(active, data, activeTrackStatus)}
			</span>
		{/if}
	</button>
	<ul
		class="dropdown-content menu bg-base-100 rounded-box z-10 p-2 shadow"
		style:visibility={isDropdownOpen ? 'visible' : 'hidden'}
	>
		{#each status as { name, color, letter }, index (letter)}
			<button
				class="flex items-center"
				on:click={(event) => {
					handleChangeState(index)
					active = event.currentTarget.textContent
				}}
			>
				<span class="btn btn-square text-gray-400">{letter}</span>
			</button>
		{/each}
	</ul>
</div>