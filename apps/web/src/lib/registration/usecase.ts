import { registrations } from '$lib'
import { db } from '$lib/db'
import type { Paginator } from '$lib/common'
import { desc, eq, inArray, ilike, or } from 'drizzle-orm'
import type { CreateRegistrationBody, Registration, UpdateRegistrationBody } from './entity'
import type { PostgresJsDatabase } from 'drizzle-orm/postgres-js'

export interface RegistrationRepository {
	Find(
		paginator: Paginator,
		search: string,
		order: string[]
	): Promise<{ page: number; pageSize: number; data: Registration[] }>
	FindOne(id: number): Promise<Registration | undefined>
	Insert(req: CreateRegistrationBody): Promise<Registration[]>
	Update(req: UpdateRegistrationBody): Promise<Registration[]>
	Delete(id: number[]): Promise<Registration[]>
}

export class RegistrationDrizzleRepo implements RegistrationRepository {
	constructor(private readonly db: PostgresJsDatabase<Record<string, never>>) {}

	async Find(
		paginator: Paginator,
		search: string,
		order: string[]
	): Promise<{ page: number; pageSize: number; data: Registration[] }> {
		paginator.Format()
		let builder = this.db
			.select()
			.from(registrations)
			.where(eq(registrations.isProcessed, false))
			.limit(paginator.PageSize())
			.offset(paginator.Offset())
			.orderBy(desc(registrations.id))

		if (search != '') {
			builder = builder.where(
				or(
					ilike(registrations.parentName, search),
					ilike(registrations.studentName, search),
					ilike(registrations.phoneNumber, search)
				)
			)
		}

		//        builder = builder.orderBy()

		const result = await builder

		return {
			data: result,
			page: paginator.Page(),
			pageSize: paginator.PageSize()
		}
	}

	async FindOne(id: number): Promise<Registration | undefined> {
		return this.db
			.select()
			.from(registrations)
			.where(eq(registrations.id, id))
			.limit(1)
			.then((res) => res.at(0))
	}

	Insert(req: CreateRegistrationBody): Promise<Registration[]> {
		return this.db.insert(registrations).values(req).returning()
	}

	Update(req: UpdateRegistrationBody): Promise<Registration[]> {
		return this.db
			.update(registrations)
			.set({
				note: req.note,
				parentName: req.parentName,
				studentDob: req.studentDob,
				phoneNumber: req.phoneNumber,
				studentName: req.studentName,
				studentClass: req.studentClass,
				isProcessed: req.isProcessed,
				agencyId: req.agencyId
			})
			.where(eq(registrations.id, req.id))
			.returning()
	}

	Delete(ids: number[]): Promise<Registration[]> {
		return this.db.delete(registrations).where(inArray(registrations.id, ids)).returning()
	}
}

export const RegistrationRepo = new RegistrationDrizzleRepo(db)

export async function findRegistrations(paginator: Paginator, search: string, order: string[]) {
	paginator.Format()
	let builder = db
		.select()
		.from(registrations)
		.where(eq(registrations.isProcessed, false))
		.limit(paginator.PageSize())
		.offset(paginator.Offset())
		.orderBy(desc(registrations.id))

	if (search != '') {
		builder = builder.where(
			or(
				ilike(registrations.parentName, search),
				ilike(registrations.studentName, search),
				ilike(registrations.phoneNumber, search)
			)
		)
	}

	//        builder = builder.orderBy()

	const result = await builder

	return {
		data: result,
		page: paginator.Page(),
		pageSize: paginator.PageSize()
	}
}

export async function findOneRegistration(id: number) {
	return db
		.select()
		.from(registrations)
		.where(eq(registrations.id, id))
		.limit(1)
		.then((res) => res.at(0))
}

export async function createRegistration(req: CreateRegistrationBody) {
	return db.insert(registrations).values(req).returning()
}

export async function updateRegistration(req: Registration) {
	return db
		.update(registrations)
		.set({
			note: req.note,
			parentName: req.parentName,
			studentDob: req.studentDob,
			phoneNumber: req.phoneNumber,
			studentName: req.studentName,
			studentClass: req.studentClass,
			isProcessed: req.isProcessed
		})
		.where(eq(registrations.id, req.id))
		.returning()
}
