import type { ClassesProps, AgencyProps, StudentProps } from '$lib/common/type'
import type { PageLoad } from './$types'
import { PUBLIC_API_SERVER_URL } from '$env/static/public'
import dayjs from 'dayjs'
import { classIdStore } from '$lib/store'
import { get } from 'svelte/store'

export const load: PageLoad = async ({ fetch, url, depends }) => {
	const page = Number(url.searchParams.get('page') || '1')
	const pageSize = Number(url.searchParams.get('pageSize') || '15')
	const classIdParam = url.searchParams.get('classId')
	const classId = classIdParam ? parseInt(classIdParam) : get(classIdStore)
	const sorted = String(url.searchParams.get('sort') || 'desc')
	const query = new URLSearchParams()
	const token = localStorage.getItem('access_token')
	query.set('page', String(page))
	query.set('pageSize', String(pageSize))
	query.set('sort', String(sorted))

	const headers = {
		Authorization: `Bearer ${token}`,
		'Content-Type': 'application/json'
	}

	const fetchData = async (url: string) => {
		const response = await fetch(url, { method: 'GET', headers })
		return response.json()
	}

	const [studentsData, agenciesData, classesData, attendancesData] = await Promise.all([
		fetchData(`${PUBLIC_API_SERVER_URL}/students?classId=${classId}&${query}`),
		fetchData(`${PUBLIC_API_SERVER_URL}/agencies?${query}`),
		fetchData(`${PUBLIC_API_SERVER_URL}/classes?${query}`),
		fetchData(
			`${PUBLIC_API_SERVER_URL}/attendances?classId=${classId}&period=${dayjs().format('MM-YYYY')}`
		)
	])

	depends('app:attendances')
	depends('app:students')

	return {
		students: studentsData as Promise<{ page: number; pageSize: number; data: StudentProps[] }>,
		agencies: agenciesData as Promise<{ page: number; pageSize: number; data: AgencyProps[] }>,
		classes: classesData as Promise<{ page: number; pageSize: number; data: ClassesProps[] }>,
		attendances: attendancesData
	}
}
