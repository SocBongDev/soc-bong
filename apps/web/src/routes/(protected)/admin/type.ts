export type RegistrationProps = {
        createdAt?: string,
        id?: number,
        isProcessed?: boolean,
        note?: string,
        parentName: string,
        phoneNumber: string,
        studentClass: string,
        studentDob: string,
        studentName: string,
        updatedAt?: string
}


export type AgencyProps = {
    address: string,
    createdAt?: string,
    email: string,
    id?: number,
    name: string,
    phone: string,
    updatedAt?: string
}

export type StudentProps = {
    agencyId: number,
    birthPlace: string,
    classId: number,
    createdAt?: string,
    dob: string, //date
    enrolledAt: string, //date
    ethnic: string,
    firstName: string,
    gender: string,
    id?: number,
    lastName: string,
    permanentAddressCommune: string,
    permanentAddressDistrict: string,
    permanentAddressProvince: string,
    tempAddress: string,
    updatedAt?: string,

    fatherBirthPlace: string,
    motherBirthPlace: string,
    fatherDob: string,
    motherDob: string,
    fatherName: string,
    motherName: string,
    fatherOccupation: string,
    motherOccupation: string,
    parentLandLord: string,
    parentPhoneNumber: string,
    parentResRegistration: string,
    parentRoi: string,
    parentZalo: string
}

export type ParentProps = {
    id?: number,
    createdAt?: string,
    updatedAt?: string,
    father_birth_place: string,
    mother_birth_place: string,
    father_dob: string,
    mother_dob: string,
    father_name: string,
    mother_name: string,
    father_occupation: string,
    mother_occupation: string,
    land_lord: string,
    parent_phone_number: string,
    parent_res_registration: string,
    parent_roi: string,
    parent_zalo: string
}

export type ClassesProps = {
    agency_id: number,
    createdAt?: string,
    grade: string,
    id?: number,
    name: string,
    teacher_id: string,
    updatedAt?: string
}