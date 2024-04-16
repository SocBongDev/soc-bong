// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
import type { RegistrationRepository, StudentRepository, ParentRepository, AgencyRepository } from '$lib'
import 'unplugin-icons/types/svelte'

declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			supabase: SupabaseClient<Database>
			registrationRepo: RegistrationRepository
			studentRepo: StudentRepository
			parentRepo: ParentRepository
			agencyRepo: AgencyRepository
			getSession(): Promise<Session | null>
		}
		interface PageData {
			session: Session | null
		}
		// interface Platform {}
	}
}

export {}
