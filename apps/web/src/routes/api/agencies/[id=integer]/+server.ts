import {
	CreateAgencyBodySchema,
	// findOneAgency,
	type UpdateAgencyBody
} from '$lib'
import { type RequestEvent, json, error } from '@sveltejs/kit'

export async function GET({ params, locals }: RequestEvent) {
	const { id } = params
    const { agencyRepo } = locals
	if (!id) {
		throw error(400, 'Bad request')
	}

	try {
		const res = await agencyRepo.FindOne(Number(id))
		if (res === undefined) {
			throw error(404, 'Not found')
		}

		return json(res)
	} catch (e: any) {
		console.error('findOneAgency err: ', e)

		if (e.status !== undefined && e.status === 404) {
			throw e
		}
		throw error(500, 'Internal server error')
	}
}

export async function PUT({ request, params, locals }: RequestEvent) {
	const { agencyRepo } = locals
	const { id } = params
	if (!id) {
		throw error(400, 'Bad request')
	}

	const body = await request.json()

	try {
		const validatedReq = await CreateAgencyBodySchema.parseAsync(body)
		const req: UpdateAgencyBody = {
			...validatedReq,
			id: Number(id),
		}
		const res = await agencyRepo.Update(req)

		return json(res)
	} catch (e) {
		console.error(e)
		throw error(500, 'Internal server error')
	}
}
