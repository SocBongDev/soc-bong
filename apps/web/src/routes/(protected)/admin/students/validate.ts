import { z } from 'zod'

export const CreateStudentSchema = z.object({
	firstName: z.string().nonempty({ message: 'Tên bé không được bỏ trống' }),
	lastName: z.string().nonempty({ message: 'Họ và tên lót bé không được bỏ trống' }),

	enrolledAt: z.string(),
	dob: z.string(),

	gender: z.string(),
	ethnic: z.string(),
	birthPlace: z.string(),
	tempAddress: z.string(),
	permanentAddressProvince: z.string(),
	permanentAddressDistrict: z.string(),
	permanentAddressCommune: z.string(),
	agencyId: z.number().int().positive({ message: 'Trường học không được bỏ trống' }),
	classId: z.number().optional(),

	fatherBirthPlace: z.string(),
	motherBirthPlace: z.string(),
	fatherDob: z.string(),
	motherDob: z.string(),
	fatherName: z.string(),
	motherName: z.string(),
	fatherOccupation: z.string(),
	motherOccupation: z.string(),
	parentLandLord: z.string(),
	fatherPhoneNumber: z.string(),
	motherPhoneNumber: z.string(),
	parentResRegistration: z.string(),
	parentRoi: z.string(),
	parentZalo: z.string()
})

export const CreateParentSchema = z.object({
	id: z.number().int().optional(),
	fatherBirthPlace: z.string(),
	motherBirthPlace: z.string(),
	fatherDob: z.string(),
	motherDob: z.string(),
	fatherName: z.string(),
	motherName: z.string(),
	fatherOccupation: z.string(),
	motherOccupation: z.string(),
	parentLandLord: z.string(),
	fatherPhoneNumber: z.string(),
	motherPhoneNumber: z.string(),
	parentResRegistration: z.string(),
	parentRoi: z.string(),
	parentZalo: z.string()
})

export const CreateClassSchema = z.object({
	agencyId: z.number().int().positive({ message: 'Trường học không được bỏ trống' }),
	grade: z.enum(['buds, seed, leaf ']),
	name: z.string().nonempty({ message: 'Tên giáo viên không được để trống' }),
	teacherId: z.number().int(),
	createdAt: z.string(),
	updatedAt: z.string()
})
