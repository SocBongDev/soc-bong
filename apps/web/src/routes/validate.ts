import { z } from 'zod'

export const loginSchema = z.object({
	email: z
		.string()
		.email({ message: 'Email không hợp lệ' })
		.nonempty({ message: 'Email không được bỏ trống' }),
	password: z.string().nonempty({ message: 'Mật khẩu không được bỏ trống' })
})

export type LoginSchema = z.infer<typeof loginSchema>
