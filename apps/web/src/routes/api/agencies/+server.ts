import { CreateAgencyBodySchema, deleteAgencies} from "$lib";
import { error, json, type RequestEvent } from "@sveltejs/kit";
import { ZodError } from "zod";
import postgres from "postgres";
import { Paginator, Sort } from "$lib/common";
import { FindAgencyQuery } from "$lib";
import type { Agency } from "$lib";
import { z } from "zod";

export async function GET({ url, locals }: RequestEvent) {
    const { agencyRepo } = locals;
    const search = url.searchParams.get('search') || '';
    const orderBy = url.searchParams.get('orderBy') || '';
    const paginator = new Paginator(url.searchParams);
    const sort = new Sort(url.searchParams);
    const query = new FindAgencyQuery(paginator, sort, search, orderBy);
    try {
        const res = await agencyRepo.Find(paginator, search, [orderBy]);
        return json(res);
    } catch (e) {
        throw error(500, 'Unknown error');
    }
}

export async function POST({ request, locals }: RequestEvent) {
    const { agencyRepo } = locals;
    const body = await request.json();
    try {
		const req = await CreateAgencyBodySchema.parseAsync(body)
        
		const res = await agencyRepo.Insert(req)
		return json(res)
	} catch (e) {
		if (e instanceof ZodError) {
            console.log('check e', e)
			throw error(400, 'Bad request')
		}

		if (e instanceof postgres.PostgresError) {
			throw error(500, 'Internal server error')
		}

		throw error(500, 'Internal server error')
	}
}

// Delete Agency when the student table and registration has that agency field was deleted 
export async function DELETE({ request}: RequestEvent) {
    const body = await request.json()
    const deleteBodySchema = z.object({ ids: z.number().array().nonempty() })

    try {
        const { ids } = await deleteBodySchema.parseAsync(body)
        const res = await deleteAgencies(ids)
        
        return json({res})
    } catch (e) {
        if (e instanceof ZodError) {
            throw error(400, 'Bad request')
        }
        console.error('server error',e)

        throw error(500, 'Internal server error')
    }
}