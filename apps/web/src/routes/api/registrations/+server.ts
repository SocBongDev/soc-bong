import { Paginator, Sort } from '$lib/common'
import { CreateRegistrationBodySchema } from '$lib/registration'
import { FindRegistrationQuery } from '$lib/registration/payload'
import { createRegistration, findRegistrations } from '$lib/registration/usecase'
import { error, json, type RequestEvent } from '@sveltejs/kit'
import postgres from 'postgres'
import { ZodError, z } from 'zod'

export async function GET({ url }: RequestEvent) {
	const search = url.searchParams.get('search') || ''
	const orderBy = url.searchParams.get('orderBy') || ''
	const paginator = new Paginator(url.searchParams)
	const sort = new Sort(url.searchParams)
	const query = new FindRegistrationQuery(paginator, sort, search, orderBy)

	try {
		const result = await findRegistrations(paginator, search, [orderBy])
		return json(result)
	} catch (e) {
		throw error(500, 'Unknown error')
	}
}

export async function POST({ request }: RequestEvent) {
	const body = await request.json()
	try {
		const req = await CreateRegistrationBodySchema.parseAsync(body)
		const res = await createRegistration(req)

		return json(res)
	} catch (e) {
		if (e instanceof ZodError) {
			throw error(400, 'Bad request')
		}

		if (e instanceof postgres.PostgresError) {
			throw error(500, 'Internal server error')
		}

		throw error(500, 'Internal server error')
	}
}

export async function DELETE({ request, locals }: RequestEvent) {
	const { registrationRepo } = locals
	const body = await request.json()
	const deleteBodySchema = z.object({ ids: z.number().array().nonempty() })

	try {
		const { ids } = await deleteBodySchema.parseAsync(body)
		const res = await registrationRepo.Delete(ids)

		return json(res)
	} catch (e) {
		console.error(e)

		if (e instanceof ZodError) {
			throw error(400, 'Bad request')
		}

		throw error(500, 'Internal server error')
	}
}
