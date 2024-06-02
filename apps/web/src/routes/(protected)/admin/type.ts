
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
    tempAdress: string,
    updatedAt?: string,
    parentBirthPlace: string,
    parentDob: string,
    parentGender: boolean,
    parentLandlord: string,
    parentName: string,
    parentOccupation: string,
    parentPhoneNumber: string,
    parentResRegistration: string,
    parentRoi: string,
    studentId: number,
    parentZalo: string
}

export type ParentProps = {
    parentBirthPlace: string,
    createdAt?: string,
    parentDob: string,
    parentGender: boolean,
    id?: number,
    parentLandlord: string,
    parentName: string,
    parentOccupation: string,
    parentPhoneNumber: string,
    parentResRegistration: string,
    parentRoi: string,
    studentId: number,
    updatedAt?: string,
    parentZalo: string
    
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