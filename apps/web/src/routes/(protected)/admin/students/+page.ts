import type { ClassesProps, AgencyProps, StudentProps } from '$lib/common/type'
import type { PageLoad } from './$types'
import { PUBLIC_API_SERVER_URL } from "$env/static/public"
import dayjs from 'dayjs'
import { classIdStore } from '$lib/store'
import { get } from 'svelte/store';

export const load: PageLoad = async ({ fetch, url, depends }) => {
	const page = Number(url.searchParams.get('page') || '1')
	const pageSize = Number(url.searchParams.get('pageSize') || '15')
	const classIdParam = url.searchParams.get('classId');
	const classId = classIdParam ? parseInt(classIdParam) : get(classIdStore);
	const sorted = String(url.searchParams.get('sort') || 'desc')
	const query = new URLSearchParams()
	const token = localStorage.getItem("access_token")
	query.set('page', String(page))
	query.set('pageSize', String(pageSize))
	query.set('sort', String(sorted))
	const res = await fetch(`${PUBLIC_API_SERVER_URL}/students?classId=${classId}&${query}`, {
		method: "GET",
		headers: {
			"Authorization": `Bearer ${token}`,
			"Content-Type": "application/json",
		}
	})
	const agencyRes = await fetch(`${PUBLIC_API_SERVER_URL}/agencies?${query}`, {
		method: "GET",
		headers: {
			"Authorization": `Bearer ${token}`,
			"Content-Type": "application/json",
		}
	})
	const classes = await fetch(`${PUBLIC_API_SERVER_URL}/classes?${query}`, {
		method: "GET",
		headers: {
			"Authorization": `Bearer ${token}`,
			"Content-Type": "application/json",
		}
	})

	const attendances = await fetch(`${PUBLIC_API_SERVER_URL}/attendances?classId=${classId}&period=${dayjs().format('MM-YYYY')}`, {
		method: "GET",
		headers: {
			"Authorization": `Bearer ${token}`,
			"Content-Type": "application/json",
		}
	})
	depends('app:attendances')
	depends('app:students')

	return {
		students: res.json() as Promise<{ page: number; pageSize: number; data: StudentProps[] }>,
		agencies: agencyRes.json() as Promise<{ page: number; pageSize: number; data: AgencyProps[] }>,
		classes: classes.json() as Promise<{ page: number; pageSize: number; data: ClassesProps[] }>,
		attendances: attendances.json()
	}
}
