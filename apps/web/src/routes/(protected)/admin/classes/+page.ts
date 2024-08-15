
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

    const classes = await fetch(`${PUBLIC_API_SERVER_URL}/classes?${query}`,
        {
            method: "GET",
            headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
            }
        }
    )

    const agencies = await fetch(`${PUBLIC_API_SERVER_URL}/agencies?${query}`,
        {
            method: "GET",
            headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
            }
        }
    )

    const users = await fetch(`${PUBLIC_API_SERVER_URL}/users?${query}`, 
        {
            method: "GET",
            headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
            }
        }
    )
    depends('app:classes')

    return {
        classes: classes.json() as Promise<{ page: number; pageSize: number; data: ClassesProps[] }>,
        agencies: agencies.json() as Promise<{page: number; pageSize: number; data: AgencyProps[]}>,
        users: users.json() as Promise<{page: number, pageSize: number, data: UserProps[]}>
    }
}