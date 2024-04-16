import { student } from '$lib'
import { db } from '$lib/db'
import type { Paginator } from '$lib/common'
import { desc, eq, inArray, ilike, or } from 'drizzle-orm'
import type { CreateStudentBody, Student, UpdateStudentBody } from './entity'
import type { PostgresJsDatabase } from 'drizzle-orm/postgres-js'
import dayjs from 'dayjs'

export interface StudentRepository {
	Find(
		paginator: Paginator,
		search: string,
		order: string[]
	): Promise<{ page: number; pageSize: number; data: Student[] }>
	FindOne(id: number): Promise<Student | undefined>
	Insert(req: CreateStudentBody): Promise<Student[]>
	Update(req: UpdateStudentBody): Promise<Student[]>
	Delete(id: number[]): Promise<Student[]>
}

export class StudentDrizzleRepo implements StudentRepository {
	constructor(private readonly db: PostgresJsDatabase<Record<string, never>>) {}

	async Find(
		paginator: Paginator,
		search: string,
		order: string[]
	): Promise<{ page: number; pageSize: number; data: Student[] }> {
		paginator.Format()
		let builder = this.db
			.select()
			.from(student)
			.limit(paginator.PageSize())
			.offset(paginator.Offset())
			.orderBy(desc(student.id))

		if (search != '') {
			builder = builder.where(
				or(
					ilike(student.firstName, search),
					ilike(student.lastName, search),
					ilike(student.classRoomId, search)
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

	async FindOne(id: number): Promise<Student | undefined> {
		return this.db
			.select()
			.from(student)
			.where(eq(student.id, id))
			.limit(1)
			.then((res) => res.at(0))
	}

	async Insert(req: CreateStudentBody): Promise<Student[]> {
		return db
		.select()
		.from(student)
		.where(
			eq(student.firstName, req.firstName) && 
			eq(student.lastName, req.lastName) && 
			eq(student.dob, dayjs(req.dob).format('YYYY-MM-DD'))
		).then((existingRow) => {
			if (existingRow.length > 0) {
				throw Error('Student already exists')
			}	else {	
				return this.db.insert(student).values(req).returning()
			}
		}).catch((error) => {
			console.error('Error:', error);
			throw Error('Student already exists')
		  });
	}

	Update(req: UpdateStudentBody): Promise<Student[]> {
		return this.db
			.update(student)
			.set({
                grade: req.grade,
                firstName: req.firstName,
                lastName: req.lastName,
                enrollDate: req.enrollDate,
                dob: req.dob,
                birthYear: req.birthYear,
                birthPlace: req.birthPlace,
                gender: req.gender,
                ethnic: req.ethnic,
                tempRes: req.tempRes,
                permResProvince: req.permResProvince,
                permResDistrict: req.permResDistrict,
                permResCommune: req.permResCommune,
                classRoomId: req.classRoomId,
				agencyId: req.agencyId
			})
			.where(eq(student.id, req.id))
			.returning()
	}

	Delete(ids: number[]): Promise<Student[]> {
		return this.db.delete(student).where(inArray(student.id, ids)).returning()
	}
}

export const StudentRepo = new StudentDrizzleRepo(db)

export async function deleteOneStudent(id: number) {
	return db.delete(student).where(eq(student.id, id)).returning()
}