import {
    CreateParentBodySchema,
    CreateStudentBodySchema,
    type UpdateParentBody,
    type UpdateStudentBody
} from '$lib'
import { db } from '$lib/db'
import { student, parent } from '$lib'
import type { Student, Parent } from '$lib'
import { type RequestEvent, json, error } from '@sveltejs/kit'
import { eq } from 'drizzle-orm'
import { ZodError } from 'zod'

export async function GET({ params }: RequestEvent) {
    const { id } = params
    if (!id) {
        throw error(400, 'Bad request')
    }

    try {
        const result = await db.select().from(student).fullJoin(parent, eq(student.id, parent.studentId)).where(eq(student.id, Number(id)))

        if (result) {
            const res = {
                ...result[0].student,
                studentId: result[0].parent?.studentId,
                parentName: result[0].parent?.parentName,
                parentDob: result[0].parent?.parentDob,
                parentGender: result[0].parent?.parentGender,
                parentPhoneNumber: result[0].parent?.parentPhoneNumber,
                parentZalo: result[0].parent?.parentZalo,
                parentOccupation: result[0].parent?.parentOccupation,
                parentLandlord: result[0].parent?.parentLandlord,
                parentRoi: result[0].parent?.parentRoi,
                parentBirthPlace: result[0].parent?.parentBirthPlace,
                parentResRegistration: result[0].parent?.parentResRegistration,

            }
            return json(res)
        }
        if (!result) { 
            throw error(404, 'Not found')
        }
    } catch (e) {
        console.error('findOneStudent err: ', e)

        if (e instanceof ZodError) {
            throw error(400, 'Bad request')
        }
        console.error('server error',e)
        throw error(500, 'Internal server error')

    }
}

export async function PUT({ request, params, locals }: RequestEvent) {
    const { studentRepo, parentRepo } = locals
    const { id } = params
    if (!id) {
        throw error(400, 'Bad request')
    }

    const body: Record<string, Student & Parent> = await request.json()
    const student: Record<string, Student | null> = {};
    const parent: Record<string, Parent | null> = {};

    try {
        for (const key in body) {
            if (key.startsWith('parent') || key.startsWith('studentId')) {
                parent[key] = body[key as keyof Parent];
            } else {
                student[key] = body[key as keyof Student];
            }
        }

        const validatedStudentReq = await CreateStudentBodySchema.parseAsync(student)
        const validatedParentReq = await CreateParentBodySchema.parseAsync(parent)
        const req: UpdateStudentBody & UpdateParentBody = {
            ...validatedStudentReq,
            enrollDate: validatedStudentReq.enrollDate || null,
            dob: validatedStudentReq.dob || null,
            birthYear: validatedStudentReq.birthYear || null,
            gender: validatedStudentReq.gender || null,
            ethnic: validatedStudentReq.ethnic || null,
            birthPlace: validatedStudentReq.birthPlace || null,
            tempRes: validatedStudentReq.tempRes || null,
            permResProvince: validatedStudentReq.permResProvince || null,
            permResDistrict: validatedStudentReq.permResDistrict || null,
            permResCommune: validatedStudentReq.permResCommune || null,
            classRoomId: validatedStudentReq.classRoomId,
            id: Number(id),
            ...validatedParentReq,
            parentName: validatedParentReq.parentName,
            parentDob: validatedParentReq.parentDob || null,
            parentGender: validatedParentReq.parentGender || 'male',
            parentPhoneNumber: validatedParentReq.parentPhoneNumber || null,
            parentZalo: validatedParentReq.parentZalo || null,
            parentOccupation: validatedParentReq.parentOccupation || null,
            parentLandlord: validatedParentReq.parentLandlord || null,
            parentRoi: validatedParentReq.parentRoi || null,
            parentBirthPlace: validatedParentReq.parentBirthPlace || null,
            parentResRegistration: validatedParentReq.parentResRegistration || null,
            studentId: Number(id),
        }

        const studentRes = await studentRepo.Update(req)

        const parentRes = await parentRepo.Update(req)

        return json({ studentRes, parentRes })
    } catch (e) {
        console.error(e)
        throw error(500, 'Internal server error')
    }
}