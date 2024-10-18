import type { PageLoad } from './$types'
import type { AgencyProps, RegistrationProps } from '$lib/common/type'
import { PUBLIC_API_SERVER_URL } from '$env/static/public'

export const load: PageLoad = async ({ fetch, url, depends }) => {
	const page = Number(url.searchParams.get('page') || '1')
	const pageSize = Number(url.searchParams.get('pageSize') || '15')
	const query = new URLSearchParams()

	const token = localStorage.getItem('access_token')
	query.set('page', String(page))
	query.set('pageSize', String(pageSize))

	const headers = {
		Authorization: `Bearer ${token}`,
		'Content-Type': 'application/json'
	}

	const fetchData = async (url: string) => {
		const response = await fetch(url, { method: 'GET', headers })
		return response.json()
	}

	const [registrationsData, agenciesData] = await Promise.all([
		fetchData(`${PUBLIC_API_SERVER_URL}/registrations?${query}`),
		fetchData(`${PUBLIC_API_SERVER_URL}/agencies?${query}`)
	])

	depends('app:registrations')

	return {
		registrations: registrationsData as Promise<{
			page: number
			pageSize: number
			data: RegistrationProps[]
		}>,
		agencies: agenciesData as Promise<{ page: number; pageSize: number; data: AgencyProps[] }>
	}
}
