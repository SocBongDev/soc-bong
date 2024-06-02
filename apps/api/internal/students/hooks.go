package students

import (
	"github.com/SocBongDev/soc-bong/internal/classes"
	"github.com/SocBongDev/soc-bong/internal/common"
	"github.com/SocBongDev/soc-bong/internal/parents"
	"github.com/pocketbase/dbx"
)

func mapStudents(dbResult []DbStudentResult) []Student {
	studentsMap := make(map[int]*Student)
	for _, v := range dbResult {
		if student, ok := studentsMap[v.Id]; ok {
			student.Parents = append(student.Parents, parents.Parent{
				BaseEntity: common.BaseEntity{
					Id:        v.ParentId,
					CreatedAt: v.ParentCreatedAt,
					UpdatedAt: v.ParentUpdatedAt,
				},
				WriteParentRequest: parents.WriteParentRequest{
					BirthPlace:      v.ParentBirthPlace,
					Dob:             v.ParentDob,
					Gender:          v.ParentGender,
					Landlord:        v.ParentLandlord,
					Name:            v.ParentName,
					Occupation:      v.ParentOccupation,
					PhoneNumber:     v.ParentPhoneNumber,
					ResRegistration: v.ParentResRegistration,
					Roi:             v.ParentRoi,
					Zalo:            v.ParentZalo,
				},
			})
		} else {
			studentsMap[v.Id] = &Student{
				BaseEntity: common.BaseEntity{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Class: classes.Class{
					BaseEntity: common.BaseEntity{
						Id:        v.ClassId,
						CreatedAt: v.ClassCreatedAt,
						UpdatedAt: v.ClassUpdatedAt,
					},
					WriteClassRequest: classes.WriteClassRequest{
						Name:  v.ClassName,
						Grade: v.ClassGrade,
					},
				},
				WriteStudentRequest: WriteStudentRequest{
					BirthPlace:               v.BirthPlace,
					Dob:                      v.Dob,
					EnrolledAt:               v.EnrolledAt,
					Ethnic:                   v.Ethnic,
					FirstName:                v.FirstName,
					Gender:                   v.Gender,
					LastName:                 v.LastName,
					PermanentAddressCommune:  v.PermanentAddressCommune,
					PermanentAddressDistrict: v.PermanentAddressDistrict,
					PermanentAddressProvince: v.PermanentAddressProvince,
					TempAddress:              v.TempAddress,
					AgencyId:                 v.AgencyId,
					ClassId:                  v.ClassId,
				},
				Parents: []parents.Parent{
					{
						BaseEntity: common.BaseEntity{
							Id:        v.ParentId,
							CreatedAt: v.ParentCreatedAt,
							UpdatedAt: v.ParentUpdatedAt,
						},
						WriteParentRequest: parents.WriteParentRequest{
							BirthPlace:      v.ParentBirthPlace,
							Dob:             v.ParentDob,
							Gender:          v.ParentGender,
							Landlord:        v.ParentLandlord,
							Name:            v.ParentName,
							Occupation:      v.ParentOccupation,
							PhoneNumber:     v.ParentPhoneNumber,
							ResRegistration: v.ParentResRegistration,
							Roi:             v.ParentRoi,
							Zalo:            v.ParentZalo,
						},
					},
				},
			}
		}
	}
	students := make([]Student, 0, len(studentsMap))
	for _, s := range studentsMap {
		students = append(students, *s)
	}

	return students
}

func allHook(q *dbx.Query, sliceA any, op func(sliceB any) error) error {
	switch v := sliceA.(type) {
	case *[]Student:
		dbStudentResults := new([]DbStudentResult)
		if err := op(dbStudentResults); err != nil {
			return err
		}

		students := mapStudents(*dbStudentResults)
		*v = students
		return nil
	default:
		return op(sliceA)
	}
}
