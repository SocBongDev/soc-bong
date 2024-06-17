import type { ClassesProps, AgencyProps, StudentProps } from '../type'
import type { PageLoad } from './$types'

const API_URL = 'http://127.0.0.1:5000/api/v1'

export const load: PageLoad = async ({ fetch, url, depends }) => {
	const page = Number(url.searchParams.get('page') || '1')
	const pageSize = Number(url.searchParams.get('pageSize') || '15')
	const classId = Number(url.searchParams.get('classId') || '1')
	const query = new URLSearchParams()
	query.set('page', String(page))
	query.set('pageSize', String(pageSize))

	const res = await fetch(`${API_URL}/students?classId=${classId}&${query}`)
	const agencyRes = await fetch(`${API_URL}/agencies?${query}`)
	const classes = await fetch(`${API_URL}/classes?${query}`)

	depends('app:students')

	return {
		students: res.json() as Promise<{ page: number; pageSize: number; data: StudentProps[] }>,
		agencies: agencyRes.json() as Promise<{ page: number; pageSize: number; data: AgencyProps[] }>,
		classes: classes.json() as Promise<{ page: number; pageSize: number; data: ClassesProps[] }>
	}
}
