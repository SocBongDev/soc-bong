import { pgTable, serial, timestamp, varchar } from 'drizzle-orm/pg-core'
import { createInsertSchema } from 'drizzle-zod'
import type { z } from 'zod'

export const agencies = pgTable('agencies', {
	id: serial('id').primaryKey(),
	createdAt: timestamp('created_at').notNull().defaultNow(),
	updatedAt: timestamp('updated_at').notNull().defaultNow(),
	agencyName: varchar('agency_name', { length: 100 }).notNull(),
	agencyAddress: varchar('agency_address', { length: 100 }).notNull(),
	agencyPhone: varchar('agency_phone', { length: 100 }).notNull(),
	agencyEmail: varchar('agency_email', { length: 100 }).notNull(),
})

export type Agency = typeof agencies.$inferSelect

export type NewAgency = typeof agencies.$inferInsert

const writeSchema = createInsertSchema(agencies)

export const CreateAgencyBodySchema = writeSchema
	.omit({ id: true, createdAt: true, updatedAt: true })
	.required()

export const UpdateAgencyBodySchema = writeSchema
	.omit({ createdAt: true, updatedAt: true })
	.required()

export type CreateAgencyBody = z.infer<typeof CreateAgencyBodySchema>

export type UpdateAgencyBody = z.infer<typeof UpdateAgencyBodySchema>
