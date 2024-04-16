import { z } from 'zod'

export const CreateAgencySchema = z.object({
	agencyName: z.string().nonempty({ message: 'Tên trung tâm không được bỏ trống' }),
	agencyAddress: z.string().nonempty({ message: 'Địa chỉ không được bỏ trống' }),
	agencyPhone: z.string().nonempty({ message: 'Số điện thoại không được bỏ trống' }),
	agencyEmail: z.string().email({ message: 'Email không hợp lệ' }),
})
