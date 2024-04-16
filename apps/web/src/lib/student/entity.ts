import { agencies } from "$lib/agency";
import { pgTable, serial, varchar, integer, foreignKey, unique} from "drizzle-orm/pg-core";
import { createInsertSchema } from 'drizzle-zod'
import type { z } from 'zod'

export const student = pgTable("student", {
    id: serial("id").primaryKey(),
    grade: varchar("grade").notNull(),
    firstName: varchar("first_name").notNull(),
    lastName: varchar("last_name").notNull(),
    enrollDate: varchar("enroll_date"),
    dob: varchar("dob"),
    birthYear: varchar("birth_year"),
    gender: varchar("gender"),
    ethnic: varchar("ethnic"),
    birthPlace: varchar("birth_place"),
    tempRes: varchar("temp_res"),
    permResProvince: varchar("perm_res_province"),
    permResDistrict: varchar("perm_res_district"),
    permResCommune: varchar("perm_res_commune"),
    classRoomId: integer("class_room_id").notNull(),
    agencyId: integer("agency_id").notNull(),
}, (table) => {
    return {
        parentReference: foreignKey({
            columns: [table.agencyId],
            foreignColumns: [agencies.id],
        }),
        unique: unique().on(table.agencyId)
    }
});

export type Student = typeof student.$inferSelect

export type NewStudent = typeof student.$inferInsert

const writeSchema = createInsertSchema(student)

export const CreateStudentBodySchema = writeSchema.omit({
	id: true,
})

export const UpdateStudentBodySchema = writeSchema.omit({}).required()

export type CreateStudentBody = z.infer<typeof CreateStudentBodySchema>

export type UpdateStudentBody = z.infer<typeof UpdateStudentBodySchema>
