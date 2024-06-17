// import type { Registration, Agency } from '$lib'
import type { PageLoad } from './$types'
import type { AgencyProps, RegistrationProps } from '../type'

const API_URL = 'http://127.0.0.1:5000/api/v1'

export const load: PageLoad = async ({ fetch, url, depends }) => {
	const page = Number(url.searchParams.get('page') || '1')
	const pageSize = Number(url.searchParams.get('pageSize') || '15')
	const query = new URLSearchParams()
	query.set('page', String(page))
	query.set('pageSize', String(pageSize))

	const res = await fetch(`${API_URL}/registrations?${query}`)
	const agencyRes = await fetch(`${API_URL}/agencies?${query}`)

	depends('app:registrations')
	depends('app:agencies')

	return {
		registrations: res.json() as Promise<{ page: number; pageSize: number; data: RegistrationProps[] }>,
		agencies: agencyRes.json() as Promise<{ page: number; pageSize: number; data: AgencyProps[] }>
	}
}
