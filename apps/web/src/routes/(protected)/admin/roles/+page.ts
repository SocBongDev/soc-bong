import type { RoleProps } from '$lib/common/type'
import type { PageLoad } from './$types'
import { PUBLIC_API_SERVER_URL } from "$env/static/public"

export const load: PageLoad = async ({ fetch, url, depends }) => {
	const page = Number(url.searchParams.get('page') || '1')
	const pageSize = Number(url.searchParams.get('pageSize') || '15')
	const search = String(url.searchParams.get('search') || '')
	const query = new URLSearchParams()
	const token = localStorage.getItem("access_token")
	query.set('page', String(page))
	query.set('pageSize', String(pageSize))
	query.set('search', String(search))
	const res = await fetch(`${PUBLIC_API_SERVER_URL}/roles`, {
		method: "GET",
		headers: {
			"Authorization": `Bearer ${token}`,
			"Content-Type": "application/json; charset=utf-8",
		}
	})
	
	depends('app:roles')

	return {
		roles: res.json() as Promise<{ page: number; pageSize: number; data: RoleProps[] }>,
	}
}
