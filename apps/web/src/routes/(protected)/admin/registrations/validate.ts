import { z } from 'zod'

export const CreateRegistrationSchema = z.object({
	parentName: z.string().nonempty({ message: 'Tên phụ huynh không được bỏ trống' }),
	studentName: z.string().nonempty({ message: 'Tên bé không được bỏ trống' }),
	phoneNumber: z
		.string()
		.regex(/^\d+$/, { message: 'Số điện thoại không được chứ ký tự khác ngoài chữ số' })
		.length(10, { message: 'Số điện thoại phải đúng 10 chữ số' }),
	studentDob: z.string().nonempty({ message: 'Ngày sinh của bé không được bỏ trống' }),
	studentClass: z.enum(['seed', 'buds', 'leaf', 'toddlers']),
	agencyId: z.number().int().positive({ message: 'Đơn vị không hợp lệ' }),
})
