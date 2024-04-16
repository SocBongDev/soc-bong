import { CreateParentBodySchema} from "$lib";
import { error, json, type RequestEvent } from "@sveltejs/kit";
import { CreateStudentBodySchema } from "$lib";
import { ZodError } from "zod";
import postgres from "postgres";
import { Paginator, Sort } from "$lib/common";
import {deleteOneStudent } from "$lib/student/usecase";
import { FindStudentQuery } from "$lib";
import type { Parent, Student } from "$lib";
import { z } from "zod";

export async function GET({ url, locals }: RequestEvent) {
    const { studentRepo } = locals;
    const search = url.searchParams.get('search') || '';
    const orderBy = url.searchParams.get('orderBy') || '';
    const paginator = new Paginator(url.searchParams);
    const sort = new Sort(url.searchParams);
    const query = new FindStudentQuery(paginator, sort, search, orderBy);
    try {
        const res = await studentRepo.Find(paginator, search, [orderBy]);
        return json(res);
    } catch (e) {
        throw error(500, 'Unknown error');
    }
}

export async function POST({ request, locals }: RequestEvent) {
    const { studentRepo, parentRepo } = locals;
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

        const reqStudent = await CreateStudentBodySchema.parseAsync(student);
        const students = await studentRepo.Insert(reqStudent);

        if (students.length > 0) {
            const studentId = students[0].id;
            const parentObj = {
                ...parent,
                studentId: parseInt(studentId.toString()),
            };

            const reqParent = await CreateParentBodySchema.parseAsync(parentObj);
            
            try {
                const parents = await parentRepo.Insert(reqParent);
                return json({ students, parents });
            } catch (parentError) {
                // Roll back student creation if parent creation fails
                await deleteOneStudent(studentId);
                throw parentError;
            }
        } else {
            throw error(400, 'Create student failed');
        }
    } catch (e) {
        if (e instanceof ZodError) {
            console.error('error zod: ', e);
            throw error(400, 'Bad request');
        }

        if (e instanceof postgres.PostgresError) {
            console.error('error postgres: ', e);
            throw error(500, 'Internal server error');
        }
        console.error('error: ', e);
        throw error(500, 'Internal server error');
    }
}

export async function DELETE({ request, locals }: RequestEvent) {
    const { studentRepo, parentRepo } = locals
    const body = await request.json()
    const deleteBodySchema = z.object({ ids: z.number().array().nonempty() })

    try {
        const { ids } = await deleteBodySchema.parseAsync(body)
        const resParent = await parentRepo.Delete(ids)
        if (resParent) {
            const res = await studentRepo.Delete(ids)
            return json({res, resParent})
        } else {
            throw error(404, 'Can not Delete')
        }

    } catch (e) {
        if (e instanceof ZodError) {
            throw error(400, 'Bad request')
        }
        console.error('server error',e)

        throw error(500, 'Internal server error')
    }
}