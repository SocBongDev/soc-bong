import { z } from 'zod'

export const CreateAgencySchema = z.object({
	name: z.string().nonempty({ message: 'Tên trung tâm không được bỏ trống' }),
	address: z.string().nonempty({ message: 'Địa chỉ không được bỏ trống' }),
	phone: z.string().nonempty({ message: 'Số điện thoại không được bỏ trống' }),
	email: z.string().email({ message: 'Email không hợp lệ' })
})
