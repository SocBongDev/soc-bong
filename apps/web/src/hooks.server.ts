import { PUBLIC_SUPABASE_ANON_KEY, PUBLIC_SUPABASE_URL } from '$env/static/public'
import { AgencyRepo, ParentRepo, RegistrationRepo, StudentRepo} from '$lib'
import { createSupabaseServerClient } from '@supabase/auth-helpers-sveltekit'
import type { Handle } from '@sveltejs/kit'
import { redirect } from '@sveltejs/kit'

export const handle: Handle = async ({ event, resolve }) => {
	event.locals.supabase = createSupabaseServerClient({
		supabaseUrl: PUBLIC_SUPABASE_URL,
		supabaseKey: PUBLIC_SUPABASE_ANON_KEY,
		event
	})

	event.locals.registrationRepo = RegistrationRepo
	event.locals.studentRepo = StudentRepo
	event.locals.parentRepo = ParentRepo
	event.locals.agencyRepo = AgencyRepo

	event.locals.getSession = async () => {
		const {
			data: { session }
		} = await event.locals.supabase.auth.getSession()
		return session
	}

	if (event.url.pathname.startsWith('/admin')) {
		const session = await event.locals.getSession()
		if (!session) {
			throw redirect(303, '/')
		}
	}

	return resolve(event, {
		filterSerializedResponseHeaders(name) {
			return name === 'content-range'
		}
	})
}
