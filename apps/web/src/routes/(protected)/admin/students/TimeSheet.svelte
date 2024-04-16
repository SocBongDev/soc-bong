<script lang="ts">
	import LegendIcon from '~icons/ic/outline-watch-later'
	import ExportIcon from '~icons/mdi/export'
	import ClickDropdown from '$lib/components/ClickDropdown.svelte'

	let inputValue: string = '2023-05'
	let yearPicked: number = parseInt(inputValue.split('-')[0], 10)
	let monthPicked: number = parseInt(inputValue.split('-')[1], 10)

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

	function handleInput(event: any) {
		const value = (event.target as HTMLInputElement).value
		inputValue = value
		yearPicked = parseInt(value.split('-')[0], 10)
		monthPicked = parseInt(value.split('-')[1], 10)
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
</script>

<div class="h-full w-full">
	<div class="mt-2 flex justify-between">
		<div class="flex items-center justify-start">
			<select class="select select-ghost h-fit min-h-0 w-fit max-w-xs pl-2 font-bold">
				<option value="1" disabled>Monthly Timesheets</option>
				<option value="2">Third-quarter Timessheets</option>
				<option value="3">Half-year Timesheets</option>
				<option value="4">Yearly Timesheets</option>
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
			{#each Array(5) as _, index (index)}
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
							<p>Cho Chang</p>
						</div>
					</th>
					{#each generateCalendar(yearPicked, monthPicked) as _, index (index)}
						<td class="w-fit p-0.5">
							<ClickDropdown />
						</td>
					{/each}
					<td class="min-w-[4rem] px-1">30 buổi</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
