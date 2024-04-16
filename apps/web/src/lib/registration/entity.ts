import { agencies } from '$lib/agency';
import { boolean, char, date, pgTable, serial, timestamp, foreignKey, unique, integer } from 'drizzle-orm/pg-core'
import { createInsertSchema } from 'drizzle-zod'
import type { z } from 'zod'

export const registrations = pgTable('registrations', {
	id: serial('id').primaryKey(),
	createdAt: timestamp('created_at').notNull().defaultNow(),
	isProcessed: boolean('is_processed').notNull().default(false),
	note: char('note'),
	parentName: char('parentName').notNull(),
	phoneNumber: char('phoneNumber').notNull(),
	studentClass: char('studentClass').notNull(),
	studentDob: date('studentDob').notNull(),
	studentName: char('studentName').notNull(),
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

export type Registration = typeof registrations.$inferSelect

export type NewRegistration = typeof registrations.$inferInsert

const writeSchema = createInsertSchema(registrations)

export const CreateRegistrationBodySchema = writeSchema.omit({
	id: true,
	createdAt: true,
	isProcessed: true
})

export const UpdateRegistrationBodySchema = writeSchema.omit({ createdAt: true }).required()

export type CreateRegistrationBody = z.infer<typeof CreateRegistrationBodySchema>

export type UpdateRegistrationBody = z.infer<typeof UpdateRegistrationBodySchema>
