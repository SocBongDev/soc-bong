// import type { Registration, Agency } from '$lib'
import type { PageLoad } from './$types'
import type { AgencyProps, RegistrationProps } from '../type'
import { PUBLIC_API_SERVER_URL } from "$env/static/public"

export const load: PageLoad = async ({ fetch, url, depends }) => {
	const page = Number(url.searchParams.get('page') || '1')
	const pageSize = Number(url.searchParams.get('pageSize') || '15')
	const query = new URLSearchParams()

	const token = localStorage.getItem("access_token");
	query.set('page', String(page))
	query.set('pageSize', String(pageSize))

	const registration = await fetch(
		`${PUBLIC_API_SERVER_URL}/registrations?${query}`,
		{
			method: "GET",
			headers: {
				Authorization: `Bearer ${token}`,
				"Content-Type": "application/json",
			}
		}
	)
	const agencies = await fetch(
		`${PUBLIC_API_SERVER_URL}/agencies?${query}`,
		{
			method: "GET",
			headers: {
				Authorization: `Bearer ${token}`,
				"Content-Type": "application/json",
			}
		}
	)


	depends('app:registrations')

	return {
		registrations: registration.json() as Promise<{ page: number; pageSize: number; data: RegistrationProps[] }>,
		agencies: agencies.json() as Promise<{ page: number; pageSize: number; data: AgencyProps[] }>
	}
}
