import { parent, student } from '$lib'
import { db } from '$lib/db'
import type { Paginator } from '$lib/common'
import { desc, eq, inArray, ilike, or } from 'drizzle-orm'
import type { CreateParentBody, Parent, UpdateParentBody } from './entity'
import type { PostgresJsDatabase } from 'drizzle-orm/postgres-js'
import dayjs from 'dayjs'

export interface ParentRepository {
    Find(
        paginator: Paginator,
        search: string,
        order: string[]
    ): Promise<{ page: number; pageSize: number; data: Parent[] }>
    FindOne(id: number): Promise<Parent | undefined>
    Insert(req: CreateParentBody): Promise<Parent[]>
    Update(req: UpdateParentBody): Promise<Parent[]>
    Delete(id: number[]): Promise<Parent[]>
}

export class ParentDrizzleRepo implements ParentRepository {
    constructor(private readonly db: PostgresJsDatabase<Record<string, never>>) { }

    async Find(
        paginator: Paginator,
        search: string,
        order: string[]
    ): Promise<{ page: number; pageSize: number; data: Parent[] }> {
        paginator.Format()
        let builder = this.db
            .select()
            .from(parent)
            .limit(paginator.PageSize())
            .offset(paginator.Offset())
            .orderBy(desc(parent.id))

        if (search != '') {
            builder = builder.where(
                or(
                    ilike(parent.parentName, search),
                    ilike(parent.studentId, search),
                    ilike(parent.parentZalo, search)
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

    async FindOne(id: number): Promise<Parent | undefined> {
        return this.db
            .select()
            .from(parent)
            .where(eq(parent.id, id))
            .limit(1)
            .then((res) => res.at(0))
    }

    async Insert(req: CreateParentBody): Promise<Parent[]> {
        return db
		.select()
		.from(parent)
		.where(
			eq(parent.parentName, req.parentName) && 
			eq(parent.parentGender, req.parentGender) && 
			eq(parent.parentDob, dayjs(req.parentDob).format('YYYY-MM-DD'))
		).then((existingRow: Parent[]) => {
			if (existingRow.length > 0) {
				throw Error('Student Parent already exists')
			}	else {	
                return db.insert(parent).values(req).returning()
			}
		}).catch((error: any) => {
			console.error('Error:', error);
			throw Error('Student Parent already exists')
		  });
    }

    Update(req: UpdateParentBody): Promise<Parent[]> {
        return this.db
            .update(parent)
            .set({
                parentName: req.parentName,
                parentDob: req.parentDob,
                parentGender: req.parentGender,
                parentPhoneNumber: req.parentPhoneNumber,
                parentZalo: req.parentZalo,
                parentOccupation: req.parentOccupation || '',
                parentLandlord: req.parentLandlord || '', 
                parentRoi: req.parentRoi || '',
                parentBirthPlace: req.parentBirthPlace || '',
                parentResRegistration: req.parentResRegistration || '',
            })
            .where(eq(parent.studentId, req.studentId))
            .returning()
    }

    Delete(ids: number[]): Promise<Parent[]> {
        return this.db.delete(parent).where(inArray(parent.studentId, ids)).returning()
    }
}

export const ParentRepo = new ParentDrizzleRepo(db)
