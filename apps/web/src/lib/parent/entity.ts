import { pgTable, serial, varchar,unique, foreignKey, integer} from "drizzle-orm/pg-core";
import { createInsertSchema } from 'drizzle-zod'
import type { z } from 'zod'
import {student} from "$lib";



export const parent = pgTable("parent", {
    id: serial("id").primaryKey(),
    studentId: integer("student_id").notNull(),
    parentName: varchar("name").notNull(),
    parentDob: varchar("dob"),
    parentGender: varchar("sex").notNull(),
    parentPhoneNumber: varchar("phone_number"),
    parentZalo: varchar("zalo"),
    parentOccupation: varchar("occupation"),
    parentLandlord: varchar("landlord"),
    parentRoi: varchar("roi"),
    parentBirthPlace: varchar("birthplace"),
    parentResRegistration: varchar("res_registration"),
}, (table) => {
    return {
        parentReference: foreignKey({
            columns: [table.studentId],
            foreignColumns: [student.id],
        }),
        unique: unique().on(table.studentId)
    }
});

export type Parent = typeof parent.$inferSelect

export type NewParent = typeof parent.$inferInsert

const writeSchema = createInsertSchema(parent)

export const CreateParentBodySchema = writeSchema.omit({
	id: true,
})

export const UpdateParentBodySchema = writeSchema.omit({id: true}).required()

export type CreateParentBody = z.infer<typeof CreateParentBodySchema>

export type UpdateParentBody = z.infer<typeof UpdateParentBodySchema>