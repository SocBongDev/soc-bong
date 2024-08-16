import { z } from 'zod'

export const CreateUserSchema = z.object({
    // address: z.string().nonempty({ message: 'Địa chỉ không được bỏ trống' }),
    phoneNumber: z
        .string()
        .regex(/^\d+$/, { message: 'Số điện thoại không được chứ ký tự khác ngoài chữ số' })
        .length(10, { message: 'Số điện thoại phải đúng 10 chữ số' }),
    email: z.string().email({ message: 'Email không hợp lệ' }),
    birthDate: z.string().nonempty({ message: "ngày tháng năm sinh không được để trống" }),
    firstName: z.string().nonempty({ message: "Tên không được để trống" }),
    lastName: z.string().nonempty({ message: "Họ không được để trống" }),
    password: z.string().regex(
        /^(?=.*[A-Z])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?])(?=.*[0-9])(?=.*[a-zA-Z]).{8,72}$/,
        "Mật khẩu cần có cả chữ viết hoa, viết thường, số và kí tự đặc biệt và có độ dài tối thiểu là 8"
    ),
    // 
    confirmPassword: z.string().regex(
        /^(?=.*[A-Z])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?])(?=.*[0-9])(?=.*[a-zA-Z]).{8,72}$/,
        "Mật khẩu cần có cả chữ viết hoa, viết thường, số và kí tự đặc biệt và có độ dài tối thiểu là 8"
    ),
}).superRefine(({ confirmPassword, password }, ctx) => {
    if (confirmPassword !== password) {
        ctx.addIssue({
            code: "custom",
            message: "Mật khẩu chưa khớp!",
            path: ['confirmPassword']
        });
    }
});
