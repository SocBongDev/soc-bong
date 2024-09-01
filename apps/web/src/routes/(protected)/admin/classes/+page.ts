
import type { PageLoad } from './$types'
import { PUBLIC_API_SERVER_URL } from "$env/static/public"
import type { AgencyProps, ClassesProps, UserProps } from '$lib/common/type'

export const load: PageLoad = async ({ fetch, url, depends }) => {
    const page = Number(url.searchParams.get('page') || '1')
    const pageSize = Number(url.searchParams.get('pageSize') || '15')
    const query = new URLSearchParams()
    const token = localStorage.getItem("access_token");

    query.set('page', String(page))
    query.set('pageSize', String(pageSize))
    
    const headers = {
        "Authorization": `Bearer ${token}`,
        "Content-Type": "application/json",
    }

    const fetchData = async (url: string) => {
        const response = await fetch(url, { method: "GET", headers })
        return response.json()
    }

	const [classesData, agenciesData, usersData] = await Promise.all([
        fetchData(`${PUBLIC_API_SERVER_URL}/classes?${query}`),
        fetchData(`${PUBLIC_API_SERVER_URL}/agencies?${query}`),
        fetchData(`${PUBLIC_API_SERVER_URL}/users?${query}`),
    ])

    depends('app:classes')

    return {
        classes: classesData as Promise<{ page: number; pageSize: number; data: ClassesProps[] }>,
        agencies: agenciesData as Promise<{page: number; pageSize: number; data: AgencyProps[]}>,
        users: usersData as Promise<{page: number, pageSize: number, data: UserProps[]}>
    }
}