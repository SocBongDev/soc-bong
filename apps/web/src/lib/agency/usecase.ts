import { agencies, registrations, student } from '$lib'
import { db } from '$lib/db'
import type { Paginator } from '$lib/common'
import { desc, eq, inArray, ilike, or } from 'drizzle-orm'
import type { CreateAgencyBody, Agency, UpdateAgencyBody } from './entity'
import type { PostgresJsDatabase } from 'drizzle-orm/postgres-js'
import { Notify } from '$lib/store'

export interface AgencyRepository {
	Find(
		paginator: Paginator,
		search: string,
		order: string[]
	): Promise<{ page: number; pageSize: number; data: Agency[] }>
	FindOne(id: number): Promise<Agency | undefined>
	Insert(req: CreateAgencyBody): Promise<Agency[]>
	Update(req: UpdateAgencyBody): Promise<Agency[]>
	Delete(id: number[]): Promise<Agency[]>
}

export class AgencyDrizzleRepo implements AgencyRepository {
	constructor(private readonly db: PostgresJsDatabase<Record<string, never>>) {}

	async Find(
		paginator: Paginator,
		search: string,
		order: string[]
	): Promise<{ page: number; pageSize: number; data: Agency[] }> {
		paginator.Format()
		let builder = this.db
			.select()
			.from(agencies)
			.limit(paginator.PageSize())
			.offset(paginator.Offset())
			.orderBy(desc(agencies.id))

		if (search != '') {
			builder = builder.where(or(
				ilike(agencies.agencyName, search),
				ilike(agencies.agencyAddress, search),
				ilike(agencies.agencyPhone, search),
				ilike(agencies.agencyEmail, search),
				
				))
		}

		//        builder = builder.orderBy()

		const result = await builder
		return {
			data: result,
			page: paginator.Page(),
			pageSize: paginator.PageSize()
		}
	}

	async FindOne(id: number): Promise<Agency | undefined> {
		return this.db
			.select()
			.from(agencies)
			.where(eq(agencies.id, id))
			.limit(1)
			.then((res) => res.at(0))
	}

	async Insert(req: CreateAgencyBody): Promise<Agency[]> {
		return db
			.select()
			.from(agencies)
			.where(eq(agencies.agencyName, req.agencyName))
			.then((existingRow) => {
				if (existingRow.length > 0) {
					throw Error('Agency already exists')
				} else {
					return this.db.insert(agencies).values(req).returning()
				}
			})
			.catch((error) => {
				console.error('Error:', error)
				throw Error('Agency already exists')
			})
	}

	Update(req: UpdateAgencyBody): Promise<Agency[]> {
		return this.db
			.update(agencies)
			.set({
				agencyName: req.agencyName,
				agencyAddress: req.agencyAddress,
				agencyPhone: req.agencyPhone,
				agencyEmail: req.agencyEmail,
			})
			.where(eq(agencies.id, req.id))
			.returning()
	}

	Delete(ids: number[]): Promise<Agency[]> {
		
		return this.db.delete(agencies).where(inArray(agencies.id, ids)).returning()
	}
}

export const AgencyRepo = new AgencyDrizzleRepo(db)

export async function deleteOneAgency(id: number) {
	return db.delete(agencies).where(eq(agencies.id, id)).returning()
}

async function findStudentsAndRegistrationsInAgencies(ids: number[]) {
	try {
	  const studentIncludesAgency = await db.select(
		{ id: student.id },
	  ).from(student).where(inArray(student.agencyId, ids));
  
	  const registrationIncludesAgency = await db.select(
		{ id: registrations.id },
	  ).from(registrations).where(inArray(registrations.agencyId, ids));
  
	  return { students: studentIncludesAgency, registrations: registrationIncludesAgency };
	} catch (error) {
	  console.error('Error:', error);
	}
  }

export async function deleteAgencies(ids: number[]) {
	const checkStudentAndRegistration = await findStudentsAndRegistrationsInAgencies(ids);	

	try {
		if (checkStudentAndRegistration && ( checkStudentAndRegistration.students.length === 0 && checkStudentAndRegistration.registrations.length === 0 )) {
			return db.delete(agencies).where(inArray(agencies.id, ids)).returning()
		}  else {
			throw new Error('Can not Delete Agency because it has students or registrations')
		}
	} catch (e) {
		console.error('server error',e)

		throw Error('Internal server error')
	}
}

