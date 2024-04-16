import type { Agency, Student } from '$lib'
import type { PageLoad } from './$types'

export const load: PageLoad = async ({ fetch, url, depends }) => {
	const page = Number(url.searchParams.get('page') || '1')
	const pageSize = Number(url.searchParams.get('pageSize') || '15')
	const query = new URLSearchParams()
	query.set('page', String(page))
	query.set('pageSize', String(pageSize))

	const res = await fetch(`/api/students?${query}`)
	const agencyRes = await fetch(`/api/agencies?${query}`)

	depends('app:students')
	depends('app:agencies')

	return {
		students: res.json() as Promise<{ page: number; pageSize: number; data: Student[] }>,
		agencies: agencyRes.json() as Promise<{ page: number; pageSize: number; data: Agency[] }>
	}
}
