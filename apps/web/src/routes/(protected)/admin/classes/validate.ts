import { z } from 'zod'

export const CreateClassesSchema = z.object({
	name: z.string().nonempty({ message: 'Tên lớp được bỏ trống' }),
	grade: z.enum(['seed', 'buds', 'leaf', 'toddler']),
    teacherId: z.string().nonempty({message: 'mã giáo viên không được để trống'}),
    agencyId: z.number().nonnegative({message: 'mã cơ sở không được để trống'})
})
