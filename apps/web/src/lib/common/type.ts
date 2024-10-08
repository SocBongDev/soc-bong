export type RegistrationProps = {
	createdAt?: string
	id?: number
	isProcessed?: boolean
	note?: string
	parentName: string
	phoneNumber: string
	studentClass: string
	studentDob: string
	studentName: string
	updatedAt?: string
	agencyId: number
}

export type AgencyProps = {
	address: string
	createdAt?: string
	email: string
	id?: number
	name: string
	phone: string
	updatedAt?: string
}

export type StudentProps = {
	agencyId: number
	birthPlace: string
	classId: number
	createdAt?: string
	dob: string //date
	enrolledAt: string //date
	ethnic: string
	firstName: string
	gender: string
	id?: number
	lastName: string
	permanentAddressCommune: string
	permanentAddressDistrict: string
	permanentAddressProvince: string
	tempAddress: string
	updatedAt?: string

	fatherBirthPlace: string
	motherBirthPlace: string
	fatherDob: string
	motherDob: string
	fatherName: string
	motherName: string
	fatherOccupation: string
	motherOccupation: string
	parentLandLord: string
	fatherPhoneNumber: string
	motherPhoneNumber: string
	parentResRegistration: string
	parentRoi: string
	parentZalo: string
}

export type ParentProps = {
	id?: number
	createdAt?: string
	updatedAt?: string
	father_birth_place: string
	mother_birth_place: string
	father_dob: string
	mother_dob: string
	father_name: string
	mother_name: string
	father_occupation: string
	mother_occupation: string
	land_lord: string
	father_phone_number: string
	mother_phone_number: string
	parent_res_registration: string
	parent_roi: string
	parent_zalo: string
}

export type ClassesProps = {
	agencyId: number
	createdAt?: string
	grade: string
	id?: number
	name: string
	teacherId: string
	updatedAt?: string
}

export type UserProps = {
	id?: number
	createdAt?: string
	updatedAt?: string
	email: string
	first_name: string
	last_name: string
	password: string
	connection: string
	phone_number: string
	dob: string
	agencyId: number
	auth0_user_id: string
	is_active?: boolean
	verify_email?: boolean
}

export type RoleProps = {
	id?: string
	name: string
	description: string
}
