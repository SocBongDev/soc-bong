import {
	CreateRegistrationBodySchema,
	findOneRegistration,
	type UpdateRegistrationBody
} from '$lib'
import { type RequestEvent, json, error } from '@sveltejs/kit'

export async function GET({ params }: RequestEvent) {
	const { id } = params
	if (!id) {
		throw error(400, 'Bad request')
	}

	try {
		const res = await findOneRegistration(Number(id))
		if (res === undefined) {
			throw error(404, 'Not found')
		}

		return json(res)
	} catch (e: any) {
		console.error('findOneRegistration err: ', e)

		if (e.status !== undefined && e.status === 404) {
			throw e
		}
		throw error(500, 'Internal server error')
	}
}

export async function PUT({ request, params, locals }: RequestEvent) {
	const { registrationRepo } = locals
	const { id } = params
	if (!id) {
		throw error(400, 'Bad request')
	}

	const body = await request.json()

	try {
		const validatedReq = await CreateRegistrationBodySchema.parseAsync(body)
		const req: UpdateRegistrationBody = {
			...validatedReq,
			note: validatedReq.note || null,
			id: Number(id),
			isProcessed: false
		}
		const res = await registrationRepo.Update(req)

		return json(res)
	} catch (e) {
		console.error(e)
		throw error(500, 'Internal server error')
	}
}
